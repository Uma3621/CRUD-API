package Configurations

//creating or defining a structure of config file
type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
}

type ServerConfig struct {
	Port int `json:"port"`
}
type DatabaseConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Name string `json:"name"`
	// User     string `json:"user"`
	// Password string `json:"password"`
}
