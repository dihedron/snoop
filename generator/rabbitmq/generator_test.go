package rabbitmq

import (
	"context"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/dihedron/rawdata"
	"github.com/dihedron/snoop/format"
	"github.com/dihedron/snoop/test"
	"github.com/joho/godotenv"
)

func TestRabbitMQContext(t *testing.T) {
	test.Setup(t)
	godotenv.Load()

	rmq := &RabbitMQ{}
	rawdata.UnmarshalInto("@"+os.Getenv("FILE"), rmq)
	slog.Debug("RabbitMQ configuration file in JSON format", "configuration", format.ToPrettyJSON(rmq))

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	count := 0
	for m := range rmq.All(ctx) {
		count++
		if count == 10 {
			break
		}
		slog.Debug("message received", "value", format.ToPrettyJSON(m))
	}
	if err := rmq.Err(); err != nil {
		slog.Error("error iterating over RabbitMQ messages", "error", err)
	}
}

func TestRabbitMQConfigurationAndValidation(t *testing.T) {
	test.Setup(t)
	godotenv.Load()

	rmq := &RabbitMQ{}
	slog.Debug("RabbitMQ configuration file in JSON format", "configuration", format.ToJSON(rmq))

	rawdata.UnmarshalInto("@"+os.Getenv("FILE"), rmq)
	slog.Debug("RabbitMQ configuration file in JSON format", "configuration", format.ToJSON(rmq))

	rmq.Validate()
}
