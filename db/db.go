package db

import (
	"context"
	"example/test-golang/config"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DB struct {
	Client *mongo.Client
}

func (db *DB) GetDB() *mongo.Database {
	return db.Client.Database("simple-gin-api")
}

func (db *DB) Shutdown(ctx context.Context) error {
	return db.Client.Disconnect(ctx)
}

func ProvideDB(cfg *config.Config) *DB {
	c, _ := mongo.Connect(options.Client().ApplyURI(cfg.MongoHost))

	return &DB{Client: c}
}
