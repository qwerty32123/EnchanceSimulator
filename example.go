package main

import (
	"EnchanceSimulator/config"
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"io"

	"log"
	"net/http"
)

func main() {
	// Establish a connection to RabbitMQ
	conn, ch := establishRabbitMQConnection()
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {

		}
	}(ch)

	//Load Global Config Values
	appConfig := config.GetConfig()
	inputQueueName := appConfig.InputQueu
	outputQueueName := appConfig.OutputQueu
	declareQueues(ch, inputQueueName, outputQueueName)

	// Declare input and output queues
	inputQueueName = "input_data_queue"
	outputQueueName = "output_data_queue"
	declareQueues(ch, inputQueueName, outputQueueName)

	//	 Make an HTTP POST request and get the response bytes
	responseBytes, err := makeHTTPPostRequest()
	if err != nil {
		log.Println("Error:", err)
		return
	}

	// Convert the response bytes to hex (optional)
	hexData := hex.EncodeToString(responseBytes)

	// Publish the hex data to the input queue
	publishToInputQueue(ch, inputQueueName, hexData)

	log.Println("Data sent successfully.")

	//	 Consume the output queue
	consumeOutputQueue(ch, outputQueueName)
}

// Function to establish a connection to RabbitMQ

func establishRabbitMQConnection() (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial("amqp:guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}

	return conn, ch
}

//  Function to declare input and output queues

func declareQueues(ch *amqp.Channel, inputQueueName, outputQueueName string) {
	declareQueue(ch, inputQueueName)
	declareQueue(ch, outputQueueName)
}

//  Function to declare a queue

func declareQueue(ch *amqp.Channel, queueName string) {
	_, err := ch.QueueDeclare(
		queueName, //Queue name
		false,     //Durable
		false,     //Delete when unused
		false,     //Exclusive
		false,     //No-wait
		nil,       //Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare the queue '%s': %s", queueName, err)
	}
}

// Function to make an HTTP POST request

func makeHTTPPostRequest() ([]byte, error) {
	payload := map[string]interface{}{
		"keyType":      0,
		"mainCategory": 10,
		"subCategory":  3,
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("Error encoding JSON: %v", err)
	}

	//Todo change to get all strings from hashcorpi vault
	url := "secret/Trademarket/GetWorldMarketHotList"
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadJSON))
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %v", err)
	}

	req.Header.Set("User-Agent", "XXX")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error making request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response: %v", err)
	}

	return responseBytes, nil
}

// Function to publish data to the input queue

func publishToInputQueue(ch *amqp.Channel, queueName, data string) {
	err := ch.Publish(
		"",        // Exchange
		queueName, // Routing key
		false,     // Mandatory
		false,     // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish message to '%s' queue: %s", queueName, err)
	}
}

//Function to consume the output queue

func consumeOutputQueue(ch *amqp.Channel, queueName string) {
	messages, err := ch.Consume(
		queueName, //Queue name
		"",        // Consumer tag
		false,     // Auto-acknowledgement set to false
		false,     // Exclusive
		false,     // No-local
		false,     //No-wait
		nil,       //Arguments
	)
	if err != nil {
		log.Fatalf("Failed to consume messages from '%s' queue: %s", queueName, err)
	}

	for msg := range messages {
		log.Println("Received output:", string(msg.Body))
		err := msg.Ack(false)
		if err != nil {
			return
		}
	}
}
