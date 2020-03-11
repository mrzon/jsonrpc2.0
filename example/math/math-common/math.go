package math_common

type MathService struct {
	Add       func(A int, B int) int `jsonrpc:"add"`
	Substract func(A int, B int) int
	Multiply  func(A int, B int) int
	Modulo    func(A int, B int) int
}
