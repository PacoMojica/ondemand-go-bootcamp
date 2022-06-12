package repository

import (
	"reflect"
	"strconv"
	"strings"
)

func setString(field reflect.Value, value string) {
	field.SetString(value)
}

func setInt(field reflect.Value, value string) error {
	pi, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return err
	}
	field.SetInt(pi)
	return nil
}

func setUint(field reflect.Value, value string) error {
	pu, err := strconv.ParseUint(value, 10, 0)
	if err != nil {
		return err
	}
	field.SetUint(pu)
	return nil
}

func setBool(field reflect.Value, value string) error {
	pb, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	field.SetBool(pb)
	return nil
}

func setStringSlice(field reflect.Value, values []string) {
	newSlice := reflect.AppendSlice(field, reflect.ValueOf(values))
	field.Set(newSlice)
}

func packStructValues(value string, v any) []string {
	values := strings.Split(value, "$")
	s := reflect.ValueOf(v).Elem()
	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		if field.Kind() == reflect.Struct {
			end := i + field.NumField()
			structValues := strings.Join(values[i:end], "$")
			rest := values[end:]
			values = append(values[0:i], structValues)
			values = append(values, rest...)
		}
	}

	return values
}

func setStructArray(field reflect.Value, structType reflect.Type, values []string) error {
	newSlice := reflect.MakeSlice(field.Type(), 0, len(values))
	for _, value := range values {
		newStruct := reflect.New(structType).Elem()
		v := newStruct.Addr().Interface()
		structValues := packStructValues(value, v)
		unmarshall(structValues, v)
		newSlice = reflect.Append(newSlice, newStruct)
	}
	field.Set(newSlice)
	return nil
}

func unmarshall(record []string, v any) error {
	s := reflect.ValueOf(v).Elem()
	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		value := record[i]

		switch field.Kind() {
		case reflect.String:
			setString(field, value)
		case reflect.Int:
			if err := setInt(field, value); err != nil {
				return err
			}
		case reflect.Uint:
			if err := setUint(field, value); err != nil {
				return err
			}
		case reflect.Bool:
			if err := setBool(field, value); err != nil {
				return err
			}
		case reflect.Struct:
			si := field.Addr().Interface()
			unmarshall(strings.Split(value, "$"), si)
		case reflect.Slice:
			structType := s.Type().Field(i).Type.Elem()
			kind := structType.Kind()
			values := strings.Split(value, ",")
			switch kind {
			case reflect.String:
				setStringSlice(field, values)
			case reflect.Struct:
				if err := setStructArray(field, structType, values); err != nil {
					return err
				}
			}
		}
	}

	return nil

}
