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

// Initialize Default values
func LoadDefaults() {
	_, err := util.GetExecDirectory()
	if err != nil {
		log.Fatal(err)
	}
}
