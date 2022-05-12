package repository

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func marshall(v any, record *[]string) error {
	s := reflect.ValueOf(v).Elem()
	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		switch field.Kind() {
		case reflect.String:
			*record = append(*record, field.String())
		case reflect.Int:
			*record = append(*record, strconv.FormatInt(field.Int(), 10))
		case reflect.Uint:
			*record = append(*record, strconv.FormatUint(field.Uint(), 10))
		case reflect.Bool:
			*record = append(*record, strconv.FormatBool(field.Bool()))
		case reflect.Struct:
			si := field.Addr().Interface()
			structFields := []string{}
			marshall(si, &structFields)
			*record = append(*record, fmt.Sprintf("%v", strings.Join(structFields, "$")))
		case reflect.Slice:
			structType := s.Type().Field(i).Type.Elem()
			kind := structType.Kind()
			switch kind {
			case reflect.String:
				values, ok := field.Interface().([]string)
				if !ok {
					return errors.New("unable to get slice of strings")
				}
				*record = append(*record, fmt.Sprintf("%v", strings.Join(values, ",")))
			case reflect.Struct:
				structRecords := []string{}
				for i := 0; i < field.Len(); i++ {
					value := field.Index(i)
					recordFields := []string{}
					marshall(value.Addr().Interface(), &recordFields)
					structRecords = append(structRecords, fmt.Sprintf("%v", strings.Join(recordFields, "$")))
				}
				*record = append(*record, fmt.Sprintf("%v", strings.Join(structRecords, ",")))
			}
		}
	}

	return nil
}
