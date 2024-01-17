package csvmarshaller

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)


func Marshall(v interface{}) ([][]string, error) {
	sliceVal := reflect.ValueOf(v)
	if sliceVal.Kind() != reflect.Slice {
		return nil, errors.New("must be a slice of structs")
	}
	structType := sliceVal.Type().Elem()
	if structType.Kind() != reflect.Struct {
		return nil, errors.New("must be a slice of structs")
	}
	var out [][]string
	header := marshallHeader(structType)
	out = append(out, header)
	for i := 0; i < sliceVal.Len(); i++ {
		row, err := marshallOne(sliceVal.Index(i))
		if err != nil {
			return nil, err
		}
		out = append(out, row)
	}
	return out, nil
}

func marshallHeader(st reflect.Type) []string {
	var row []string
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if curTag, ok := field.Tag.Lookup("csv"); ok {
			row = append(row, curTag)
		}
	}
	return row
}

func marshallOne(sv reflect.Value) ([]string, error) {
	var row []string
	svt := sv.Type()
	for i := 0; i < sv.NumField(); i++ {
		fieldVal := sv.Field(i)
		if _, ok := svt.Field(i).Tag.Lookup("csv"); !ok {
			continue
		}
		switch fieldVal.Kind() {
		case reflect.Int:
			row = append(row, strconv.FormatInt(fieldVal.Int(), 10))
		case reflect.String:
			row = append(row, fieldVal.String())
		case reflect.Bool:
			row = append(row, strconv.FormatBool(fieldVal.Bool()))
		default:
			return nil, fmt.Errorf("cannot handle field of kind %v", fieldVal.Kind())
		}
	}
	return row, nil
}