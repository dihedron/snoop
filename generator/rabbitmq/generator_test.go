package rabbitmq

import (
	"context"
	"log"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/dihedron/rawdata"
	"github.com/dihedron/snoop/format"
	"github.com/dihedron/snoop/test"
	"github.com/joho/godotenv"
)

func TestRabbitMQContextWithConsumeOnce(t *testing.T) {
	test.Setup(t)
	if err := godotenv.Load(); err != nil {
		slog.Error("error loading environment", "error", err)
	}

	rmq := &RabbitMQ{}
	rawdata.UnmarshalInto("@"+os.Getenv("FILE"), rmq)
	slog.Debug("RabbitMQ configuration file in JSON format", "configuration", format.ToPrettyJSON(rmq))

	log.Printf("CONFIGURATION:\n%s\n", format.ToPrettyJSON(rmq))

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
	if err := godotenv.Load(".env"); err != nil {
		slog.Error("error loading environment", "error", err)
	}

	rmq := &RabbitMQ{}
	slog.Debug("RabbitMQ configuration file in JSON format", "configuration", format.ToJSON(rmq))

	rawdata.UnmarshalInto("@"+os.Getenv("FILE"), rmq)
	slog.Debug("RabbitMQ configuration file in JSON format", "configuration", format.ToJSON(rmq))

	rmq.Validate()
}
