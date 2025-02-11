package playback

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"time"

	"github.com/dihedron/snoop/format"
	"github.com/dihedron/snoop/generator/file"
	"github.com/dihedron/snoop/openstack/amqp"
	"github.com/dihedron/snoop/openstack/notification"
	"github.com/dihedron/snoop/openstack/oslo"
	. "github.com/dihedron/snoop/transform"
	"github.com/dihedron/snoop/transformers"
)

func doCountSpecificEventTypes(args []string, acceptedEvents ...string) error {

	ctx := context.Background()
	start := time.Now()

	stopwatch := &transformers.StopWatch[string, notification.Notification]{}
	multicounter := &transformers.MultiCounter[notification.Notification, string]{}

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
							multicounter.AddIf(
								func(n notification.Notification) string {
									return n.Summary().EventType
								},
								func(n notification.Notification) bool {
									return slices.Contains(acceptedEvents, n.Summary().EventType)
								},
							),
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
	os.Stdout.Sync()

	counts, total := multicounter.Count()
	fmt.Printf("\nprocessed %d messages total in %s\n", total, time.Now().Sub(start).String())
	subcount := int64(0)
	for k, v := range counts {
		fmt.Printf("  %-50s: %d\n", k, v)
		subcount = subcount + v
	}
	fmt.Printf("  %-50s: %d\n", "others", total-subcount)

	return nil
}

func doRecordSpecificEventTypesWithFormat(args []string, format string, filter func(n notification.Notification) bool) error {

	ctx := context.Background()

	stopwatch := &transformers.StopWatch[string, string]{}

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
							transformers.AcceptIf(filter),
							Then(
								transformers.Format[notification.Notification](format),
								Then(
									transformers.Write[string](os.Stdout, true),
									stopwatch.Stop(),
								),
							),
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
	os.Stdout.Sync()

	return nil
}

func doEventTypesStats(args []string) error {

	ctx := context.Background()
	start := time.Now()

	stopwatch := &transformers.StopWatch[string, notification.Notification]{}
	multicounter := &transformers.MultiCounter[notification.Notification, string]{}

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
							multicounter.Add(func(n notification.Notification) string {
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
	os.Stdout.Sync()

	counts, total := multicounter.Count()
	fmt.Printf("\nprocessed %d messages total in %s\n", total, time.Now().Sub(start).String())
	subcount := int64(0)
	for k, v := range counts {
		fmt.Printf("  %-50s: %d\n", k, v)
		subcount = subcount + v
	}
	fmt.Printf("  %-50s: %d\n", "others", total-subcount)

	return nil
}

func printEventsAsYAML(args []string, filter func(n notification.Notification) bool) error {

	ctx := context.Background()
	stopwatch := &transformers.StopWatch[string, notification.Notification]{}
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
							transformers.AcceptIf(filter),
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
			fmt.Println("# --------------------------------------------------------------------------------")
			fmt.Printf("%s", format.ToYAML(value))
		}
	}
	fmt.Println("# --------------------------------------------------------------------------------")
	os.Stdout.Sync()

	return nil
}
