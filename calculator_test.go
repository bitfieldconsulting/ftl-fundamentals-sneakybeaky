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

	for _, test := range tcs {
		t.Run(fmt.Sprintf(test.name), func(t *testing.T) {
			got := calculator.Add(test.a, test.b)
			if test.want != got {
				t.Errorf("want %f, got %f", test.want, got)
			}

		})
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		a, b float64
		want float64
	}{
		{name: "Subtracting two positive numbers giving a positive", a: 4, b: 2, want: 2},
		{name: "Subtracting two positive numbers giving zero", a: 2, b: 2, want: 0},
		{name: "Subtracting two positive numbers giving a negative number", a: 2, b: 4, want: -2},
	}

	for _, test := range tcs {
		t.Run(fmt.Sprintf(test.name), func(t *testing.T) {
			got := calculator.Subtract(test.a, test.b)
			if test.want != got {
				t.Errorf("want %f, got %f", test.want, got)
			}

		})
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	tcs := map[string]struct {
		a, b float64
		want float64
	}{
		"Product of 2 positives is a positive":   {a: 2, b: 5, want: 10},
		"Product of a positive and zero is zero": {a: 2, b: 0, want: 0},
		"Product of a negative and zero is zero": {a: -10, b: 0, want: 0},
	}

	for name, test := range tcs {
		t.Run(fmt.Sprintf(name), func(t *testing.T) {
			got := calculator.Multiply(test.a, test.b)
			if test.want != got {
				t.Errorf("want %f, got %f", test.want, got)
			}

		})
	}
}
