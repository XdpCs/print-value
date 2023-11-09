package print_value

// @Title        print.go
// @Description
// @Create       XdpCs 2023-11-09 10:43
// @Update       XdpCs 2023-11-09 10:43

import (
	"fmt"
	"reflect"
)

type PrintValue struct{}

func (p *PrintValue) Print(v interface{}) string {
	return print(v)
}

func Print(v interface{}) string {
	return print(v)
}

func print(v interface{}) string {
	value := reflect.ValueOf(v)
	return printValue(value)
}

func printValue(v reflect.Value) string {
	var result string
	switch v.Kind() {
	case reflect.Invalid:
		result = "nil"
	case reflect.Ptr:
		if !v.IsNil() {
			result = printValue(v.Elem())
		} else {
			result = "nil"
		}
	case reflect.Struct:
		result += v.Type().Name() + "{"
		for i := 0; i < v.NumField(); i++ {
			result += v.Type().Field(i).Name + ":"
			result += printValue(v.Field(i))
			if i != v.NumField()-1 {
				result += ","
			}
		}
		result += "}"
	case reflect.Slice, reflect.Array:
		result += "["
		for i := 0; i < v.Len(); i++ {
			result += printValue(v.Index(i))
			if i != v.Len()-1 {
				result += ","
			}
		}
		result += "]"
	case reflect.Map:
		result += "map["
		for i, key := range v.MapKeys() {
			result += printValue(key) + ":" + printValue(v.MapIndex(key))
			if i != len(v.MapKeys())-1 {
				result += ","
			}
		}
		result += "]"
	default:
		result = fmt.Sprintf("%+v", v)
	}
	return result
}
