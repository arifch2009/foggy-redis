package main

import 

func main() {
	pubsub := client.Subscribe("mychannel1")
	defer pubsub.Close()

	// Wait for subscription to be created before publishing message.
	subscr, err := pubsub.ReceiveTimeout(time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println(subscr)

	err = client.Publish("mychannel1", "hello").Err()
	if err != nil {
		panic(err)
	}

	msg, err := pubsub.ReceiveMessage()
	if err != nil {
		panic(err)
	}

	fmt.Println(msg.Channel, msg.Payload)
	// Output: subscribe: mychannel1
	// mychannel1 hello
}
