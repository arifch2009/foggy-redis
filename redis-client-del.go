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

 
	result, err := client.Del("ubuntu:latest").Result()
	if err != nil {
		fmt.Println("There was an error")
	}	

	if result == 1 {
		fmt.Println("The key has been deleted")
	} else {
		fmt.Println("The key is not present.")
	}
}


