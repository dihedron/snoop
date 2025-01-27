package record

/*
import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/dihedron/snoop/format"
	"github.com/dihedron/snoop/message"
	"github.com/dihedron/snoop/pipeline"
	"github.com/dihedron/snoop/pipeline/source/file"
	"go.uber.org/zap"
)

type OpenStackJSONMessageParser struct{}

func (*OpenStackJSONMessageParser) Name() string {
	return "github.com/dihedron/snoop/command/record/OpenStackJSONMessageParser"
}

func (c *OpenStackJSONMessageParser) Process(ctx context.Context, quit chan<- bool, msg pipeline.Message) (pipeline.Message, context.Context, error) {
	select {
	case <-ctx.Done():
		slog.Debug("context cancelled")
		return msg, ctx, pipeline.ErrInterrupted
	default:
		slog.Debug("parsing JSON message into OpenStack notification", "type", format.TypeAsString(msg))
		// TODO: read message
		if msg, ok := msg.(*file.Message); ok {
			not, err := message.NewNotificationFromJSON(msg.Value)
			if err != nil {
				slog.Error("error pasring OpenStack notification from JSON", "error", err)
			}
			slog.Debug("notification parsed")
			return &message.OpenStackMessage{
				Message: not,
			}, ctx, nil
		}
	}
	return msg, ctx, nil
}

func TestCreateVM(t *testing.T) {
	source, err := file.NewSource("create_vm.messages")
	if err != nil {
		slog.With(zap.Error(err)).Error("error creating file source")
	}
	// startTime := filter.NewStartTime("0")
	// endTime := filter.NewEndTime("0")
	parser := &OpenStackJSONMessageParser{}
	correlator, err := NewCorrelator()
	if err != nil {
		slog.With(zap.Error(err)).Error("error creating correlator")
	}
	pipeline := pipeline.New(
		pipeline.WithSource(source),
		pipeline.WithFilters(
			// startTime,
			// filter.NewThrottler(100*time.Millisecond),
			parser,
			correlator,
			// endTime,
		),
		pipeline.WithErrorCallback(func(ctx context.Context, quit chan<- bool, filter string, err error) {
			slog.With(zap.Error(err)).Errorf("callback called on filter %s", filter)
			t.Fail()
		}),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	results, quit, _ := pipeline.Start(ctx)
	defer pipeline.Close()
loop:
	for {
		select {
		case result := <-results:
			slog.Debugf("result retrieved: %v", result)
			result.Ack(false)
		case <-ctx.Done():
			slog.Debug("pipeline context cancelled")
			break loop
		case <-quit:
			slog.Debug("pipeline received quit message")
			break loop
		}
	}
	// slog.Debugf("result (%v each):\n%v", endTime.Elapsed())
}
*/
