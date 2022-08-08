package parse

import "reflect"

func StructToMap(s interface{}) map[interface{}]interface{} {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	m := make(map[interface{}]interface{})
	for k := 0; k < t.NumField(); k++ {
		name := t.Field(k).Name
		value := v.Field(k).Interface()
		m[name] = value
	}
	return m
}
