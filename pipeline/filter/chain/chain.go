package chain

import "github.com/dihedron/snoop/pipeline"

// New creates a chain of Handlers that will be executed one after the
// other, until one handler returns and error or the chain is completely
// applied.
func New[T any](handlers ...pipeline.Handler[T]) pipeline.Handler[T] {
	return func(value T) (T, error) {
		var err error
		for _, handler := range handlers {
			value, err = handler(value)
			if err != nil {
				return value, err
			}
		}
		return value, nil
	}
}
