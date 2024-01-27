package utils

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"strings"
	"time"
)

type StructField struct {
	Name    string
	Type    ast.Expr
	JsonTag string
}

type MyStructType struct {
	StructName string
	Fields     []StructField
}

func getFieldList(fields *ast.FieldList) []StructField {
	var result []StructField

	if fields != nil && fields.List != nil {
		for _, field := range fields.List {

			var name = field.Names[0].Name

			var tags = field.Tag

			var tag string

			if tags != nil {
				var tagsInfo = reflect.StructTag(tags.Value[1 : len(tags.Value)-1])
				if t := tagsInfo.Get("json"); t == "" || t == "-" {
					tag = name
				} else {
					tag = t
				}
			} else {
				tag = name
			}

			result = append(result, StructField{
				Name:    name,
				Type:    field.Type,
				JsonTag: tag,
			})
		}
	}

	return result
}

func startStructToJSON(str string) map[string]MyStructType {
	var maps = make(map[string]MyStructType)

	var fset = token.NewFileSet()

	var filePath = fmt.Sprintf("%d.go", time.Now().UnixMicro())

	os.WriteFile(filePath, []byte("package main \n\n "+str), os.ModePerm)

	var f, _ = parser.ParseFile(fset, filePath, nil, parser.ParseComments)

	os.Remove(filePath)

	for _, decl := range f.Decls {
		genDecl, _ := decl.(*ast.GenDecl)
		if genDecl.Tok == token.TYPE {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					var structName = typeSpec.Name.Name
					if structType, ok := typeSpec.Type.(*ast.StructType); ok {
						var fileds = getFieldList(structType.Fields)
						maps[structName] = MyStructType{
							StructName: structName,
							Fields:     fileds,
						}
					}
				}
			}
		}
	}
	return maps
}

func getTypeValue(astType ast.Expr) interface{} {
	switch t := astType.(type) {
	case *ast.Ident:
		var _type = t.Name
		switch _type {
		case "int", "uint", "int64", "int32", "int16", "int8", "uint64", "uint32", "uint16", "uint8", "float64", "float32":
			return 0
		case "string":
			return "\"\""
		case "bool":
			return false
		default:
			return _type
		}
	case *ast.ArrayType:
		return fmt.Sprintf("[%v]", getTypeValue(t.Elt))
	case *ast.MapType:
		return "{}"
	default:
		return "null"
	}
}

func switchAppend(maps map[string]MyStructType, fields []StructField) (result []string) {
	for _, field := range fields {
		var _type = fmt.Sprintf("%v", field.Type)
		var v = getTypeValue(field.Type)
		if r, found := maps[_type]; found {
			result = append(result, fmt.Sprintf(`"%s":%v`, field.JsonTag, "{"+strings.Join(switchAppend(maps, r.Fields), ",")+"}"))
		} else {
			result = append(result, fmt.Sprintf(`"%s":%v`, field.JsonTag, v))
		}
	}
	return result
}

func foreachStruct(maps map[string]MyStructType, list []StructField) strings.Builder {
	var build strings.Builder

	build.WriteString("{")

	build.WriteString(strings.Join(switchAppend(maps, list), ","))

	build.WriteString("}")

	return build
}

func StructToJson(str string) string {

	var maps = startStructToJSON(str)

	var result = make([]string, 0)

	for _, structType := range maps {
		var build = foreachStruct(maps, structType.Fields)
		var str2 = strings.ReplaceAll(build.String(), " ", "")
		result = append(result, str2)
	}

	return fmt.Sprintf("[%v]", strings.Join(result, ","))
}
