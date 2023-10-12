package mongodb

import (
	"context"
	"more-tech/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDb() (*mongo.Client, error) {
	ctx := context.Background()

	for {
		time.Sleep(time.Second*2)
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Cfg.MongoURI))
		if err != nil {
			continue
		}

		if err := client.Ping(ctx, nil); err != nil {
			continue
		}

		return client, nil
	}
}