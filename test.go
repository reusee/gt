package gt

import (
	"context"
	"errors"
	"io"
	"testing"
)

func TestGenerator[T any](
	t *testing.T,
	newInstance func() Generator[T],
) {

	// canceled context
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	gen := newInstance()
	err := gen(ctx, nil, nil)
	if !errors.Is(err, context.Canceled) {
		t.Fatal()
	}
	err = gen(ctx, nil, nil)
	if !errors.Is(err, context.Canceled) {
		t.Fatal()
	}

	// control cancel
	ctx = context.Background()
	gen = newInstance()
	err = gen(ctx, &Control{
		Cancel: true,
	}, nil)
	if !errors.Is(err, context.Canceled) {
		t.Fatal()
	}
	err = gen(ctx, &Control{
		Cancel: true,
	}, nil)
	if !errors.Is(err, context.Canceled) {
		t.Fatal()
	}

	// next error
	ctx = context.Background()
	gen = newInstance()
	err = gen(ctx, nil, func(_ T) error {
		return io.ErrUnexpectedEOF
	})
	if !errors.Is(err, io.ErrUnexpectedEOF) {
		t.Fatal()
	}
	err = gen(ctx, nil, nil)
	if !errors.Is(err, io.ErrUnexpectedEOF) {
		t.Fatal()
	}

}
