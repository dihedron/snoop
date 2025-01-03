# snoop

Snoop is a tool to collect and analyse OpenStack messages by sniffing them on the internal RabbitMQ bus.

## Subcommands

Snoop has the following subcommands:

`snoop drain`: connects to the cluster and downloads all messages, processing them; it acceps the `--record` flag to record all messages to disk and the `--process` to apply a set of instructions to the input stream.

`snoop replay`: replays a recorded input stream from a file; it accepts the `--process` flag to apply a set of instructions to the input stream.

`snoop inspect`: allows to load a recording and inspect it record by record.