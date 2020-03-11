package server

import (
	"reflect"
	"testing"
)

func Test_getMethodFromTagAndParam(t *testing.T) {
	type A struct {
		A1 func()               `jsonrpc:"a1,0"`
		A2 func(p1 int)         `jsonrpc:"a1,1"`
		A3 func(p1 int, p2 int) `jsonrpc:"a1,2"`
	}
	a := A{}
	aVal := reflect.ValueOf(a)
	type args struct {
		service      interface{}
		serviceValue reflect.Value
		tagName      string
		params       []interface{}
	}

	var AA args = args{
		service:      a,
		serviceValue: aVal,
		tagName:      "a1",
		params:       []interface{}{1, 2},
	}

	tests := []struct {
		name string
		args args
		want string
	}{{
		name: "get jsonrpc",
		args: AA,
		want: "A3",
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMethodNameFromTagAndParam(tt.args.service, tt.args.serviceValue, tt.args.tagName, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMethodFromTagAndParam() = %v, want %v", got, tt.want)
			}
		})
	}
}
