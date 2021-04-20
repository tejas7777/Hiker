package models

import (
	"github.com/gomodule/redigo/redis"
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
