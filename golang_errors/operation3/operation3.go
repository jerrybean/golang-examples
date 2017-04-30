package operation3

import (
	"errors"
)

type Operation3 struct {
	number1 int
	number2 int
	err     error
}

var (
	ErrDivideByZero = errors.New("divide by zero")
	ErrSubLess      = errors.New("minuend must greater or equal than reduction")
)

func NewOperation3(a, b int) *Operation3 {
	var err error
	if a < b {
		err = ErrSubLess
	}
	return &Operation3{
		number1: a,
		number2: b,
		err:     err,
	}
}

func (opt *Operation3) Err() error {
	return opt.err
}

func (opt *Operation3) Add() int {
	return opt.number1 + opt.number2
}

func (opt *Operation3) Sub() int {
	return opt.number1 - opt.number2
}

func (opt *Operation3) Multiply() int {
	return opt.number1 * opt.number2
}

func (opt *Operation3) Divide() int {
	if opt.number2 == 0 {
		opt.err = errors.New("divide by zero")
		return 0
	}
	return opt.number1 / opt.number2
}
