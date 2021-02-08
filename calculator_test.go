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

	for _, tc := range tcs {
		t.Run(fmt.Sprintf(tc.name), func(t *testing.T) {
			got := calculator.Add(tc.a, tc.b)
			if tc.want != got {
				t.Errorf("want %f, got %f", tc.want, got)
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
		{name: "Subtracting two positive fractional numbers giving a positive whole number", a: 5.5, b: 1.5, want: 4},
		{name: "Subtracting two positive fractional numbers giving a positive fractional number", a: 5.5, b: 2.4, want: 3.1},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf(tc.name), func(t *testing.T) {
			got := calculator.Subtract(tc.a, tc.b)
			if tc.want != got {
				t.Errorf("want %f, got %f", tc.want, got)
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

	for name, tc := range tcs {
		t.Run(fmt.Sprintf(name), func(t *testing.T) {
			got := calculator.Multiply(tc.a, tc.b)
			if tc.want != got {
				t.Errorf("want %f, got %f", tc.want, got)
			}

		})
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name        string
		a, b        float64
		errExpected bool
		want        float64
	}{
		{name: "Dividing two positive numbers giving a positive", a: 8, b: 2, want: 4},
		{name: "Dividing by zero giving an error", a: 10, b: 0, errExpected: true},
		{name: "Dividing two negative numbers giving a positive", a: -10, b: -2, want: 5},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf(tc.name), func(t *testing.T) {
			got, err := calculator.Divide(tc.a, tc.b)

			if tc.errExpected {
				if err == nil {
					t.Error("expected an error but got none")
				}

			} else {

				if err != nil {
					t.Errorf("got an error but wasn't expecting one : %v", err)
					return
				}

				if tc.want != got {
					t.Errorf("want %f, got %f", tc.want, got)
				}
			}

		})
	}
}
