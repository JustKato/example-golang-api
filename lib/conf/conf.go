package conf

import (
	"fmt"
	"os"
	"strconv"
)

type ServerConfig struct {
	Host string
	Port int
}

type DatabaseConfig struct {
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

var configuration *Config = nil

func GetConfig() *Config {

	if configuration != nil {
		return configuration
	}

	port, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		fmt.Println("Error:", err)
		port = 8081
	}

	// Generate the configuration from the Environment Variables
	sc := ServerConfig{
		Host: os.Getenv("host"),
		Port: port,
	}

	configuration = &Config{
		Server:   sc,
		Database: DatabaseConfig{},
	}

	return configuration

}
