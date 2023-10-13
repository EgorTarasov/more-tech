package mongodb

import (
	"context"
	"errors"
	"fmt"
	"more-tech/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDb() (*mongo.Client, error) {
	ctx := context.Background()

	attemptsCount := 0
	for {
		if attemptsCount > 5 {
			return nil, errors.New("too many attempts while connecting to mongo")
		}
		attemptsCount++
		time.Sleep(time.Second * 2)
		// con_string := "mongodb://mongouser:mongopass@localhost:27017/"
		conString := fmt.Sprintf("mongodb://%s:%s@%s:%s/", config.Cfg.MongoUser, config.Cfg.MongoPassword, config.Cfg.MongoHost, config.Cfg.MongoPort)
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(conString))
		if err != nil {
			fmt.Printf("can't connect to mongo: %+v", err)
			continue
		}

		if err := client.Ping(ctx, nil); err != nil {
			fmt.Printf("can't access mongo: %+v", err)
			continue
		}

		return client, nil
	}
}
