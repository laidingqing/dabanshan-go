package config

import (
	"log"

	"github.com/laidingqing/dabanshan/common/util"
)

var Service struct {
	DomainName       string
	ServiceName      string
	Port             int
	RegistryLocation string
}

var Logger struct {
	LogFile    string
	MaxSize    int
	MaxBackups int
	MaxAge     int
}

// Initialize Default values
func LoadDefaults() {
	_, err := util.GetExecDirectory()
	if err != nil {
		log.Fatal(err)
	}

	Logger.LogFile = "service.log"
	Logger.MaxSize = 10
	Logger.MaxBackups = 3
	Logger.MaxAge = 30
}
