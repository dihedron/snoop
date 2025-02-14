package check

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"log/slog"

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
}

const MaxMessages = 1

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
