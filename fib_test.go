package gt

import (
	"context"
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	TestGenerator(t, func() Generator[int] {
		return NewFib()
	})
}

func TestFibSeq(t *testing.T) {
	ctx := context.Background()
	gen := NewFib()
	var seq []int
	for i := 0; i < 10; i++ {
		if err := gen(ctx, nil, func(i int) error {
			seq = append(seq, i)
			return nil
		}); err != nil {
			t.Fatal(err)
		}
	}
	if fmt.Sprintf("%v", seq) != `[1 1 2 3 5 8 13 21 34 55]` {
		t.Fatal()
	}
}
