package helper

// buat handle error update field

import (
	"reflect"
)

func GetFieldValue(data interface{}, jsonTag string) interface{} {
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("json")
		if tag == jsonTag {
			return val.Field(i).Interface()
		}
	}
	return nil
}
