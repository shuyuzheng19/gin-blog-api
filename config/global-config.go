package config

import (
	"common-web-framework/helper"
	"gopkg.in/yaml.v3"
	"os"
)

// CONFIG 全局配置
var CONFIG *GlobalConfig

// ServerConfig 服务器配置
type ServerConfig struct {
	Release      bool   `yaml:"release" json:"release"`
	Cron         bool   `yaml:"cron" json:"cron"`
	Addr         string `yaml:"addr" json:"addr"`
	ApiPrefix    string `yaml:"apiPrefix" json:"apiPrefix"`
	ReadTimeOut  int    `yaml:"readTimeOut" json:"readTimeOut"`
	WriteTimeOut int    `yaml:"writeTimeOut" json:"writeTimeOut"`
}

type CorsConfig struct {
	AllOrigins       bool     `yaml:"allOrigins" json:"allOrigins"`
	Enable           bool     `yaml:"enable" json:"enable"`
	AllowOrigins     []string `yaml:"allowOrigins"`
	AllowMethods     []string `yaml:"allowMethods"`
	AllowHeaders     []string `yaml:"allowHeaders"`
	ExposeHeaders    []string `yaml:"exposeHeaders"`
	AllowCredentials bool     `yaml:"allowCredentials"`
}

// GlobalConfig 全局配置
type GlobalConfig struct {
	DataBaseKey string `yaml:"databaseKey" json:"-"`
	//我的邮箱
	MyEmail string `yaml:"myEmail" json:"myEmail"`
	//是否开启cors
	EnableCors string `yaml:"enableCors" json:"enableCors"`
	//跨域配置
	Cors CorsConfig `yaml:"cors" json:"cors"`
	//IP数据库路径
	IpDbPath string `yaml:"ipDbPath" json:"ipDbPath"`
	//server配置
	Server ServerConfig `yaml:"server" json:"server"`
	//db配置
	Db DbConfig `yaml:"db" json:"db"`
	//邮箱配置
	Email EmailConfig `yaml:"email" json:"email"`
	//日志配置
	Logger LoggerConfig `yaml:"logger" json:"logger"`
	//redis配置
	Redis RedisConfig `yaml:"redis" json:"redis"`
	//上传文件配置
	Upload UploadConfig `yaml:"upload" json:"upload"`
	//搜索配置
	Search MeiliSearchConfig `yaml:"meilisearch" json:"meilisearch"`
}

// LoadGlobalConfig 加载全局配置
func LoadGlobalConfig() {
	var file, err = os.ReadFile("application.yml")

	helper.ErrorPanicAndMessage(err, "读取配置文件失败")

	yaml.Unmarshal(file, &CONFIG)
}
