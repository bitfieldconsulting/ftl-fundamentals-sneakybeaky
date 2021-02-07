package calculator_test

import (
	"calculator"
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		a, b float64
		want float64
	}{
		{name: "Sum of 2 positives is a positive", a: 2, b: 2, want: 4},
		{name: "A negative and a positive that sum to zero", a: -1, b: 1, want: 0},
		{name: "A negative and a negative that sum to a negative", a: -1, b: -4, want: -5},
	}

	for i, test := range tcs {
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
