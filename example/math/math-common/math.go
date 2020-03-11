package math_common

type MathService struct {
	Add       func(A int, B int) int        `jsonrpc:"add"`
	Add3      func(A int, B int, C int) int `jsonrpc:"add,3"`
	Substract func(A int, B int) int
	Multiply  func(A int, B int) int
	Modulo    func(A int, B int) int
}
