package mongo_db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, error) {
	// Set up the MongoDB client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to the MongoDB server
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Connection successful
	fmt.Println("Connected to MongoDB successful")

	return client, nil
}


func CreateConnectionPool(uri string, maxConnections int) (*mongo.Client, error) {
	// Set up the MongoDB client options
	clientOptions := options.Client().ApplyURI(uri)

	// Create the connection pool
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	// Connect to the MongoDB server
	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	// Set the maximum number of connections in the pool
	
	// Connection successful
	fmt.Println("Connected to MongoDB successful")

	return client, nil
}


func CloseDb(client *mongo.Client) error {
	err := client.Disconnect(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func InsertOne(client *mongo.Client, dbName string, collectionName string, documents interface{}) (*mongo.InsertOneResult, error) {
	// Insert the document
	result, err := client.Database(dbName).Collection(collectionName).InsertOne(context.Background(), documents)

	// defer CloseDb(client)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func InsertMany(client *mongo.Client, dbName string, collectionName string, documents []interface{}) (*mongo.InsertOneResult, error) {
	// Insert the document
	result, err := client.Database(dbName).Collection(collectionName).InsertOne(context.Background(), documents)

	// defer CloseDb(client)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func FindData(client *mongo.Client, dbName string, collectionName string, filter interface{}) ([]interface{}, error) {
	// Access  from the client
	collection := client.Database(dbName).Collection(collectionName)

	// Define an empty slice to store the results
	var results []interface{}

	// Find documents based on the filter
	cursor, err := collection.Find(context.Background(), filter)

	// defer CloseDb(client)

	if err != nil {
		return nil, err
	}

	// Iterate through the cursor and append the documents to the results slice
	for cursor.Next(context.Background()) {
		var document interface{}
		err := cursor.Decode(&document)
		if err != nil {
			return nil, err
		}
		results = append(results, document)
	}

	return results, nil
}
