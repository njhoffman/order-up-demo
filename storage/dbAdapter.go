package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"sync"
)

var connect_db string
var ok bool

const (
	DBCONNECT   = "mongodb://localhost:27017"
	DBNAME      = "order_up"
	C_ORDERS    = "orders"
	C_CUSTOMERS = "customers"
)

var mongoInstance *mongo.Client
var mongoError error
var mongoOnce sync.Once

func GetMongoClient() (*mongo.Client, error) {
  // init db once and assign instance for app-wide singleton use
	mongoOnce.Do(func() {

    // env variable for dynamic address assignment
    connect_db, ok = os.LookupEnv("MONGO_CONNECT")
    if !ok {
      connect_db = DBCONNECT
    }
		mongoOptions := options.Client().ApplyURI(connect_db)
		mongoClient, err := mongo.Connect(context.TODO(), mongoOptions)
		if err != nil {
			mongoError = err
		}

		// ensure mongodb server is up
		err = mongoClient.Ping(context.TODO(), nil)
		if err != nil {
			mongoError = err
		}
		mongoInstance = mongoClient
	})
	return mongoInstance, mongoError
}
