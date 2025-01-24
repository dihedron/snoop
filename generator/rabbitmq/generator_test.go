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
	test.Setup(t, test.Text)
	godotenv.Load()

	rmq := &RabbitMQ{}
	rawdata.UnmarshalInto("@"+os.Getenv("FILE"), rmq)
	slog.Debug("RabbitMQ configuration file in JSON format", "configuration", format.ToPrettyJSON(rmq))

	options := rmq.ToOptions()
	slog.Debug("RabbitMQ options in JSON format", "configuration", format.ToPrettyJSON(options))

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	count := 0
	for m := range RabbitMQContext(ctx, rmq) {
		count++
		if count == 10 {
			break
		}
		slog.Debug("message received", "value", format.ToPrettyJSON(m))
	}
}
