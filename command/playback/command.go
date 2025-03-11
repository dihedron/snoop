package playback

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"log/slog"

	"github.com/dihedron/snoop/generator/textfile"
	"github.com/dihedron/snoop/metadata"
	"github.com/dihedron/snoop/openstack/amqp"
	"github.com/dihedron/snoop/openstack/notification"
	"github.com/dihedron/snoop/openstack/oslo"
	"github.com/dihedron/snoop/syslog"
	"github.com/dihedron/snoop/transform/chain"
	"github.com/dihedron/snoop/transform/transformers"
	"github.com/juju/rfc/v2/rfc5424"
)

// Embed the file content as string.
//
//go:embed compute.instance.tmpl
var templ string

type Handler func(n *notification.Notification, sl *syslog.Syslog) error

var handlers map[string]Handler

// Playback is the command that reads message from a recording on file and
// processes them one by one.
// ./snoop playback 20220818.amqp.messages
type Playback struct {
}

// Execute is the real implementation of the Playback command.
func (cmd *Playback) Execute(args []string) error {
	if len(args) == 0 {
		slog.Error("no input files")
		return errors.New("no input files provided")
	}
	slog.Debug("reading messages from recording..", "files", args)

	unwrap := chain.Of4(
		transformers.StringToByteArray(),
		amqp.JSONToMessage(),
		oslo.MessageToOslo(false),
		notification.OsloToNotification(false),
	)

	sl, err := syslog.New(syslog.WithApplication(metadata.Name))
	if err != nil {
		slog.Error("error initialising syslog", "error", err)
		return err
	}

	ctx := context.Background()
	files := textfile.New()
	for line := range files.AllLinesContext(ctx, args...) {
		var (
			n   notification.Notification
			err error
		)
		if n, err = unwrap(line); err != nil {
			slog.Error("error processing line", "line", line)
			continue
		}
		slog.Info("processed line", "line", line, "notification", n)

		switch n.Summary().EventType {
		case "identity.authenticate":
			if n, ok := n.(*notification.Identity); ok {
				onIdentityAuthenticate(n, sl)
			}

		}

	}

	return nil
}

func onIdentityAuthenticate(n *notification.Identity, sl *syslog.Syslog) error {
	fmt.Printf("identity.authenticate: user %s\n", n.Payload.Reason.ReasonType)
	type SyslogIdentityEvent struct {
		UserID  string  `json:"userid"`
		Success bool    `json:"success"`
		Reason  *string `json:"reason,omitempty"`
	}
	e := &SyslogIdentityEvent{
		UserID: n.ContextUserID,
		//...
	}

	if err := sl.Send(&syslog.Message{
		Facility: rfc5424.FacilityAuthpriv,
		Severity: rfc5424.SeverityEmergency,
		ID:       "Login",
		Content:  e,
		Data: map[string][]string{
			"user":   {"name=John", "surname=Doe", "id=a123456"},
			"tenant": {"region=regionOne", "name=event-broker", "id=1234567890"},
		},
	}); err != nil {
		slog.Error("error submitting message to syslog", "error", err)
		return err
	}
	return nil
}

/*
// Execute is the real implementation of the Playback command.
func (cmd *Playback) writeOutSingleEventsAsFiles(args []string) error {
	if len(args) == 0 {
		slog.Error("no input files")
		return errors.New("no input files provided")
	}
	slog.Debug("reading messages from recording..", "files", args)

	stopwatch := &transformers.StopWatch[string, notification.Notification]{}
	multicache := &transformers.MultiCache[string, notification.Notification]{}
	xform := chain.Of7(
		stopwatch.Start(),
		transformers.StringToByteArray(),
		amqp.JSONToMessage(),
		oslo.MessageToOslo(false),
		notification.OsloToNotification(false),
		multicache.Set(func(n notification.Notification) string {
			return n.Summary().EventType
		}),
		stopwatch.Stop(),
	)

	ctx := context.Background()
	files := textfile.New()
	for line := range files.AllLinesContext(ctx, args...) {
		if value, err := xform(line); err != nil {
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
*/
