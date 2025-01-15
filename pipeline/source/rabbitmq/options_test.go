package rabbitmq

import (
	"os"
	"testing"

	"github.com/dihedron/rawdata"
	"github.com/dihedron/snoop/format"
	"github.com/joho/godotenv"
)

func TestRabbitMQToJSON(t *testing.T) {

	godotenv.Load()

	rmq := &RabbitMQ{}
	t.Logf("RabbitMQ configuration file in JSON format: %v", format.ToPrettyJSON(rmq))

	rawdata.UnmarshalInto("@"+os.Getenv("FILE"), rmq)
	t.Logf("RabbitMQ configuration file in JSON format: %v", format.ToPrettyJSON(rmq))

	options := rmq.ToOptions()
	t.Logf("RabbitMQ options in JSON format: %v", format.ToPrettyJSON(options))

}
