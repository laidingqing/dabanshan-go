// Stores configuration for the core services
package config

import (
	"fmt"
	"log"

	"github.com/alyu/configparser"
)

// Load config values from file
func LoadConfig(filename string) {
	LoadDefaults()
	log.Printf("\nLoading Configuration from %v\n", filename)
	config, err := configparser.Read(filename)
	fmt.Print(config)
	if err != nil {
		log.Fatal(err)
	}

}
