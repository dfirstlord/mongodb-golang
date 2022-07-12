package config

import "os"

type ApiConfig struct {
	ApiPort string
	ApiHost string
}

type MongoConfig struct {
	Host     string
	Port     string
	DbName   string
	User     string
	Password string
}

type Config struct {
	ApiConfig
	MongoConfig
}

func (c Config) readConfig() Config {
	c.MongoConfig = MongoConfig{
		Host:     os.Getenv("MONGO_HOST"),
		Port:     os.Getenv("MONGO_PORT"),
		DbName:   os.Getenv("MONGO_DB"),
		User:     os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PASSWORD"),
	}
	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_HOST"),
		ApiHost: os.Getenv("API_PORT"),
	}
	return c
}

func NewConfig() Config{
	cfg := Config{}
	return cfg.readConfig()
}
