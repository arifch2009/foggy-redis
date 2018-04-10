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

        
	// Set a key value entry using Set()
	//err := client.Set("ubuntu:latest", "Pulling", 0).Err()
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("The value has been succesfully set")
	//}

	result := client.HSetNX("ubuntu:latest", "Pulling", 0).Val()
	
	if result {
		fmt.Println("The value has been succesfully set")
	} else {
		fmt.Println("The value is already present")
	}
}

