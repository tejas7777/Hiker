package models

import (
	"context"
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetHash(hash string, field string, value string) error {

	conn := Pool.Get()
	defer conn.Close()

	_, err := conn.Do("hset", hash, field, value)

	if err != nil {
		return err
	}

	return nil

}

func GetHash(hash string, field string) error {

	conn := Pool.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("HGET", hash, field))

	if err != nil {
		return err
	}

	return nil
}

func ChechTrailAllowed(trail string) (bool, error) {
	return true, nil
}

func CheckKey(key string) (bool, error) {
	conn := Pool.Get()
	defer conn.Close()
	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return ok, err
	}
	return ok, err
}

func GetAllTrails() (map[string]string, error) {
	client, err := GetMongoClient()

	if err != nil {
		fmt.Errorf(err.Error())
		return nil, err
	}

	defer Disconnect()

	var result []struct {
		Park_Name string `bson:"park_name"`
	}

	collection := client.Database("Hiker").Collection("trails")

	cursor, err := collection.Find(context.TODO(), bson.M{}, options.Find().SetProjection(bson.M{"Name": 1}))

	if err != nil {
		panic(err)
		return nil, err
	}

	cursor.Decode(&result)

	//Take out the value from the crusor in batches
	var smap = make(map[string]string)
	for cursor.Next(context.TODO()) {
		var trail bson.M
		if err = cursor.Decode(&trail); err != nil {
			log.Fatal(err)
			return nil, err
		}
		fmt.Println(trail["Name"])
		if trail["_id"] == nil {
			fmt.Println("NIL MAP")
		}
		id := trail["_id"]
		sid := id.(primitive.ObjectID).Hex()
		park := fmt.Sprint(trail["Name"])
		smap[sid] = park
	}
	fmt.Println(smap)
	return smap, err
}
