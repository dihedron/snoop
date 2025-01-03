package sink

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/dihedron/snoop/pipeline"
)

// Acknowledge is a sink that acknowledges message.
type Acknowledge struct{}

func (*Acknowledge) Collect(ctx context.Context, message pipeline.Message) error {
	if message != nil {
		slog.Debug("absorbing message with acknowledgement", "message", message, "type", fmt.Sprintf("%T", message))
		if err := message.Ack(false); err != nil {
			slog.Error("error acknowledging message", "error", err)
		}
	} else {
		slog.Warn("discarding nil message!")
	}
	return nil
}
