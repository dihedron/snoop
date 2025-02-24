package check

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"log/slog"

	"github.com/dihedron/rawdata"
	"github.com/dihedron/snoop/command/base"
	"github.com/dihedron/snoop/command/common"
	"github.com/dihedron/snoop/format"
	"github.com/dihedron/snoop/generator/rabbitmq"
	"github.com/fatih/color"
)

// Ping is the command that checks connectivity against the RabbitMQ servers
// in the given configuration.
type Check struct {
	base.Command
	// ConnectionInfo contains the path to the (optional) configuration file to use to
	// connect to a RabbitMQ instance; if no value is provided (neither on the
	// command line nor in the environment via the SNOOP_CONNECT variable), the
	// application will look for a viable configuration file named .snoop.yaml
	// under a few well-known paths: /etc, the current directory etc.
	Profile string `short:"p" long:"profile" description:"The path to the file containing the RabbitMQ connection info (aka profile)." required:"yes" env:"SNOOP_PROFILE" validate:"file"`
}

const MaxMessages = 1

// Execute is the real implementation of the Check command.
func (cmd *Check) Execute(args []string) error {

	// TODO: is this needed?
	if cmd.Profile == "" {
		slog.Error("no profile provided")
		return errors.New("no profile provided")
	}

	if err := common.Validate(cmd); err != nil {
		slog.Error("error validating command struct", "error", err)
		return err
	}

	slog.Debug("connection profile available", "path", cmd.Profile)

	rmq := &rabbitmq.RabbitMQ{}
	err := rawdata.UnmarshalInto("@"+cmd.Profile, rmq)
	if err != nil {
		slog.Error("error reading connection info file", "error", err)
	}

	if err := common.Validate(rmq); err != nil {
		slog.Error("error validating connection info struct", "error", err)
		return err
	}

	slog.Debug("RabbitMQ connection info file in JSON format", "connection info", format.ToJSON(rmq))
	fmt.Printf("%s:\n%s", color.YellowString("connection info"), color.BlueString(format.ToYAML(rmq)))

	// check if servers can be contacted, one at a time
	servers := rmq.Servers
	var result error = nil
	for _, server := range servers {

		func() {
			fmt.Printf("connecting to server %s on port %s:", color.YellowString(server.Address), color.YellowString(fmt.Sprintf("%d", server.Port)))

			// overwrite the address of the serves in the configuration
			// with the single one under test now
			rmq.Servers = []rabbitmq.Server{server}
			rmq.Reset()

			// now prepare the processing chain
			slog.Debug("initialising context...")
			ctx := context.Background()

			count := 0
			for m := range rmq.All(ctx) {
				fmt.Printf(" %s", color.GreenString(" ✔"))
				m.Nack(true, true)
				//slog.Debug("message received", "value", format.ToPrettyJSON(m))
				slog.Debug("range loop: message received")
				count++
				if count >= MaxMessages {
					slog.Debug("range loop: maximum count of messages, breaking from loop")
					break
				}
			}
			if err := rmq.Err(); err != nil {
				slog.Error("error from iterator", "error", err)
				fmt.Printf(" %s (%s)", color.RedString("✘"), err.Error())
				result = errors.Join(result, fmt.Errorf("error connecting to %s: %w", server.Address, err))
			}
			fmt.Println()
		}()
	}

	return result
}
