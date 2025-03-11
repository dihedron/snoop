package transformers

import (
	"github.com/dihedron/snoop/transform/chain"
	"github.com/schollz/progressbar/v3"
)

// Counter holds the state needed to count items across multiple
// invocations of the chain. If Every is specified and more than
// 0, it will print a dot to STDOUT every count%Every values
type ProgressBar[S any, T any] struct {
	pb *progressbar.ProgressBar
}

// Init initialises the progress bar; set max to -1 to have a
// spinner.
func (p *ProgressBar[S, T]) Reset(max int64, description ...string) chain.X[T, T] {
	if p.pb == nil {
		p.pb = progressbar.Default(max, description...)
	}
	return func(value T) (T, error) {
		return value, nil
	}
}

// Add the given amount to the progress bar; if no amount is
// specified, it defaults to 1. This filter does not affect
// the value flowing through.
func (p *ProgressBar[S, T]) Add(amount ...int) chain.X[S, S] {
	a := 0
	if len(amount) == 1 {
		a = amount[0]
	}
	return func(value S) (S, error) {
		p.pb.Add(a)
		return value, nil
	}
}
