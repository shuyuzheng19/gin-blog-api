package route

import (
	"common-web-framework/controller"
	"common-web-framework/service"
)

// LoadToolsController 工具相关接口
func (r *RouterSetup) LoadToolsController(name string) *RouterSetup {
	var group = r.apiGroup.Group(name)

	var toolsService = service.NewToolsService()

	var toolsController = controller.NewToolsController(toolsService)

	{
		group.POST("quote", toolsController.StringQuote)
		group.POST("un_quote", toolsController.UnStringQuote)
		group.POST("base64", toolsController.Base64EncodingOrDecoding)
		group.POST("compress", toolsController.Compress)
		group.POST("compress_file", toolsController.CompressFromFile)
		group.POST("struct2json", toolsController.StructToJSON)
		group.POST("json2struct", toolsController.Json2Struct)
		group.POST("json2yaml", toolsController.JsonToYaml)
		group.POST("yaml2json", toolsController.YamlToJson)
		group.POST("format_json", toolsController.FormatJson)
		group.POST("compress_json", toolsController.CompressJson)

	}

	return r
}
