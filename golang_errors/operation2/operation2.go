package operaion2

import (
	"errors"
)

type Operation2 struct {
	number1 int
	number2 int
}

func NewOperation2(a, b int) *Operation2 {
	return &Operation2{
		number1: a,
		number2: b,
	}
}

func (opt *Operation2) Add() int {
	return opt.number1 + opt.number2
}

func (opt *Operation2) Sub() int {
	return opt.number1 - opt.number2
}

func (opt *Operation2) Multiply() int {
	return opt.number1 * opt.number2
}

func (opt *Operation2) Divide() (int, error) {
	if opt.number2 == 0 {
		err := errors.New("integer divide by zero")
		return 0, err
	}
	return opt.number1 / opt.number2, nil
}
