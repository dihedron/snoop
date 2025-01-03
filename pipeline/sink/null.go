package sink

import (
	"context"
	"log/slog"

	"github.com/dihedron/snoop/pipeline"
)

// Null is the default, do-nothing sink.
type Null struct{}

func (*Null) Collect(ctx context.Context, message pipeline.Message) error {
	slog.Debug("absorbing message without acknowledgement", "message", message)
	message.Ack(false)
	return nil
}
