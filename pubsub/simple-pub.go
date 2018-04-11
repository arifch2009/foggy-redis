package main

import "github.com/go-redis/redis"



func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})


	err := client.Publish("mychannel1", "hello").Err()
	if err != nil {
		panic(err)
	}

	
}

