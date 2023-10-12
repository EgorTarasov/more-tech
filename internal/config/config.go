package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI   string
	ServerPort string
	MongoDb    string
	SecretKey  string
}

var Cfg *Config

func NewConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	Cfg = &Config{
		MongoURI:   os.Getenv("MONGO_URI"),
		ServerPort: os.Getenv("SERVER_PORT"),
		MongoDb:    os.Getenv("MONGO_DB"),
		SecretKey:  os.Getenv("SECRET_KEY"),
	}

	return nil
}
