package event

import amqp "github.com/rabbitmq/amqp091-go"

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", // name of the exchange
		"topic",      // type
		true,         // we want this exchange to hang around as long as this is running(is it durable)
		false,        // do you get rid of it when you're done with it (auto-deleted?)
		false,        // is this an exchange that's gonna used internally? No, we'll be using it between microservices(internal?)
		false,
		nil,
	)
}

// get a random queue
func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",
		false, // get rid of it when we're done with it, so it's not durable
		false, // do we delete it when it's unused. No so pass false
		true,
		false,
		nil,
	)
}
