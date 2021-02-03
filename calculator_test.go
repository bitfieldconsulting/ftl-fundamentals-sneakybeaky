package calculator_test

import (
	"calculator"
	"fmt"
	"testing"
)

type testcase struct {
	a    float64
	b    float64
	want float64
}

func TestAdd(t *testing.T) {
	t.Parallel()

	tests := []testcase{
		{a: 2, b: 2, want: 4},
		{a: -1, b: 1, want: 0},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			got := calculator.Add(test.a, test.b)
			if test.want != got {
				t.Errorf("want %f, got %f", test.want, got)
			}

		})
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 10
	got := calculator.Multiply(2, 5)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
