package main

import "github.com/go-redis/redis"
import "fmt"
import "strings"


func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val := client.Get("ubuntu:latest").String()
	

	if strings.Contains(val , "nil") {
		fmt.Println("The record does not exist.")
	} else {
		fmt.Println(val)
	}
 
	
	
 
}


