package xtools

import "testing"

func TestNewFileLogger(t *testing.T) {
	logger, _ := NewFileLogger("./test.log")

	logger.Infoln("hahaha hhehehe")
	logger.Errorln("hahaha hhehehe")
	logger.Println("hahaha hhehehe")
}
