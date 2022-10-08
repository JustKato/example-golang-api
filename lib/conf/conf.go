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
	Host string
	Port int
	User string
	Pass string
	Db   string

	MaxConn     int
	MaxIdleConn int
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

var configuration *Config = nil

// Get the configuration of the app, this will return all configs such as database, server, etc...
func GetConfig(forceReload bool) *Config {

	if configuration != nil && !forceReload {
		return configuration
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		fmt.Println("Error:", err)
		port = 8081
	}

	// Generate the configuration from the Environment Variables
	sc := ServerConfig{
		Host: os.Getenv("HOST"),
		Port: port,
	}

	configuration = &Config{
		Server:   sc,
		Database: DatabaseConfig{},
	}

	// Handle database cofiguration

	DATABASE_HOST, exists := os.LookupEnv("DATABASE_HOST")
	if !exists {
		DATABASE_HOST = "localhost"
	}

	DATABASE_PORT_STRING, exists := os.LookupEnv("DATABASE_PORT")
	if !exists {
		DATABASE_PORT_STRING = "3306"
	}

	DATABASE_PORT, err := strconv.Atoi(DATABASE_PORT_STRING)
	if err != nil {
		DATABASE_PORT = 3306
	}

	DATABASE_NAME, exists := os.LookupEnv("DATABASE_NAME")
	if !exists {
		DATABASE_NAME = "main"
	}

	DATABASE_USER, exists := os.LookupEnv("DATABASE_USER")
	if !exists {
		DATABASE_USER = "root"
	}

	DATABASE_PASS, exists := os.LookupEnv("DATABASE_PASS")
	if !exists {
		DATABASE_PASS = ""
	}

	DATABASE_MAXC_STRING, exists := os.LookupEnv("DATABASE_MAXC")
	if !exists {
		DATABASE_MAXC_STRING = "10"
	}

	DATABASE_MAXC, err := strconv.Atoi(DATABASE_MAXC_STRING)
	if err != nil {
		DATABASE_MAXC = 10
	}

	DATABASE_MAXT_STRING, exists := os.LookupEnv("DATABASE_MAXT")
	if !exists {
		DATABASE_MAXT_STRING = "10"
	}

	DATABASE_MAXT, err := strconv.Atoi(DATABASE_MAXT_STRING)
	if err != nil {
		DATABASE_MAXT = 10
	}

	configuration.Database.Host = DATABASE_HOST
	configuration.Database.Port = DATABASE_PORT
	configuration.Database.User = DATABASE_USER
	configuration.Database.Pass = DATABASE_PASS
	configuration.Database.Db = DATABASE_NAME
	configuration.Database.MaxConn = DATABASE_MAXC
	configuration.Database.MaxIdleConn = DATABASE_MAXT

	return configuration

}
