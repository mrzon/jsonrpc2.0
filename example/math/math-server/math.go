package math_server

import math_common "github.com/mrzon/jsonrpc2.0/example/math/math-common"

func NewMathServiceImpl() math_common.MathService {
	mImpl := math_common.MathService{
		Add: func(A int, B int) int {
			//time.Sleep(3 * time.Second)
			return (A + B)
		},
		Add3: func(A int, B int, C int) int {
			//time.Sleep(3 * time.Second)
			return (A + B + C)
		},
		Substract: func(A int, B int) int {
			return A - B
		},
		Multiply: func(A int, B int) int {
			return A * B
		},
		Modulo: func(A int, B int) int {
			return A % B
		},
	}
	return mImpl
}
