package xtools

import (
	"github.com/astaxie/beego/logs"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"github.com/uniplaces/carbon"
	"os"
)

func LogrusStdLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: carbon.DefaultFormat,
	})
	return logger
}

func LogrusFileLogger(path string) (*logrus.Logger, error) {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(&lumberjack.Logger{
		Filename:   path,
		MaxSize:    500, // megabytes
		MaxBackups: 30,
		MaxAge:     365,   //days
		Compress:   false, // disabled by default
	})
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: carbon.DefaultFormat,
	})
	return logger, nil
}

func LogsStdLogger() {
	logs.SetLogFuncCall(true)
	logs.SetLogFuncCallDepth(3)
}
