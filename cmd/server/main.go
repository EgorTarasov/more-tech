package main

import (
	"context"
	"fmt"
	"more-tech/internal/api/router"
	"more-tech/internal/config"
	mongodb "more-tech/internal/database/mongo"
	"more-tech/internal/logging"
	"os"
)

//	@title			More Tech API
//	@description	More Tech API server
//	@host			localhost:9999
func main() {
	if err := config.NewConfig(); err != nil {
		fmt.Printf("can't load config: %+v", err)
		os.Exit(1)
	}

	if err := logging.NewLogger(); err != nil {
		fmt.Printf("can't create logger instance: %+v", err)
		os.Exit(1)
	}

	mongoClient, err := mongodb.NewMongoDb()
	if err != nil {
		logging.Log.Fatalf("can't create mongo client: %+v", err)
	}
	defer mongoClient.Disconnect(context.Background())

	router := router.NewRouter(mongoClient)
	logging.Log.Infof("starting server on port %s", config.Cfg.ServerPort)
	if err := router.Run(config.Cfg.ServerPort); err != nil {
		logging.Log.Fatalf("can't start server %+v", err)
	}
}
