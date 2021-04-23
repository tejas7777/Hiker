package models

import (
	"github.com/gomodule/redigo/redis"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
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
