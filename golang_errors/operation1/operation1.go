package operation1

type Operation1 struct {
	Number1 int
	Number2 int
}

func (opt *Operation1) Add() int {
	return opt.Number1 + opt.Number2
}

func (opt *Operation1) Sub() int {
	return opt.Number1 - opt.Number2
}

func (opt *Operation1) Multiply() int {
	return opt.Number1 * opt.Number2
}

func (opt *Operation1) Divide() int {
	return opt.Number1 / opt.Number2
}
