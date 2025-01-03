package rabbitmq

import (
	"testing"

	"github.com/dihedron/snoop/format"
)

func TestRabbitMQToJSON(t *testing.T) {

	rmq := &RabbitMQ{}
	t.Logf("RabbitMQ configuration file in JSON format: %v", format.ToPrettyJSON(rmq))
}
