package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI("<MONGODB_URI>"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

}

var (
	// ErrOrderNotFound is returned when the specified order cannot be found
	ErrOrderNotFound = errors.New("order not found")

	// ErrOrderExists is returned when a new order is being inserted but an order
	// with the same ID already exists
	ErrOrderExists = errors.New("order already exists")
)

////////////////////////////////////////////////////////////////////////////////

// GetOrder should return the order with the given ID. If that ID isn't found then
// the special ErrOrderNotFound error should be returned.
func (i *Instance) GetOrder(ctx context.Context, id string) (Order, error) {
	// TODO: get order from DB based on the id
	return Order{}, errors.New("unimplemented")
}

////////////////////////////////////////////////////////////////////////////////

// GetOrders should return all orders with the given status. If status is the
// special -1 value then it should return all orders regardless of their status.
func (i *Instance) GetOrders(ctx context.Context, status OrderStatus) ([]Order, error) {
	// TODO: get orders from DB based based on the status sent, unless status is -1
	return nil, errors.New("unimplemented")
}

////////////////////////////////////////////////////////////////////////////////

// SetOrderStatus should update the order with the given ID and set the status
// field. If that ID isn't found then the special ErrOrderNotFound error should
// be returned.
func (i *Instance) SetOrderStatus(ctx context.Context, id string, status OrderStatus) error {
	// TODO: update the order's status field to status for the id
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// InsertOrder should fill in the order's ID with a unique identifier if it's not
// already set and then insert it into the database. It should return the order's
// ID. If the order already exists then ErrOrderExists should be returned.
func (i *Instance) InsertOrder(ctx context.Context, order Order) (string, error) {
	// TODO: if the order's ID field is empty, generate a random ID, then insert
	// into the database
	return "", errors.New("unimplemented")
}
