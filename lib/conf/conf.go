package conf

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
