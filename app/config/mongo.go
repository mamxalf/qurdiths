package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hadithgo/app/exceptions"
	"strconv"
	"time"
)

func ConnectMongoDatabase(configuration Config) *mongo.Database {
	ctx, cancel := NewMongoContext()
	defer cancel()

	mongoPoolMin, err := strconv.Atoi(configuration.Get("MONGO_POOL_MIN"))
	exceptions.PanicIfNeeded(err)

	mongoPoolMax, err := strconv.Atoi(configuration.Get("MONGO_POOL_MAX"))
	exceptions.PanicIfNeeded(err)

	mongoMaxIdleTime, err := strconv.Atoi(configuration.Get("MONGO_MAX_IDLE_TIME_SECOND"))
	exceptions.PanicIfNeeded(err)

	option := options.Client().
		ApplyURI(configuration.Get("MONGO_URI")).
		SetMinPoolSize(uint64(mongoPoolMin)).
		SetMaxPoolSize(uint64(mongoPoolMax)).
		SetMaxConnIdleTime(time.Duration(mongoMaxIdleTime) * time.Second)

	client, err := mongo.NewClient(option)
	exceptions.PanicIfNeeded(err)

	err = client.Connect(ctx)
	exceptions.PanicIfNeeded(err)

	database := client.Database(configuration.Get("MONGO_DATABASE"))
	return database
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
