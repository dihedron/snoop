package base

import (
	"log/slog"
	"os"
	"runtime"
	"runtime/pprof"
)

type Command struct {
	// CPUProfile sets the (optional) path of the file for CPU profiling info.
	CPUProfile string `short:"C" long:"cpu-profile" description:"The (optional) path where the CPU profiler will store its data." optional:"yes"`
	// MemProfile sets the (optional) path of the file for memory profiling info.
	MemProfile string `short:"M" long:"mem-profile" description:"The (optional) path where the memory profiler will store its data." optional:"yes"`
}

type ConfiguredCommand struct {
	Command
	// Configuration contains the path to the (optional) configuration file to use; if
	// no value is provided (neither on the command line nor in the environment via the
	// SNOOP_CONFIGURATION variable), the application will look for a viable configuration
	// file named XXX under a few well-known paths: /etc, the current directory etc.
	Configuration *string `short:"c" long:"configuration" description:"The path to the configuration file." optional:"yes" env:"SNOOP_CONFIGURATION"`
}

func (cmd *Command) ProfileCPU() *Closer {
	var f *os.File
	if cmd.CPUProfile != "" {
		var err error
		f, err = os.Create(cmd.CPUProfile)
		if err != nil {
			slog.Error("could not create CPU profile", "file", cmd.CPUProfile, "error", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			slog.Error("could not start CPU profiler", "error", err)
		}
	}
	return &Closer{
		file: f,
	}
}

func (cmd *Command) ProfileMemory() {
	if cmd.MemProfile != "" {
		f, err := os.Create(cmd.MemProfile)
		if err != nil {
			slog.Error("could not create memory profile", "file", cmd.MemProfile, "error", err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			slog.Error("could not write memory profile", "error", err)
		}
	}
}

type Closer struct {
	file *os.File
}

func (c *Closer) Close() {
	if c.file != nil {
		pprof.StopCPUProfile()
		c.file.Close()
	}
}
