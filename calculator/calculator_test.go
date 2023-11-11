package calculator_test

import (
	"testing"

	"github.com/mobamoh/codewars/calculator"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		lhs, rhs, want float64
	}{
		{2, 2, 4},
		{3, 99, 102},
		{-67, 100, 33},
		{-3, -100, -103},
	}

	for _, tc := range tcs {
		got := calculator.Add(tc.lhs, tc.rhs)
		if tc.want != got {
			t.Errorf("Add(%.1f,%.1f) want %.1f, got %.1f", tc.lhs, tc.rhs, tc.want, got)
		}
	}

}

func TestSubstract(t *testing.T) {
	t.Parallel()
	var lhs, rhs float64 = 7, 2
	want := float64(5)
	got := calculator.Substract(lhs, rhs)
	if want != got {
		t.Errorf("Substract(%.1f,%.1f) want %.1f, got %.1f", lhs, rhs, want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var lhs, rhs float64 = 8, 2
	want := float64(16)
	got := calculator.Multiply(lhs, rhs)
	if want != got {
		t.Errorf("Multiply(%.1f,%.1f) want %.1f, got %.1f", lhs, rhs, want, got)
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		lhs, rhs, want float64
		err            error
	}{
		{8, 4, 2, nil},
		{3, 2, 1.5, nil},
		{77, 0, 0, calculator.ErrDivisionByZero},
		{-66, -6, 11, nil},
	}

	for _, tc := range tcs {
		got, err := calculator.Divide(tc.lhs, tc.rhs)
		if err != tc.err {
			t.Fatalf("Expected error got nil! %v", err)
		}
		if got != tc.want {
			t.Errorf("Divide(%.1f,%.1f) got: %.1f, want: %.1f", tc.lhs, tc.rhs, got, tc.want)
		}
	}
}

func TestCalculatorMethods(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		fn             func(float64, float64) float64
		lhs, rhs, want float64
	}{
		{calculator.Add, 2, 2, 4},
		{calculator.Multiply, 12, 99, 1188},
		{calculator.Substract, -67, 100, -167},
		{calculator.Multiply, -3, -100, 300},
	}

	for _, tc := range tcs {
		got := tc.fn(tc.lhs, tc.rhs)
		if tc.want != got {
			t.Errorf("(%.1f,%.1f) want %.1f, got %.1f", tc.lhs, tc.rhs, tc.want, got)
		}
	}

}
