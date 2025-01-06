package rabbitmq

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/dihedron/rawdata"
	"github.com/dihedron/snoop/format"
	"github.com/joho/godotenv"
)

func TestRabbitMQContext(t *testing.T) {

	godotenv.Load()

	rmq := &RabbitMQ{}
	rawdata.UnmarshalInto("@"+os.Getenv("FILE"), rmq)
	t.Logf("RabbitMQ configuration file in JSON format: %v", format.ToPrettyJSON(rmq))

	options := rmq.ToOptions()
	t.Logf("RabbitMQ options in JSON format: %v", format.ToPrettyJSON(options))

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	for m := range RabbitMQContext(ctx, rmq) {
		t.Logf("message received: %s", format.ToPrettyJSON(m))
	}
}
