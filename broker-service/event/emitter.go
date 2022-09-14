package event

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Emiiter struct {
	connection *amqp.Connection
}

func (e *Emiiter) setup() error {
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()
	return declareExchange(channel)
}

func (e *Emiiter) Push(event string, severity string) error {
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	log.Println("Pushing to channel")

	err = channel.Publish(
		"logs_topic",
		severity,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func NewEventEmitter(conn *amqp.Connection) (Emiiter, error) {
	emitter := Emiiter{
		connection: conn,
	}

	err := emitter.setup()
	if err != nil {
		return Emiiter{}, err
	}

	return emitter, nil
}
