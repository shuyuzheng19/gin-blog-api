package config

import (
	"common-web-framework/helper"
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
	"strings"
)

// LoggerConfig 日志配置
type LoggerConfig struct {
	Dev         bool   `yaml:"dev" json:"dev"`
	Encoding    string `yaml:"encoding" json:"encoding"`
	OutputPaths string `yaml:"outputPaths" json:"outputPaths"`
	ErrorPaths  string `yaml:"errorPaths" json:"errorPaths"`
	Level       string `yaml:"level" json:"level"`
	LoggerDir   string `yaml:"loggerDir" json:"loggerDir"`
	DefaultName string `yaml:"defaultName" json:"defaultName"`
}

var maps = map[string]zapcore.Level{
	"info":  zapcore.InfoLevel,
	"debug": zapcore.DebugLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"panic": zapcore.PanicLevel,
}

var LOGGER *zap.Logger

// LoadLogger 加载zap日志
func LoadLogger(loggerConfig LoggerConfig) {

	if loggerConfig.LoggerDir == "" {
		helper.ErrorPanicAndMessage(errors.New("log dir path empty"), "日志文件目录不能为空")
	}

	if loggerConfig.DefaultName == "" {
		helper.ErrorPanicAndMessage(errors.New("log defaultName empty"), "日志默认文件名不能为空")
	}

	config := zap.NewProductionConfig()

	config.Encoding = loggerConfig.Encoding

	config.Development = loggerConfig.Dev

	var level, found = maps[loggerConfig.Level]

	if !found {
		level = zap.InfoLevel
	}

	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	config.Level = zap.NewAtomicLevelAt(level)

	var paths = make([]string, 0)

	if loggerConfig.OutputPaths != "" {
		paths = strings.Split(loggerConfig.OutputPaths, ",")
	}

	paths = append(paths, filepath.Join(loggerConfig.LoggerDir, loggerConfig.DefaultName))

	config.OutputPaths = paths

	if loggerConfig.ErrorPaths != "" {

		var errorPaths = strings.Split(loggerConfig.ErrorPaths, ",")

		if len(errorPaths) > 0 {
			config.ErrorOutputPaths = errorPaths
		}
	}

	var logger, err = config.Build()

	helper.ErrorPanicAndMessage(err, "加载日志配置失败")

	LOGGER = logger

	defer LOGGER.Sync()
}
