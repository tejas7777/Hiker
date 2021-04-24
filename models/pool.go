package models

import (
	"context"
	"fmt"
	"log"

	"os"

	"github.com/gomodule/redigo/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Pool = &redis.Pool{
	// Maximum number of idle connections in the pool.
	MaxIdle: 10,
	// max number of connections
	MaxActive: 100,
	// Dial is an application supplied function for creating and
	// configuring a connection.
	Dial: func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", ":6379")
		if err != nil {
			panic(err.Error())
		}
		return c, err
	},
}

//Initialized and exposed through  GetMongoClient().*/
var clientInstance *mongo.Client

//Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error

//Used to execute client creation procedure only once.
//var mongoOnce sync.Once

//I have used below constants just to hold required database config's.
const (
	USERNAME = "root"
)

//GetMongoClient - Return mongodb connection to work with
func GetMongoClient() (*mongo.Client, error) {
	//Perform connection creation operation only once.
	PASSWORD := os.Getenv("PASSWORD")

	//----------NOT USING GO SYNC FOR NOW-------------------------
	// mongoOnce.Do(func() {
	// 	// Set client options
	// 	connectionstring := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.pgmtq.mongodb.net/myFirstDatabase?retryWrites=true&w=majority", USERNAME, PASSWORD)
	// 	clientOptions := options.Client().ApplyURI(connectionstring)
	// 	// Connect to MongoDB
	// 	ctx := context.TODO()
	// 	client, err := mongo.Connect(ctx, clientOptions)
	// 	if err != nil {
	// 		clientInstanceError = err
	// 	}
	// 	// Check the connection
	// 	err = client.Ping(ctx, nil)
	// 	if err != nil {
	// 		clientInstanceError = err
	// 	}
	// 	clientInstance = client
	// })
	//-------------------------------------------------------------------

	// Set client options
	connectionstring := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.pgmtq.mongodb.net/myFirstDatabase?retryWrites=true&w=majority", USERNAME, PASSWORD)
	clientOptions := options.Client().ApplyURI(connectionstring)
	// Connect to MongoDB
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		clientInstanceError = err
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		clientInstanceError = err
	}
	clientInstance = client

	return clientInstance, clientInstanceError
}

func Disconnect() error {
	if clientInstance == nil {
		return nil
	}

	err := clientInstance.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return err
	}

	// TODO optional you can log your closed MongoDB client
	fmt.Println("Connection to MongoDB closed.")
	return nil
}
