package command

import (
	"github.com/dihedron/snoop/command/check"
	"github.com/dihedron/snoop/command/playback"
	"github.com/dihedron/snoop/command/record"
	"github.com/dihedron/snoop/command/version"
)

// Commands is the set of root command groups.
type Commands struct {
	// // Inspect provides facilities to inspect AMQP, Oslo and OpenStack messages.
	// Inspect inspect.Inspect `command:"inspect" alias:"i" description:"Interactively inspect the contents of one or more files."`
	// // Aggregate provides a way to split into multiple files a single log.
	// Aggregate aggregate.Aggregate `command:"aggregate" alias:"A" description:"Split a large log into multiple files according to some pattern."`
	// // Split provides a way to split into multiple files a single log.
	// Split split.Split `command:"split" alias:"s" description:"Split a large log into multiple files according to some pattern."`

	// // Admin provides a set of cluster administation tools.
	// Admin admin.Admin `command:"administration" alias:"admin" alias:"a" description:"Run administration command against the cluster."`

	// // AdminHelp prints information about the cluster administration tools.
	// AdminHelp admin.AdminHelp `command:"adminhelp" alias:"ah" alias:"h" description:"Print help abount administration commands."`

	// // Store manages data in the cluster's K/V store.
	// Store store.Store `command:"store" alias:"s" description:"Manage data in the cluster K/V store."`

	// Process runs the snoop command against the RabbitMS server as specified in the configuration.
	//Process process.Process `command:"process" alias:"p" description:"Run the snoop utility against RabbitMQ."`

	// Check checks the connectivity to RabbitMQ.
	Check check.Check `command:"check" alias:"c" description:"Try to connect to the RabbitMQ server."`

	// Record reads messages from RabbitMQ and outputs them (to disk or STDOUT).
	Record record.Record `command:"record" alias:"r" description:"Read messages from RabbitMQ and output them (to disk or STDOUT)."`

	// Playback reads messages from a text file and outputs them (to disk or STDOUT).
	Playback playback.Playback `command:"playback" alias:"p" description:"Plays messages back from a recording on disk."`

	// Version prints brokerd version information and exits.
	//lint:ignore SA5008 commands can have multiple aliases
	Version version.Version `command:"version" alias:"ver" alias:"v" description:"Show the command version and exit."`
}
