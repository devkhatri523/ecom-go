package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"v01/domain"
)

type CustomerRepositoryImpl struct {
	db *mongo.Database
}

func NewCustomerRepositoryImpl(db *mongo.Database) CustomerRepository {
	return &CustomerRepositoryImpl{
		db: db,
	}
}

func (c CustomerRepositoryImpl) CreateCustomer(customer domain.Customer) (string, error) {
	collection := c.db.Collection("customers")
	res, err := collection.InsertOne(context.Background(), customer)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (c CustomerRepositoryImpl) FindAllCustomer() ([]domain.Customer, error) {
	collection := c.db.Collection("customers")
	var customers []domain.Customer
	filter := bson.D{}

	// Find all documents
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	// Iterate through the cursor and decode each document
	for cur.Next(context.Background()) {
		var customer domain.Customer
		err = cur.Decode(&customer)
		if err != nil {
			log.Fatal(err)
		}
		customers = append(customers, customer)
	}

	if err = cur.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

func (c CustomerRepositoryImpl) UpdateCustomer(customerId string, customer domain.Customer) (string, error) {
	collection := c.db.Collection("customers")
	oid, err := primitive.ObjectIDFromHex(customerId)
	filter := bson.M{"_id": oid}
	if err != nil {
		errors.New(fmt.Sprintf("Error while updating customer with Id :%s", customerId))
	}
	update := bson.M{
		"$set": bson.M{
			"first_name": customer.FirstName,
			"last_name":  customer.LastName,
			"email":      customer.Email,
			"address": domain.Address{
				Street:      customer.Address.Street,
				HouseNumber: customer.Address.HouseNumber,
				ZipCode:     customer.Address.ZipCode,
			},
		},
	}
	results := collection.FindOneAndUpdate(context.Background(), filter, update)

	if results.Err() != nil {
		return "", results.Err()
	}
	return customerId, nil
}

func (c CustomerRepositoryImpl) IsCustomerExists(customerId string) (bool, error) {
	collection := c.db.Collection("customers")
	var customer domain.Customer
	oid, err := primitive.ObjectIDFromHex(customerId)
	if err != nil {
		log.Fatalf("Invalid customer ID: %v", err)
	}
	filter := bson.M{"_id": oid}
	err = collection.FindOne(context.Background(), filter).Decode(&customer)
	if err != nil {
		return false, err
	}
	return customer.Id.Hex() != "", nil
}

func (c CustomerRepositoryImpl) FindById(customerId string) (domain.Customer, error) {
	collection := c.db.Collection("customers")
	var customer domain.Customer
	oid, err := primitive.ObjectIDFromHex(customerId)
	if err != nil {
		log.Fatalf("Invalid customer ID: %v", err)
	}
	filter := bson.M{"_id": oid}
	err = collection.FindOne(context.Background(), filter).Decode(&customer)
	if err != nil {
		return domain.Customer{}, err
	}
	return customer, nil
}

func (c CustomerRepositoryImpl) DeleteCustomer(customerId string) (string, error) {
	collection := c.db.Collection("customers")
	var customer domain.Customer
	oid, err := primitive.ObjectIDFromHex(customerId)
	if err != nil {
		return "", err
	}
	filter := bson.M{"_id": oid}
	results := collection.FindOneAndDelete(context.Background(), filter)
	if results.Err() != nil {
		return "", results.Err()
	}
	return customer.Id.Hex(), nil
}
