package record

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/dihedron/rawdata"
	"github.com/dihedron/snoop/command/base"
	"github.com/dihedron/snoop/command/common"
	"github.com/dihedron/snoop/format"
	"github.com/dihedron/snoop/generator/rabbitmq"
	"github.com/dihedron/snoop/openstack/amqp"
	"github.com/dihedron/snoop/transform/chain"
	"github.com/dihedron/snoop/transform/transformers"
	"github.com/rabbitmq/amqp091-go"
)

// Record is the command that reads message from RabbitMQ and dumps them
// out to standard output or to a file.
// ./snoop record --connnection-info=tests/rabbitmq/brokerd.yaml --output=20220818.amqp.messages
type Record struct {
	base.Command
	// Profile contains the path to the configuration file to use to connect to
	// a RabbitMQ instance.
	Profile string `short:"p" long:"profile" description:"The path to the file containing the RabbitMQ connection info (aka profile)." required:"yes" env:"SNOOP_PROFILE"`
	// Truncate is used to specify whether the output file (if one is specified)
	// should be truncated before writing to it.
	Truncate *bool `short:"t" long:"truncate" description:"Whether the output file should be truncated or appended to (default)." optional:"yes" env:"SNOOP_TRUNCATE"`
	// Limit is used to specify the number of messages to process before exiting.
	Limit *int `short:"l" long:"limit" description:"Whether to process only the given amount of messages." optional:"yes" hidden:"yes" env:"SNOOP_LIMIT"`
}

// Execute is the real implementation of the Record command.
func (cmd *Record) Execute(args []string) error {
	slog.Debug("draining and recording messages from RabbitMQ")

	// validate input parameters first
	if err := common.Validate(*cmd); err != nil {
		slog.Error("error validating command struct", "error", err)
		return err
	}

	// get output path
	path := "-" // stdout
	if len(args) > 0 {
		path = args[0]
	}

	// get the messages writer
	writer, err := common.GetWriter(path, cmd.Truncate)
	if err != nil {
		slog.Error("error getting writer", "error", err)
		return err
	}
	if w, ok := writer.(io.Closer); ok {
		defer w.Close()
	}

	// get the RabbitMQ connection
	if cmd.Profile == "" {
		slog.Error("no connection info provided")
		return errors.New("no connection info provided")
	}

	slog.Debug("reading connection info", "connection info", cmd.Profile)

	rmq := &rabbitmq.RabbitMQ{}
	err = rawdata.UnmarshalInto("@"+cmd.Profile, rmq)
	if err != nil {
		slog.Error("error reading connection info", "error", err)
	}
	slog.Debug("RabbitMQ connection info file in JSON format", "configuration", format.ToJSON(rmq))

	// now prepare the processing chain
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	stopwatch := &transformers.StopWatch[*amqp091.Delivery, []byte]{}
	xform := chain.Of4(
		stopwatch.Start(),
		amqp.DeliveryToMessage(false),
		transformers.ToJSON[*amqp.Message](),
		stopwatch.Stop(),
	)

	count := 0
	limit := 0
	if cmd.Limit != nil && *cmd.Limit > 0 {
		limit = *cmd.Limit
	}
	for m := range rmq.All(ctx) {
		count++
		if count%100 == 0 {
			fmt.Printf(". ")
		}
		if cmd.Limit != nil && count >= limit {
			break
		}
		value, err := xform(m)
		if err != nil {
			slog.Error("error applying chain to message", "error", err)
		} else {
			slog.Debug("AMQP091 message received", "value", format.ToPrettyJSON(value))
			fmt.Fprintf(writer, "%s\n", value)
			slog.Debug("acknowledging incoming AMQP message")
			if err := m.Ack(false); err != nil {
				slog.Error("error acknowledging message", "id", m.MessageId, "error", err)
			}
		}
	}
	if err := rmq.Err(); err != nil {
		slog.Error("error connecting to RabbitMQ", "error", err)
	}

	return nil
}
