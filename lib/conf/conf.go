package conf

type ServerConfig struct {
	Host string
	Port int
}

type Config struct {
	Server ServerConfig
}
