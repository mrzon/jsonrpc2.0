package util

import (
	"reflect"
	"encoding/json"
	"unicode"
	"log"
)

func ToBoolPtr(val bool) *bool{
	return &val;
}

func ToIntPtr(val int) *int{
	return &val;
}

func ToCamel(val string) string {
	a := []rune(val)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}


func GetVal(typeOf reflect.Type, val interface{}) (value reflect.Value) {
	if val == nil {
		value = reflect.Zero(typeOf)
	} else if (typeOf.Kind() == reflect.Ptr ){
		temp := GetVal(typeOf.Elem(), val)
		if !temp.CanAddr() {
			log.Println("Primitive return can not be a pointer")
		} else {
			value = temp.Addr()
		}
	} else if typeOf.Kind() == reflect.Int {
		value = reflect.ValueOf(int(val.(float64)))
	} else if typeOf.Kind() == reflect.Int64 {
		value = reflect.ValueOf(int64(val.(float64)))
	} else if typeOf.Kind() == reflect.Float32 {
		value = reflect.ValueOf(float32(val.(float64)))
	} else if typeOf.Kind() == reflect.Struct {
		intPtr := reflect.New(typeOf).Interface()
		paramStr, err := json.Marshal(val)
		if err == nil {
			json.Unmarshal(paramStr, &intPtr)
			return reflect.ValueOf(intPtr).Elem()
		}
	} else {
		value = reflect.ValueOf(val)
	}
	return value
}