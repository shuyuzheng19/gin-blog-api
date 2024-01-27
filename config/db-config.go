package config

import (
	"common-web-framework/common"
	"common-web-framework/helper"
	"common-web-framework/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"strings"
)

var DB *gorm.DB

// DbConfig 关系型数据库配置
type DbConfig struct {
	Log bool `yaml:"log" json:"log"`
	//空闲数量
	MaxIdle int `yaml:"maxIdle" json:"maxIdle"`
	//最大连接量
	MaxSize int `yaml:"maxSize" json:"maxSize"`
	//数据库时区
	Timezone string `yaml:"timezone" json:"timezone"`
	//数据库厂商 如mysql postgresql等 仅限gorm支持的数据库
	Database string `yaml:"database" json:"database"`
	//数据库远程地址
	Host string `yaml:"host" json:"host"`
	//数据库端口
	Port int `yaml:"port" json:"port"`
	//数据库登录账号
	Username string `yaml:"username" json:"username"`
	//数据库密码
	Password string `yaml:"password" json:"password"`
	//数据库名称
	Dbname string `yaml:"dbname" json:"dbname"`

	AutoCreate bool `yaml:"autoCreate" json:"autoCreate"`
}

// getDSN 根据 DbConfig 返回数据库连接字符串
func getDataBaseDSN(config DbConfig) string {
	var database = strings.ToLower(config.Database)
	switch database {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
			config.Username, config.Password, config.Host, config.Port, config.Dbname, config.Timezone)
	case "postgresql":
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable Timezone=%s",
			config.Host, config.Port, config.Username, config.Dbname, config.Password, config.Timezone)
	case "sqlite":
		return config.Dbname // Assuming that Dbname is the path to SQLite file
	case "sqlserver":
		return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
			config.Username, config.Password, config.Host, config.Port, config.Dbname)
	case "oracle":
		return fmt.Sprintf("%s/%s@%s:%d/%s",
			config.Username, config.Password, config.Host, config.Port, config.Dbname)
	case "cockroachdb":
		return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
			config.Username, config.Password, config.Host, config.Port, config.Dbname)
	case "clickhouse":
		return fmt.Sprintf("tcp://%s:%d?username=%s&password=%s&database=%s",
			config.Host, config.Port, config.Username, config.Password, config.Dbname)
	case "bigquery":
		return fmt.Sprintf("bigquery://%s:%s@projectID:%s/datasetID",
			config.Username, config.Password, config.Dbname)
	default:
		var err interface{} = fmt.Sprintf("未知的数据库 %s", config.Database)
		panic(err)
	}
}

// LoadDBConfig 加载数据库配置
func LoadDBConfig(dbConfig DbConfig) {

	//数据库链接
	var dsn = getDataBaseDSN(dbConfig)

	var gormConfig = &gorm.Config{}

	if dbConfig.Log {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	} else {
		gormConfig.Logger = logger.Default.LogMode(logger.Silent)
	}

	//假设是mysql 当然也可以使用其他的数据，只需导入相应的数据库驱动即可
	var db, err = gorm.Open(postgres.Open(dsn), gormConfig)
	//判断是否链接成功
	helper.ErrorPanicAndMessage(err, "数据库加载失败")

	var connection, _ = db.DB()

	connection.SetMaxIdleConns(dbConfig.MaxIdle)

	connection.SetMaxOpenConns(dbConfig.MaxSize)

	DB = db

	if dbConfig.AutoCreate {
		DB.AutoMigrate(&models.Role{})
		var roles = []models.Role{
			{
				ID:          common.UserRole,
				Name:        "USER",
				Description: "普通用户",
			},
			{
				ID:          common.AdminRole,
				Name:        "ADMIN",
				Description: "管理员",
			},
			{
				ID:          common.SuperAdminRole,
				Name:        "SUPER_ADMIN",
				Description: "超级管理员",
			},
		}
		DB.Model(&models.Role{}).Save(&roles)
		DB.AutoMigrate(&models.User{}, &models.Tag{}, &models.Category{},
			&models.FileInfo{}, &models.Blog{}, &models.Topic{})
	}

	//如果连接成功 打印 DataBase Connection SUCCESS
	log.Println("DataBase Connection SUCCESS")
}
