package types

type T0 interface {
	Add() int
}

type T1 struct {
}

func (t *T1) Add() int {
	return 1
}

var _ T0 = new(T1) // 等于 T0 相当于接口T0的变量形式

func NewT0() T0 {
	t := new(T1)
	return t
}
func NewT1() T1 {
	t := new(T1)
	return *t
}
