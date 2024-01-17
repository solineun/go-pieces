package csvmarshaller

import (
	"errors"
	"reflect"
)

func Unmarshall(data [][]string, v interface{}) error {
	sliceValPtr := reflect.ValueOf(v)
	if sliceValPtr.Kind() != reflect.Ptr {
		return errors.New("must be a pointer to a slice of structs")
	}
	sliceVal := sliceValPtr.Elem()
	if sliceVal.Kind() != reflect.Slice {
		return errors.New("must be a pointer to a slice of structs")
	}
	structType := sliceVal.Type().Elem()
	if structType.Kind() != reflect.Struct {
		return errors.New("must be a pointer to a slice of structs")
	}
}