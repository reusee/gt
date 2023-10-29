package gt

import "context"

type Generator[T any] func(
	ctx context.Context,
	control *Control,
	next func(T) error,
) error

type Control struct {
	Cancel bool
	Spill  bool
}
