package print_value

// @Title        print.go
// @Description
// @Create       XdpCs 2023-11-09 10:43
// @Update       XdpCs 2023-11-09 10:43

import (
	"fmt"
	"reflect"
	"strings"
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
	var result strings.Builder
	switch v.Kind() {
	case reflect.Invalid:
		result.WriteString("nil")
	case reflect.Ptr:
		if !v.IsNil() {
			result.WriteString(printValue(v.Elem()))
		} else {
			result.WriteString("nil")
		}
	case reflect.Struct:
		result.WriteString(v.Type().Name() + "{")
		for i := 0; i < v.NumField(); i++ {
			result.WriteString(v.Type().Field(i).Name + ":")
			result.WriteString(printValue(v.Field(i)))
			if i != v.NumField()-1 {
				result.WriteString(",")
			}
		}
		result.WriteString("}")
	case reflect.Slice, reflect.Array:
		result.WriteString("[")
		for i := 0; i < v.Len(); i++ {
			result.WriteString(printValue(v.Index(i)))
			if i != v.Len()-1 {
				result.WriteString(",")
			}
		}
		result.WriteString("]")
	case reflect.Map:
		result.WriteString("map[")
		for i, key := range v.MapKeys() {
			result.WriteString(printValue(key) + ":" + printValue(v.MapIndex(key)))
			if i != len(v.MapKeys())-1 {
				result.WriteString(",")
			}
		}
		result.WriteString("]")
	default:
		result.WriteString(fmt.Sprintf("%+v", v))
	}
	return result.String()
}
