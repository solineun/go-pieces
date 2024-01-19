package csvmarshaller

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func Unmarshall(data [][]string, v interface{}) error {
	if len(data) == 0 {
		return nil
	}
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
	header := data[0]
	namePos := make(map[string]int, len(header))
	for k, v := range header {
		namePos[v] = k
	}
	for _, row := range data[1:] {
		newVal := reflect.New(structType).Elem()
		err := unmarshallOne(row, namePos, newVal)
		if err != nil {
			return err
		}
		sliceVal.Set(reflect.Append(sliceVal, newVal))
	}	
	return nil
}

func unmarshallOne(row []string, namePos map[string]int, vv reflect.Value) error {
	vt := vv.Type()
	for i := 0; i < vv.NumField(); i++ {
		typeField := vt.Field(i)
		pos, ok := namePos[typeField.Tag.Get("csv")]
		if !ok {
			continue
		}
		val := row[pos]
		field := vv.Field(i)
		switch field.Kind() {
		case reflect.Int:
			i, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return err
			}
			field.SetInt(i)
		case reflect.String:
			field.SetString(val)
		case reflect.Bool:
			b, err := strconv.ParseBool(val)
			if err != nil {
				return err
			}
			field.SetBool(b)
		default:
			return fmt.Errorf("cannot handle field of kind %v", field.Kind())
		}
	}
	return nil
}