package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func getTypeString(t reflect.Type) string {
	switch t.Kind() {
	case reflect.String:
		return "string"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
		return "int"
	case reflect.Bool:
		return "bool"
	default:
		return "interface{}"
	}
}

func createStruct(builds *[]strings.Builder, key string, v map[string]interface{}) {
	var build strings.Builder

	build.WriteString(fmt.Sprintf("type %s struct{\n", toCamelCase(key)))

	for key, value := range v {
		var newKey = toCamelCase(key)
		if value == nil {
			build.WriteString(fmt.Sprintf("     %s\t%s `json:\"%s\"`\n", newKey, "*interface{}", key))
			continue
		}
		var _type = reflect.TypeOf(value)
		var t = _type.Kind()
		switch t {
		case reflect.Map:
			build.WriteString(fmt.Sprintf("     %s\t%s `json:\"%s\"`\n", newKey, newKey, key))
			createStruct(builds, newKey, value.(map[string]interface{}))
			break
		case reflect.Slice:
			var field = toCamelCase(key)
			var fieldType = ArrayIsMap(builds, key, value.([]interface{}))
			build.WriteString(fmt.Sprintf("     %s\t%s `json:\"%s\"`\n", field, fieldType, key))
			break
		default:
			build.WriteString(fmt.Sprintf("     %s\t%s `json:\"%s\"`\n", newKey, getTypeString(_type), key))
		}
	}

	build.WriteString("}")

	*builds = append(*builds, build)
}

func ArrayIsMap(builds *[]strings.Builder, key string, _type []interface{}) string {
	if len(_type) == 0 {
		return "[]interface{}"
	}

	var v = _type[0]

	var kind = reflect.TypeOf(v).Kind()

	if kind == reflect.Map {
		createStruct(builds, key, v.(map[string]interface{}))
		return toCamelCase(key)
	} else {
		return "[]" + kind.String()
	}

	return ""
}

// Json2Struct 解析json字符串
func Json2Struct(builds *[]strings.Builder, build *strings.Builder, maps map[string]interface{}) {
	for key, value := range maps {
		var newKey = toCamelCase(key)
		if value == nil {
			build.WriteString(fmt.Sprintf("     %s\t%s `json:\"%s\"`\n", newKey, "*interface{}", key))
			continue
		}
		var _type = reflect.TypeOf(value)
		switch _type.Kind() {
		case reflect.Map:
			createStruct(builds, newKey, value.(map[string]interface{}))
			build.WriteString(fmt.Sprintf("     %s\t%s `json:\"%s\"`\n", newKey, newKey, key))
			break
		case reflect.Slice:
			var field = toCamelCase(key)
			var fieldType = ArrayIsMap(builds, key, value.([]interface{}))
			build.WriteString(fmt.Sprintf("     %s\t%s `json:\"%s\"`\n", field, fieldType, key))
			break
		default:
			build.WriteString(fmt.Sprintf("     %s\t%s `json:\"%s\"`\n", newKey, getTypeString(_type), key))
			break
		}
	}
}

func toCamelCase(s string) string {
	s = strings.ReplaceAll(s, "-", "_")
	s = strings.ReplaceAll(s, "_", " ")

	words := strings.Split(s, " ")
	for i, word := range words {
		words[i] = strings.Title(word)
	}

	return strings.Join(words, "")
}
