package test

import (
	"log/slog"
	"testing"

	"github.com/neilotoole/slogt"
)

type Format uint8

const (
	JSON Format = iota
	Text
)

func Setup(t *testing.T, format ...Format) {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	if format != nil && format[0] == JSON {
		slog.SetDefault(slogt.New(t, slogt.JSON()))
	} else {
		slog.SetDefault(slogt.New(t, slogt.Text()))
	}
}
