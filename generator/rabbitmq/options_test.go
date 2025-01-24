package rabbitmq

import (
	"log/slog"
	"os"
	"testing"

	"github.com/dihedron/rawdata"
	"github.com/dihedron/snoop/format"
	"github.com/dihedron/snoop/test"
	"github.com/joho/godotenv"
)

func TestRabbitMQToJSON(t *testing.T) {
	test.Setup(t, test.Text)
	godotenv.Load()

	rmq := &RabbitMQ{}
	slog.Debug("RabbitMQ configuration file in JSON format", "configuration", format.ToJSON(rmq))

	rawdata.UnmarshalInto("@"+os.Getenv("FILE"), rmq)
	slog.Debug("RabbitMQ configuration file in JSON format", "configuration", format.ToJSON(rmq))

	options := rmq.ToOptions()
	slog.Debug("RabbitMQ options in JSON format", "options", format.ToJSON(options))
}
