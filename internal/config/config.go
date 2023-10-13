package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoUser     string
	MongoPassword string
	MongoHost     string
	MongoPort     string
	MongoDb       string
	ServerPort    string
	SecretKey     string
}

var Cfg *Config

func NewConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	Cfg = &Config{
		MongoUser:     os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
		MongoPassword: os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
		MongoHost:     os.Getenv("MONGO_HOST"),
		MongoPort:     os.Getenv("MONGO_PORT"),
		MongoDb:       os.Getenv("MONGO_DB"),
		ServerPort:    os.Getenv("SERVER_PORT"),
		SecretKey:     os.Getenv("SECRET_KEY"),
	}

	dockerMode := os.Getenv("DOCKER_MODE")
	if dockerMode == "1" {
		Cfg.MongoHost = "mongo"
	} else {
		Cfg.MongoHost = "localhost"
	}
	if Cfg.MongoPort == "" {
		Cfg.MongoPort = "27017"
	}

	return nil
}
