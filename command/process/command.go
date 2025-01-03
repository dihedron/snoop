package process

import (
	"log/slog"

	"github.com/dihedron/snoop/command/base"
)

// Record is the command that reads message from RabbitMQ and dumps them
// out to standard output or to a file.
// ./brokerd record --configuration=tests/rabbitmq/brokerd.yaml --output=20220818.amqp.messages
type Process struct {
	base.ConfiguredCommand
	// Output is the path to the file to write to; the default value is
	// "<stdout>", which means that the messages will be dumped to STDOUT.
	Output string `short:"o" long:"output" description:"The path to the file to write to, or stdout." required:"yes" default:"<stdout>"`
	// Truncate is used to specify whether the output file (is the path
	// refers to a file on disk) should be truncated before writing to it.
	Truncate bool `short:"t" long:"truncate" description:"Whether the output file should be truncated or appended to (default)." optional:"yes"`
	// TODO: remove once checked
	// Configuration string `short:"c" long:"configuration" description:"The path to the configuration file." optional:"yes"`
}

// Execute is the real implementation of the Record command.
func (cmd *Process) Execute(args []string) error {
	slog.Debug("draining and recording messages from RabbitMQ")

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
	return nil
}

// // Execute is the real implementation of the Record command.
// func (cmd *Record) Execute_Ingestor(args []string) error {
// 	slog.Debug("draining and recording messages from RabbitMQ")

// 	// read configuration
// 	cfg, err := helpers.LoadConfiguration(cmd.Configuration)
// 	if err != nil {
// 		slog.Errorf("error loading configuration from %q", cmd.Configuration)
// 		return err
// 	}
// 	slog.Info("configuration successfully loaded")

// 	// open the output stream
// 	stream, err := helpers.OpenOutputStream(cmd.Output, cmd.Truncate)
// 	if err != nil {
// 		slog.Errorf("error opening output stream: %s", cmd.Output)
// 		return err
// 	}
// 	defer (stream.(io.WriteCloser)).Close()
// 	slog.Info("output stream successfully opened")

// 	ingestor, err := helpers.NewRabbitMQIngestor(cfg.RabbitMQ.Servers, cfg.RabbitMQ.Bindings, cfg.RabbitMQ.Queue, cfg.RabbitMQ.Client)

// 	if err != nil {
// 		slog.Error("error creating new RabbitMQ ingestor")
// 		return err
// 	}
// 	slog.Info("input ingestor successfully opened")

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	// open the RabbitMQ channel
// 	deliveries, err := ingestor.Ingest(ctx)
// 	if err != nil {
// 		slog.Error("error opening queue to RabbitMQ server")
// 		return err
// 	}
// 	slog.Info("RabbitMQ channes ready for processing")

// 	// get ready to handle signals (CTRL+C etc.) from user
// 	signals := make(chan os.Signal, 1)
// 	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
// 	defer close(signals)

// loop:
// 	for {
// 		select {
// 		case sig := <-signals:
// 			fmt.Printf("signal received: %v\n", sig)
// 			break loop
// 		case delivery := <-deliveries:
// 			if delivery, ok := delivery.(amqp.Delivery); ok {
// 				amqpMsg, err := message.NewAMQPDelivery(&delivery)
// 				if err != nil {
// 					slog.Error("error reading AMQP message")
// 					delivery.Ack(false)
// 					continue
// 				}
// 				fmt.Fprintf(stream, "%v\n", amqpMsg)
// 				delivery.Ack(true)
// 			}
// 		}
// 	}
// 	return nil
// }
