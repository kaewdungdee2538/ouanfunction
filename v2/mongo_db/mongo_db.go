package mongo_db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(uri string) (*mongo.Client, error) {
	// Set up the MongoDB client options
	clientOptions := options.Client().ApplyURI(uri)

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

// the result parameter is address
func FindData(client *mongo.Client, dbName string, collectionName string, filter interface{}, result interface{}) error {
	// Access  from the client
	collection := client.Database(dbName).Collection(collectionName)

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return err
	}
	defer cur.Close(context.TODO())
	// get data array to slice
	err = cur.All(context.TODO(), result)

	res := result
	fmt.Println(res)
	if err != nil {
		return err
	}

	return nil
}
