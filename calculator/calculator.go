package calculator

import "errors"

var (
	ErrDivisionByZero = errors.New("Divide: cannot divide by 0")
)

func Add(lhs, rhs float64) float64 {
	return lhs + rhs
}

func Substract(lhs, rhs float64) float64 {
	return lhs - rhs
}

func Multiply(lhs, rhs float64) float64 {
	return lhs * rhs
}

func Divide(lhs, rhs float64) (float64, error) {
	if rhs == 0 {
		return 0, ErrDivisionByZero
	}
	return lhs / rhs, nil
}
