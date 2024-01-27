package service

import (
	"common-web-framework/common"
	"common-web-framework/helper"
	"common-web-framework/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	mjson "github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
	"reflect"
	"sigs.k8s.io/yaml"
	"strings"
)

type ToolsService struct {
	m *minify.M
}

func (*ToolsService) JsonToYaml(buff []byte) string {

	var valid = json.Valid(buff)

	if !valid {
		helper.ErrorToResponse(common.NotJson)
	}

	var yamlStr, _ = yaml.JSONToYAML(buff)

	return string(yamlStr)
}

func (*ToolsService) YamlToJson(buff []byte) string {

	var maps = map[string]interface{}{}

	var err = yaml.Unmarshal(buff, &maps)

	if err != nil {
		helper.ErrorToResponse(common.ConvertFail)
	}

	return utils.ObjectToJson(maps)
}

func (*ToolsService) StructToJSON(buff []byte) string {
	return utils.StructToJson(string(buff))
}

func (*ToolsService) JsonToStruct(buff []byte) string {
	var data interface{}

	var builds = []strings.Builder{{}}

	var err = json.Unmarshal(buff, &data)

	if err != nil {
		helper.ErrorToResponse(common.NotJson)
	}

	var build = strings.Builder{}

	if reflect.TypeOf(data).Kind() == reflect.Slice {
		var datas = data.([]interface{})
		for index, data := range datas {
			var m = data.(map[string]interface{})
			build.WriteString(fmt.Sprintf("type T%d struct { \n", index))
			utils.Json2Struct(&builds, &build, m)
			build.WriteString("}\n")
		}
	} else {
		var maps = data.(map[string]interface{})
		build.WriteString("type T struct { \n")

		utils.Json2Struct(&builds, &build, maps)

		build.WriteString("}\n")
	}

	for _, builder := range builds {
		build.WriteString(builder.String() + "\n\n")
	}

	return build.String()
}

func (t ToolsService) Compress(_type string, str string) string {
	var result, err = t.m.String(_type, str)
	if err != nil {
		helper.ErrorToResponse(common.FailCode)
	}
	return result
}

func (t ToolsService) Base64Encoding(buff []byte) string {
	return base64.URLEncoding.EncodeToString(buff)
}

func (t ToolsService) Base64Decoding(buff []byte) string {
	var result, err = base64.URLEncoding.DecodeString(string(buff))
	if err != nil {
		helper.ErrorToResponse(common.FailCode)
	}
	return string(result)
}

func (t ToolsService) CompressJson(buff []byte) string {
	return t.Compress("json", string(buff))
}

func (t ToolsService) CompressHTML(buff []byte) string {
	return t.Compress("html", string(buff))
}

func (t ToolsService) CompressCSS(buff []byte) string {
	return t.Compress("css", string(buff))
}

func (t ToolsService) CompressJavaScript(buff []byte) string {
	return t.Compress("js", string(buff))
}

func (*ToolsService) FormatJson(buff []byte) string {
	var maps interface{}
	err := json.Unmarshal(buff, &maps)
	if err != nil {
		helper.ErrorToResponse(common.NotJson)
	}
	result, err := json.MarshalIndent(&maps, "", "    ")
	if err != nil {
		helper.ErrorToResponse(common.ConvertFail)
	}
	return string(result)
}

func NewToolsService() ToolsService {
	m := minify.New()
	m.AddFunc("css", css.Minify)
	m.AddFunc("html", html.Minify)
	m.AddFunc("svg", svg.Minify)
	m.AddFunc("json", mjson.Minify)
	m.AddFunc("js", js.Minify)
	return ToolsService{m: m}
}
