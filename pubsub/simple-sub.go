package main

import "github.com/go-redis/redis"
import "fmt"
import "time"
//import "strings"


func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pubsub := client.Subscribe("mychannel1")
	defer pubsub.Close()

	// Wait for subscription to be created before publishing message.
	subscr, err := pubsub.ReceiveTimeout(time.Second) 
	if err != nil {
		panic(err)
	}
	fmt.Println(subscr)

	msg, err := pubsub.ReceiveMessage()
	if err != nil {
		panic(err)
	}

	fmt.Println(msg.Channel, msg.Payload)
	// Output: subscribe: mychannel1
	// mychannel1 hello
}

