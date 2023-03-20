package main

import (
	"fmt"
	"listener/event"
	"log"
	"math"
	"os"

	/* The reason we declared an alias for this is this package replaces a community developed package and people using this package,
	still reference it using amqp.*/
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func main() {
	// try to connect to rabbitmq
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer rabbitConn.Close()

	// start listening for messages
	log.Println("Listening for and consuming RabbitMQ messages...")

	// create consumer(consumes messages from the queue)
	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		// if we can't connect to rabbit, we can't go any further
		panic(err)
	}

	// watch the queue and consume events
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// don't continue until rabbit is ready
	for {
		/* The host after @, is whatever we call the `rabbitmq` service in our docker-compose.yaml . */
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			log.Println("Connected ro RabbitMQ!")
			connection = c
			break
		}

		/* We don't want to run the loop endlessly, so put a limit. In other words if we can't connect after 5 times, sth is wrong */
		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		// increase the delay each time we backoff. backoff gets bigger each time
		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(backOff)

		// we don't need continue here BTW!
		continue
	}

	return connection, nil
}
