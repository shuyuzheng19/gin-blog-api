package main

import (
	"common-web-framework/config"
	"common-web-framework/middle"
	"common-web-framework/route"
	"common-web-framework/utils"
	"os"
)

func AutoCreateUploadPath() {
	os.MkdirAll(config.CONFIG.Upload.Path, os.ModePerm)
}

func LoadConfig() {

	config.LoadGlobalConfig()

	var conf = config.CONFIG

	config.LoadDBConfig(conf.Db)

	config.LoadLogger(conf.Logger)

	config.LoadRedis(conf.Redis)

	utils.LoadIpDB(conf.IpDbPath)

	config.LoadSearchConfig(conf.Search)

	AutoCreateUploadPath()
}

// @title Yuice Blog
// @version 1.0
// @description Gin+GORM+Redis+MySQL+MeiliSearch 个人博客
// @termsOfService https://github.com/shuyuzheng19
// @BasePath /api/v1/
// @securityDefinitions.apikey  JWT
// @in                          header
// @name                        Authorization
// @description   输入token, `Bearer: ` 前缀, 示例: "Bearer abcde12345".
func main() {

	LoadConfig()

	var serverConfig = config.CONFIG.Server

	route.NewRouter(serverConfig).
		AddMiddles(middle.ErrorMiddle).
		SetupController().
		RunServer()

}
