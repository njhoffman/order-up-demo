// Package storage contains the code to persist and retrieve orders from a database
package storage

import (
	"context"
	"github.com/levenlabs/go-llog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"sync"
	"time"
)

// database name is randomized to prevent wiping before each test
type Instance struct {
	dbconnect   string
	database    string
	c_orders    string
	c_customers string
	client      *mongo.Client
	mongoError  error
	mongoOnce   sync.Once
}

func New(overrideDatabase string) *Instance {
	// create a pointer to an Instance that we will return after initialization
	var ok bool
	inst := &Instance{}
	// if they sent overrideDatabase then use that, like for tests
	if overrideDatabase != "" {
		inst.database = overrideDatabase
	} else {
		inst.database, ok = os.LookupEnv("MONGO_DATABASE")
		if !ok {
			inst.database = "order_up"
		}
	}
	// env variable for dynamic address assignment
	inst.dbconnect, ok = os.LookupEnv("MONGO_DBCONNECT")
	if !ok {
		inst.dbconnect = "mongodb://localhost:27017"
	}

	// collection constants
	inst.c_orders = "orders"
	inst.c_customers = "customers"

	mongoOptions := options.Client().ApplyURI(inst.dbconnect)
	mongoClient, err := mongo.Connect(context.TODO(), mongoOptions)
	if err != nil {
		inst.mongoError = err
	}
	// ping to ensure mongodb server is responsive
	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		inst.mongoError = err
	}

	inst.client = mongoClient

	// give the ensureSchema function only 15 seconds to complete
	// after 15 seconds the context will return DeadlineExceeded errors which should
	// cause any functions downstream to error out
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// if we don't call cancel then the ctx will leak so we make sure that cancel
	// is called no matter what when we're done
	defer cancel()
	// we want to make sure the database is ready to accept requests and if that
	// fails we need to fatal
	if err := inst.ensureSchema(ctx); err != nil {
		llog.Fatal("failed to ensure schema", llog.ErrKV(err))
	}
	return inst
}

func (i *Instance) ensureSchema(ctx context.Context) error {
	// TODO: this is where you'll do any schema setup (CREATE DATABASE or CREATE
	// TABLE), if necessary, and since this will be called every time the service
	// starts or every time you run tests, it should not fail if the schema is
	// already setup
	// for example you might need to add a unique index on the order's ID field ;)
	return nil
}
