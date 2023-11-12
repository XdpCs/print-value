package print_value

// @Title        print.go
// @Description
// @Create       XdpCs 2023-11-09 10:51
// @Update       XdpCs 2023-11-12 21:23

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

// Print returns a string representation of v.
// If you want to print a struct, struct field names must be exported.
func Print(v interface{}) string {
	return printValue(reflect.ValueOf(v))
}

func printValue(v reflect.Value) string {
	result := getStringBuilder()
	defer putStringBuilder(result)
	printFieldValue(result, v)
	return result.String()
}

func printFieldValue(result *strings.Builder, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		result.WriteString("nil")
	case reflect.Ptr:
		if !v.IsNil() {
			printFieldValue(result, v.Elem())
		} else {
			result.WriteString("nil")
		}
	case reflect.Struct:
		result.WriteString(v.Type().Name() + "{")
		for i := 0; i < v.NumField(); i++ {
			result.WriteString(v.Type().Field(i).Name + ":")
			printFieldValue(result, v.Field(i))
			if i != v.NumField()-1 {
				result.WriteString(",")
			}
		}
		result.WriteString("}")
	case reflect.Slice, reflect.Array:
		result.WriteString("[")
		for i := 0; i < v.Len(); i++ {
			printFieldValue(result, v.Index(i))
			if i != v.Len()-1 {
				result.WriteString(",")
			}
		}
		result.WriteString("]")
	case reflect.Map:
		result.WriteString("map[")
		keys := v.MapKeys()
		for i, key := range keys {
			printFieldValue(result, key)
			result.WriteString(":")
			printFieldValue(result, v.MapIndex(key))
			if i != len(keys)-1 {
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
}

var stringBuilderPool = sync.Pool{
	New: func() interface{} {
		return &strings.Builder{}
	},
}

func getStringBuilder() *strings.Builder {
	return stringBuilderPool.Get().(*strings.Builder)
}

func putStringBuilder(builder *strings.Builder) {
	builder.Reset()
	stringBuilderPool.Put(builder)
}
