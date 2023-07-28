package database

import (
	"context"
	"ecom/configs"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	config *configs.DBConfig
	client *mongo.Client
}

func NewMongoDB(config *configs.DBConfig) *MongoDB {
	var mongo *MongoDB = &MongoDB{}

	mongo.config = config
	mongo.client = createDBInstance(config)

	return mongo
}

func createDBInstance(config *configs.DBConfig) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Host + ":" + config.Port))
	if err != nil {
		log.Fatal("error connecting database:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err = client.Connect(ctx); err != nil {
		log.Fatal("error connection database:", err)
	}
	defer cancel()

	log.Println("connected to mongodb")
	return client
}

func (m *MongoDB) OpenCollection(collectionName string) *mongo.Collection {
	colletion := m.client.Database(m.config.Name).Collection(collectionName)
	return colletion
}
