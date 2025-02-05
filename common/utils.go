package common

import (
	"errors"
	"fmt"
	"reflect"
)

func StructToMap(obj interface{}) (map[string]interface{}, error) {

	if obj == nil {
		return nil, errors.New("input is nil")
	}

	result := make(map[string]interface{})
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)

	if val.Kind() != reflect.Struct {
		return nil, errors.New("input is not a struct")
	}

	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil, errors.New("input is a nil pointer")
		}
		val = val.Elem()
		typ = typ.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		if !field.CanInterface() {
			continue
		}

		// 检查字段是否是结构体
		if field.Kind() == reflect.Struct {
			nestedMap, err := StructToMap(field.Interface())
			if err != nil {
				return nil, fmt.Errorf("failed to convert nested struct field '%s': %w", fieldType.Name, err)
			}
			result[fieldType.Name] = nestedMap
		} else {
			result[fieldType.Name] = field.Interface()
		}
	}

	return result, nil
}
