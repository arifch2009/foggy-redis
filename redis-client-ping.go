package main

import "github.com/go-redis/redis"
import "fmt"

//import "context"

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

This text is written in Emacs
This text is written in Emacs

