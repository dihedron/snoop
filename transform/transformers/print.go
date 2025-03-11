package transformers

import (
	"fmt"
	"io"
	"os"

	"github.com/dihedron/snoop/transform/chain"
)

func Print[T any](stream io.Writer) chain.X[T, T] {
	var s io.Writer = os.Stdout
	if stream != nil {
		s = stream
	}
	return func(value T) (T, error) {
		fmt.Fprintf(s, "%v", value)
		return value, nil
	}
}
