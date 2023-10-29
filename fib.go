package gt

import (
	"context"
)

func NewFib() Generator[int] {
	// states
	a := 0
	b := 1
	var finalErr error

	return func(ctx context.Context, control *Control, next func(i int) error) (err error) {
		// final error
		defer func() {
			if finalErr == nil && err != nil {
				finalErr = err
			}
		}()
		if finalErr != nil {
			return finalErr
		}

		// context error
		if err := ctx.Err(); err != nil {
			return err
		}

		// cancel
		if control != nil && control.Cancel {
			return context.Canceled
		}

		// next
		if next != nil {
			if err := next(b); err != nil {
				return err
			}
			a, b = b, a+b
		}

		return nil
	}
}
