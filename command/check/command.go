package check

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/dihedron/rawdata"
	"github.com/dihedron/snoop/command/base"
	"github.com/dihedron/snoop/format"
	"github.com/dihedron/snoop/generator/rabbitmq"
	"github.com/fatih/color"
	"github.com/go-playground/validator/v10"
)

// Ping is the command that checks connectivity against the RabbitMQ servers
// in the given configuration.
type Check struct {
	base.ConfiguredCommand
	// DrainCount represents the number of messages to drain from the
	// RabbitMQ source in order to test if things actually work; it will
	// not acknowledge any message which will therefore be re-delivered.
	DrainCount int `short:"d" long:"drain-count" description:"The number of messages to drain for testing purposes." optional:"yes" default:"1" validate:"gte=1"`
}

// Execute is the real implementation of the Check command.
func (cmd *Check) Execute(args []string) error {

	if cmd.Configuration == nil {
		slog.Error("no configuration provided")
		return errors.New("no configuration provided")
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(cmd); err != nil {
		slog.Error("error validating command struct", "error", err)
		return err
	}

	slog.Debug("running with configuration", "configuration", cmd.Configuration)

	rmq := &rabbitmq.RabbitMQ{}
	err := rawdata.UnmarshalInto("@"+*cmd.Configuration, rmq)
	if err != nil {
		slog.Error("error reading configuration file", "error", err)
	}

	if err := validate.Struct(rmq); err != nil {
		slog.Error("error validating configuration struct", "error", err)
		return err
	}

	slog.Debug("RabbitMQ configuration file in JSON format", "configuration", format.ToJSON(rmq))

	fmt.Printf("%s:\n%s", color.YellowString("configuration"), color.BlueString(format.ToYAML(rmq)))

	// options := rmq.ToOptions()
	// slog.Debug("RabbitMQ options in JSON format", "options", format.ToJSON(options))

	// fmt.Printf("%s:\n%s", color.YellowString("options"), color.BlueString(format.ToYAML(options)))

	// now prepare the processing chain
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	if cmd.DrainCount <= 0 {
		cmd.DrainCount = 1
	}
	fmt.Printf("%s %s...\n", color.YellowString("connecting to"), color.RedString("servers"))
	count := 0
	for m := range rmq.All(ctx) {
		fmt.Printf("%s...\n", color.GreenString("retrieved message"))
		slog.Debug("message received", "value", format.ToPrettyJSON(m))
		count++
		if count >= cmd.DrainCount {
			break
		}
	}
	if err := rmq.Error(); err != nil {
		slog.Error("error connecting to RabbitMQ", "error", err)
		return err
	}

	return nil
}
