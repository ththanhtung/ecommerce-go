package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Name     string
	Port     string
	Host     string
	User     string
	Type     string
	Password string
}

type ServerConfig struct {
	Port string
}

func LoadDBConfig(filename string) (*DBConfig, *ServerConfig) {
	if err := godotenv.Load(filename); err != nil {
		log.Fatal("error loading env file")
	}

	dbConfig := &DBConfig{
		Name:     os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Type:     os.Getenv("DB_TYPE"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	serverConfig := &ServerConfig{
		Port: os.Getenv("SERVER_PORT"),
	}
	return dbConfig, serverConfig
}
