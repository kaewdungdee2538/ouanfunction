package sturcture

import (
	"fmt"
	"reflect"
	"strings"
)

func MapToStruct(m map[string]interface{}, s interface{}) error {
	structValue := reflect.ValueOf(s).Elem()
	structType := structValue.Type()

	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		fieldType := structType.Field(i)

		fieldNameLowerCase := strings.ToLower(fieldType.Name)

		if value, ok := m[fieldNameLowerCase]; ok {
			if reflect.TypeOf(value) == field.Type() {
				field.Set(reflect.ValueOf(value))
			} else {
				return fmt.Errorf("Type mismatch for field '%s'", fieldType.Name)
			}
		}
	}

	return nil
}
