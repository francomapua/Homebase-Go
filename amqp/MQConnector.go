// Version 2

package commons

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

const (
	connectionAttemptLimit = 5
	actionAttemptLimit     = 5
)

// Struct
type mqConnector struct {
	uri  amqp.URI
	conn *amqp.Connection
	ch   *amqp.Channel

	connectionAttempts int
	actionAttemtps     int
}

// NewMQConnector : Connects to an AMQP server and facilitates consumption and publishing of messages
// Constructor that returns an instance of messageQueueConstructor
func NewMQConnector(vhost, username, password, host string, port int) *mqConnector {
	// Check for empty vHost
	if vhost == "" {
		vhost = "/"
	}

	// Check for empty port
	if port == 0 {
		port = 5672 // Default Port
	}

	mqc := new(mqConnector)
	mqc.uri = createURI(vhost, username, password, host, port)

	return mqc
}

// PUBLIC
// ConsumeQueue : Begins Consuming a Queue
func (mqc mqConnector) ConsumeQueue(queueName string, collection <-chan amqp.Delivery) error {
	err := mqc.ensureQueue(queueName)
	if err != nil {
		return err
	}

	// Begin Consuming Queue
	collection, err = mqc.ch.Consume(
		queueName,      // queue
		"mcw-splitter", // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	return err
}

func (mqc mqConnector) PublishToQueue(queueName string, payload string) error {
	err := mqc.ensureQueue(queueName)
	if err != nil {
		return err
	}

	err = mqc.ch.Publish(
		"go-test-exchange", // exchange
		"go-test-key",      // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Transient,
			ContentType:  "application/json",
			Body:         []byte(payload),
			Timestamp:    time.Now(),
		})
	return err
}

// Disconnect : Disconnects the Connector it is connected
func (mqc mqConnector) Disconnect() {
	if mqc.conn != nil {
		mqc.conn.Close()
	}
}

// PRIVATE
func (mqc mqConnector) connect() error {
	var err error

	// Initiate Connection
	mqc.conn, err = amqp.Dial(mqc.uri.String())
	if err != nil {
		return err
	}

	// Connect to Channel
	mqc.ch, err = mqc.conn.Channel()
	return err
}

func (mqc mqConnector) ensureQueue(queueName string) error {
	if mqc.ch == nil {
		err := executeWithRetry(connectionAttemptLimit, 2000, mqc.connect)
		if err != nil {
			return err
		}
	}

	_, err := mqc.ch.QueueDeclare(
		queueName, // name, leave empty to generate a unique name
		true,      // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	return err
}

type runnable func() error

func executeWithRetry(retryCount, delayMS int, fn runnable) error {
	fmt.Println("Executing Runnable...")
	var err error
	for i := 0; i < retryCount; i++ {
		err = fn()
		if err == nil {
			break
		}
		millis := time.Duration(int64(delayMS))
		time.Sleep(millis * time.Millisecond)
	}
	return err
}

func createURI(vhost, username, password, host string, port int) amqp.URI {
	uri := amqp.URI{}
	uri.Scheme = "amqp"
	uri.Vhost = vhost
	uri.Username = username
	uri.Password = password
	uri.Host = host
	uri.Port = port
	return uri
}
