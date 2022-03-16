package storage

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/json"
	"go.mongodb.org/mongo-driver/json/primitive"
	"log"
)

var (
	// ErrOrderNotFound is returned when the specified order cannot be found
	ErrOrderNotFound = errors.New("order not found")

	// ErrOrderExists is returned when a new order insert has existing ID
	ErrOrderExists = errors.New("order already exists")
)

////////////////////////////////////////////////////////////////////////////////

// GetOrder should return the order with the given ID. If that ID isn't found then
// the special ErrOrderNotFound error should be returned.
func (i *Instance) GetOrder(ctx context.Context, id string) (Order, error) {

	// create placeholder and json filter for order query
	result := Order{}
	filter := json.D{primitive.E{Key: "ID", Value: id}}

	// return order object or ErrOrderNotFound
	ordersCollection := i.client.Database(i.DB).Collection(i.c_orders)
	err = ordersCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, ErrOrderNotFound
	}
	return result, nil
}

////////////////////////////////////////////////////////////////////////////////

// GetOrders should return all orders with the given status. If status is the
// special -1 value then it should return all orders regardless of their status.
func (i *Instance) GetOrders(ctx context.Context, status OrderStatus) ([]Order, error) {

	// setup json filter query for fetching all documents
	orderFilter := json.D{{}}
	var orders []Order
	if err != nil {
		log.Fatal("GetOrders failed db", err)
		return orders, err
	}

	orderCollection := i.client.Database(i.DB).Collection(i.c_orders)
	cur, findError := orderCollection.Find(context.TODO(), orderFilter)
	if findError != nil {
		log.Fatal("GetOrders failed performing find", findError)
		return orders, findError
	}

	// use a slice for mapping results since size varies
	for cur.Next(context.TODO()) {
		var t Order
		mapError := cur.Decode(&t)
		if mapError != nil {
			return orders, mapError
		}

		if status == -1 || status == t.Status {
			orders = append(orders, t)
		}
	}

	// close cursor when through orders
	cur.Close(context.TODO())
	if len(orders) == 0 {
		return orders, ErrOrderNotFound
	}

	return orders, nil
}

////////////////////////////////////////////////////////////////////////////////

// SetOrderStatus should update the order with the given ID and set the status
// field. If that ID isn't found then the special ErrOrderNotFound error should
// be returned.
func (i *Instance) SetOrderStatus(ctx context.Context, id string, status OrderStatus) error {

	//Define filter query for fetching specific document from collection
	filter := json.D{primitive.E{Key: "ID", Value: id}}

	//Define updater for to specifiy change to be updated.
	updater := json.D{primitive.E{Key: "$set", Value: json.D{
		primitive.E{Key: "completed", Value: true},
	}}}

	orderCollection := i.client.Database(i.client.DB).Collection(i.c_orders)

	//Perform UpdateOne operation & validate against the error.
	_, err = orderCollection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		log.Fatal(ErrOrderNotFound, err)
		return ErrOrderNotFound
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

// InsertOrder should fill in the order's ID with a unique identifier if it's not
// already set and then insert it into the database. It should return the order's
// ID. If the order already exists then ErrOrderExists should be returned.
func (i *Instance) InsertOrder(ctx context.Context, order Order) (string, error) {

	// TODO: if the order's ID field is empty, generate a random ID, then insert

	if err != nil {
		log.Fatal(err)
		return err
	}
	orderCollection := i.client.Database(dbAdapter.DB).Collection(dbAdapter.C_ORDERS)
	insertResult, insertError := orderCollection.InsertOne(context.TODO(), order)
	if err != nil {
		log.Fatal(ErrOrderExists, insertError)
		return insertResult.ID, ErrOrderExists
	}
	return insertResult.id, nil
}
