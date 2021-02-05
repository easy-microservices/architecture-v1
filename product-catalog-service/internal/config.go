package internal

import (
	"fmt"

	"github.com/joho/godotenv"
)

// Config defines all configuration variables used
// in the whole service.
type Config struct {
	vars map[string]string
}

// LoadConfig loads all properties from the configuration file
func LoadConfig(fileName string) *Config {

	configFileContent, err := godotenv.Read(fileName)
	if err != nil {
		panic(err)
	}

	return &Config{
		vars: configFileContent,
	}
}

// GetConfigProperty retrieves the value of a configuration property
// if exists in the configuration map.
func (c *Config) GetConfigProperty(property string) string {

	k, exist := c.vars[property]
	if !exist {
		fmt.Printf("No configuration value for key %s", property)
	}

	return k
}
