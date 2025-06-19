// Package utils
// Copyright: Copyright (c) 2020<br>
// Company: 易宝支付(YeePay)<br>
// @author    : yunmei.wu
// @time      : 2023/3/16 3:22 PM
package utils

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// Logger 全局日志记录器，可以被外部自定义
var Logger *logrus.Logger

func init() {
	// 默认日志配置：使用 logrus
	Logger = logrus.New()
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(logrus.InfoLevel)
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   false,
	})
}

// SetLogger 设置自定义日志记录器
func SetLogger(logger *logrus.Logger) {
	Logger = logger
}

// DisableLogging 禁用日志输出
func DisableLogging() {
	Logger.SetOutput(io.Discard)
}

// EnableLogging 启用日志输出到标准输出
func EnableLogging() {
	Logger.SetOutput(os.Stdout)
}

// SetLogLevel 设置日志级别
func SetLogLevel(level logrus.Level) {
	Logger.SetLevel(level)
}

// SetLogFormatter 设置日志格式化器
func SetLogFormatter(formatter logrus.Formatter) {
	Logger.SetFormatter(formatter)
}
