package process

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"sort"

	"github.com/dihedron/snoop/command/base"
	"github.com/dihedron/snoop/generator/file"
	"github.com/dihedron/snoop/openstack/amqp"
	"github.com/dihedron/snoop/openstack/notification"
	"github.com/dihedron/snoop/openstack/oslo"

	//lint:ignore ST1001 we want to use the API in a fluent way with no clutter
	"github.com/dihedron/snoop/transform/chain"
	"github.com/dihedron/snoop/transform/transformers"
)

// Process is the command that reads message from RabbitMQ and processes them to
// output events to syslog; in the process, it may record the messages to a file
// if the --record flag is specified.
// If --playback is specified, the command reads messages from one or more files
// instead of RabbitMQ; if --dry-run is specified, the command simulates processing
// without actually writing events to syslog or acknowledging incoming messages to
// RabbitMQ.
type Process struct {
	base.ConnectedCommand
	// Playback indicates whether the command should read the incoming messages
	// from a file instead of connecting to RabbitMQ; either --connect or --playback
	// can be specified, but not both.
	Playback string `short:"p" long:"playback" description:"Whether the messages will be read from some input file(s)." optional:"yes"`
	// Record indicates the optional path to a file used for recording incoming
	// messages. This flag is only compatible with the
	Record *string `short:"r" long:"record" description:"The path to the file to write to (use '-' for STDOUT)." optional:"yes"`
	// Truncate is used to specify whether the output file (if --record refers
	// to a file on disk) should be truncated before writing to it.
	Truncate bool `short:"t" long:"truncate" description:"Whether the output file should be truncated or appended to (default)." optional:"yes"`
	// DryRun is used to specify whether the command should be run so that it
	// has no side effects, i.e. it simulates processing without actually writing
	// events to syslog and acknowledging incoming messages to RabbitMQ.
	DryRun bool `short:"d" long:"dry-run" description:"Whether to perform a dry run, i.e. simulating processing with no side effects." optional:"yes"`
}

// Execute is the real implementation of the Record command.
func (cmd *Process) Execute(args []string) error {
	var err error

	//
	// validate command line flags first
	//
	if cmd.Playback {
		// running from file, so the input files are mandatory
		if len(args) == 0 {
			slog.Error("no input files")
			return errors.New("no input files provided")
		}
		if cmd.Record != nil {
			slog.Warn("ignoring --record flag when running in playback mode")
			cmd.Record = nil
		}
	} else {
		// running from RabbitMQ, so the configuration file is mandatory
		if cmd.Configuration == nil {
			slog.Error("no configuration file provided")
			return errors.New("no configuration file provided")
		}
	}

	//
	// prepare the writer
	//
	var writer io.Writer = io.Discard
	if cmd.Record != nil {
		switch *cmd.Record {
		case "":
			// this should never happen
			slog.Error("no output file provided")
			return errors.New("no output file provided")
		case "-":
			slog.Info("writing to STDOUT")
			writer = os.Stdout
		default:
			slog.Info("writing to file", "path", *cmd.Record)
			flags := 0
			if cmd.Truncate {
				slog.Debug("opening output file in truncate mode", "path", *cmd.Record)
				flags = os.O_TRUNC | os.O_CREATE | os.O_WRONLY
			} else {
				slog.Debug("opening output file in append mode", "path", *cmd.Record)
				flags = os.O_APPEND | os.O_CREATE | os.O_WRONLY
			}
			file, err := os.OpenFile(*cmd.Record, flags, 0666)
			if err != nil {
				slog.Error("error opening recorder output file in append mode", "path", *cmd.Record, "mode", cmd.Truncate, "error", err)
				return errors.New("error openinig output file")
			}
			defer file.Close()
			writer = file
		}
	}
	slog.Debug("writer is ready", "type", fmt.Sprintf("%T", writer))

	//
	// if dry run, do not output to syslog nor ack the messages
	//
	if cmd.DryRun {
		slog.Info("running in dry-run mode")
		//syslog := &transformers.SysLogWriter{}
	}

	//
	// if in playback mode, retrieve the messages from file
	//
	if cmd.Playback {
		err = cmd.processFromFile(args)
	} else {
		err = cmd.processFromRabbitMQ()
	}
	/*
			// read configuration
			cfg, err := helpers.LoadConfiguration(cmd.Configuration)
			if err != nil {
				slog.Error("error loading configuration", "path", cmd.Configuration, "error", err)
				return err
			}
			slog.Info("configuration successfully loaded")

			// open the output stream
			stream, err := helpers.OpenOutputStream(cmd.Output, cmd.Truncate)
			if err != nil {
				slog.Error("error opening output stream", "stream", cmd.Output, "error", err)
				return err
			}
			defer (stream.(io.WriteCloser)).Close()
			slog.Info("output stream successfully opened")

			// create the source
			source, err := helpers.NewRabbitMQSource(cfg.RabbitMQ.Servers, cfg.RabbitMQ.Bindings, cfg.RabbitMQ.Queue, cfg.RabbitMQ.Client)
			if err != nil {
				slog.Error("error creating new RabbitMQ source", "error", err)
				return err
			}

			// 1. record the incoming messages to file
			recorder := filter.NewRecorder(stream, false)
			// 2. then unwrap them 2 times (AMQP->Oslo->OpenStack)
			amqpUnwrapper := message.NewAMQPMessageUnwrapper()
			osloUnwrapper := message.NewOsloMessageUnwrapper()
			ospUnwrapper := message.NewOpenStackMessageUnwrapper()
			// 3. then log KeyStone notifications to Syslog
			syslogger, err := syslogger.NewSysLogWriter(
				syslogger.WithApplicationName("brokerd"),
				syslogger.WithEnterpriseId("bancaditalia"),
				syslogger.WithProcessId(fmt.Sprintf("%d", os.Getpid())),
				syslogger.WithAcceptor(func(msg dataflow.Message) bool {
					slog.Debug("analysing message for inclusion into syslog...", "type", fmt.Sprintf("%T", msg))
					if m, ok := msg.(*message.OpenStackMessage); ok {
						slog.Debug("OpenStack notification", "type", fmt.Sprintf("%T", m.Message))
						if m, ok := m.Message.(*message.IdentityNotification); ok {
							slog.Info("sending message to syslog", "type", m.EventType)
							return true
						}
					}
					return false
				}),
			)
			if err != nil {
				slog.Fatal("error initialising syslogd writer")
			}
			// 4. then count messages
			counter := filter.NewCounter()

			p := pipeline.New(
				pipeline.WithSource(source),
				pipeline.WithFilters(
					amqpUnwrapper,
					recorder,
					osloUnwrapper,
					ospUnwrapper,
					syslogger,
					counter,
				),
				pipeline.WithErrorCallback(func(ctx context.Context, quit chan<- bool, filter string, err error) {
					slog.Errorf("callback called on filter %s", filter)
				}),
			)

			// get ready to handle signals (CTRL+C etc.) from user
			signals := make(chan os.Signal, 1)
			signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
			defer close(signals)

			// start the pipeline with the possibility to terminate
			// it via COntext cancellation
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			results, quit, err := p.Start(ctx)
			if err != nil {
				slog.Fatal("error opening pipeline")
			}
			defer p.Close()
		loop:
			for {
				select {
				case signal := <-signals:
					slog.Debugf("signal received: %v", signal)
					fmt.Printf("signal received: %v\n", signal)
					break loop
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
			slog.Infof("record command complete: %d messages recorded", counter.Count())
	*/
	return err
}

func (cmd *Process) processFromFile(args []string) error {
	var err error

	slog.Debug("playing back messages from recordings...", "files", args)

	ctx := context.Background()
	stopwatch := &transformers.StopWatch[string, notification.Notification]{}
	multicounter := &transformers.MultiCounter[notification.Notification, string]{}

	//acceptedEvents := []string{"compute.instance.shutdown.end", "compute.instance.shutdown.start"}

	unwrap := chain.Of7(
		stopwatch.Start(),
		transformers.StringToByteArray(),
		amqp.JSONToMessage(),
		oslo.MessageToOslo(false),
		notification.OsloToNotification(false),
		multicounter.Add(func(n notification.Notification) string { return n.Summary().EventType }),
		stopwatch.Stop(),
	)

	files := file.New()
	for line := range files.AllLinesContext(ctx, args...) {
		var n notification.Notification

		if n, err = unwrap(line); err != nil {
			slog.Error("error unwrapping line", "line", line, "error", err)
			continue
		}
		slog.Info("unwrapped line", "line", line, "output", n, "elapsed", stopwatch.Elapsed())

		err = cmd.processNotification(n)
		if err != nil {
			slog.Error("error processing notification", "error", err)
		}
	}

	stats, _ := multicounter.Count()
	cmd.PrintStatistics(stats)
	os.Stdout.Sync()
	return nil
}

func (cmd *Process) processFromRabbitMQ() error {
	// TODO
	return nil
}

func (cmd *Process) processNotification(_ notification.Notification) error {
	// TODO
	return nil
}

func (cmd *Process) PrintStatistics(stats map[string]int64) error {
	keys := make([]string, len(stats))
	i := 0
	for k := range stats {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%-50s: %d\n", k, stats[k])
	}
	return nil
}
