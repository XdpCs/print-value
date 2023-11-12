package print_value

// @Title        print.go
// @Description
// @Create       XdpCs 2023-11-09 10:43
// @Update       XdpCs 2023-11-09 10:43

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

func Print(v interface{}) string {
	return print(v)
}

func print(v interface{}) string {
	value := reflect.ValueOf(v)
	return printValue(value)
}

func printValue(v reflect.Value) string {
	result := GetStringBuilder()
	defer PutStringBuilder(result)
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
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		result.WriteString(strconv.FormatInt(v.Int(), 10))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		result.WriteString(strconv.FormatUint(v.Uint(), 10))
	case reflect.String:
		result.WriteString(v.String())
	case reflect.Bool:
		result.WriteString(strconv.FormatBool(v.Bool()))
	case reflect.Float32, reflect.Float64:
		result.WriteString(strconv.FormatFloat(v.Float(), 'f', 2, 64))
	default:
		result.WriteString(fmt.Sprintf("%+v", v))
	}
	return result.String()
}

var stringBuilderPool = sync.Pool{
	New: func() interface{} {
		return &strings.Builder{}
	},
}

func GetStringBuilder() *strings.Builder {
	return stringBuilderPool.Get().(*strings.Builder)
}

func PutStringBuilder(builder *strings.Builder) {
	builder.Reset()
	stringBuilderPool.Put(builder)
}
