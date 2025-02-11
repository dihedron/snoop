package playback

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/dihedron/snoop/format"
	"github.com/dihedron/snoop/generator/file"
	"github.com/dihedron/snoop/openstack/amqp"
	"github.com/dihedron/snoop/openstack/notification"
	"github.com/dihedron/snoop/openstack/oslo"
	. "github.com/dihedron/snoop/transform"
	"github.com/dihedron/snoop/transformers"
)

// Embed the file content as string.
//
//go:embed compute.instance.tmpl
var templ string

// Playback is the command that reads message from a recording on file and
// processes them one by one.
// ./snoop play 20220818.amqp.messages
type Playback struct{}

// Execute is the real implementation of the Playback command.
func (cmd *Playback) Execute(args []string) error {
	if len(args) == 0 {
		slog.Error("no input files")
		return errors.New("no input files provided")
	}
	slog.Debug("reading messages from recording..", "files", args)

	ctx := context.Background()
	stopwatch := &transformers.StopWatch[string, notification.Notification]{}
	multicache := &transformers.MultiCache[string, notification.Notification]{}
	chain := Apply(
		stopwatch.Start(),
		Then(
			transformers.StringToByteArray(),
			Then(
				amqp.JSONToMessage(),
				Then(
					oslo.MessageToOslo(false),
					Then(
						notification.OsloToNotification(false),
						Then(
							multicache.Set(func(n notification.Notification) string {
								return n.Summary().EventType
							}),
							stopwatch.Stop(),
						),
					),
				),
			),
		),
	)

	for line := range file.LinesContext(ctx, args...) {
		if value, err := chain(line); err != nil {
			slog.Error("error processing line", "line", line)
		} else {
			slog.Info("processed line", "line", line, "output", value)
		}
	}

	for k, messages := range multicache.Contents() {
		path, err := filepath.Abs(filepath.Join(".", "output", k)) //strings.Replace(k, ".", "_", -1)))
		if err != nil {
			slog.Error("error making absolute path of path", "path", filepath.Join(".", "output", k), "error", err)
			return err
		}
		fmt.Printf("would make directory: %s (%d items)\n", path, len(messages))
		err = os.MkdirAll(path, 0755)
		if err != nil {
			slog.Error("error creating directories", "path", path, "error", err)
			return err
		}
		for i, m := range messages {
			func(i int, message notification.Notification) {
				name := filepath.Join(path, fmt.Sprintf("%04d-%s.yaml", i, k))
				file, err := os.Create(name)
				if err != nil {
					slog.Error("error creating output file", "name", name, "error", err)
					return
				}
				defer file.Close()
				file.WriteString(format.ToYAML(message))
			}(i, m)
		}
	}

	// fmt.Println("# --------------------------------------------------------------------------------")
	// fmt.Printf("%s", format.ToYAML(value))

	// fmt.Println("# --------------------------------------------------------------------------------")
	// os.Stdout.Sync()

	// return printEventsAsYAML(args, func(n notification.Notification) bool {
	// 	return strings.HasPrefix(n.Summary().EventType, "compute.instance.")
	// })

	// return printEventsAsYAML(args, func(n notification.Notification) bool { return true })
	return nil
}
