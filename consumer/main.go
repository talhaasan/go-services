package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
)

func main() {
	// Get the connection string from the environment variable
	url := os.Getenv("AMQP_URL")

	//If it doesnt exist, use the default connection string
	if url == "" {
		url = "amqp://guest:guest@localhost:5672"
	}

	// Connect to the rabbitMQ instance
	connection, err := amqp.Dial(url)

	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}

	// Create a channel from the connection. We'll use channels to access the data in the queue rather than the
	// connection itself
	channel, err := connection.Channel()

	if err != nil {
		panic("could not open RabbitMQ channel:" + err.Error())
	}

	// We create an exahange that will bind to the queue to send and receive messages
	err = channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	// We consume data from the queue named Test using the channel we created in go.
	msgs, err := channel.Consume("test", "", false, false, false, false, nil)

	if err != nil {
		panic("error consuming the queue: " + err.Error())
	}

	// We loop through the messages in the queue and print them in the console.
	// The msgs will be a go channel, not an amqp channel
	for msg := range msgs {
		fmt.Println("message received: " + string(msg.Body))
		msg.Ack(false)
	}

	// We close the connection after the operation has completed.
	defer connection.Close()
}