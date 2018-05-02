package trigger

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/Sirupsen/logrus"
	"github.com/streadway/amqp"
)

const (
	DefaultConsumerTag = "dkron"
)

type rabbitMQ struct {
	AbstractTrigger
	configuration              *Configuration
	brokerConn                 *amqp.Connection
	brokerQueue                amqp.Queue
	brokerChannel              *amqp.Channel
	replies                    <-chan amqp.Delivery
	brokerInputMessagesChannel <-chan amqp.Delivery
	done                       chan error
}

func newRabbitMQTrigger(logger *logrus.Logger, configuration *Configuration) (Trigger, error) {
	return &rabbitMQ{
		AbstractTrigger: AbstractTrigger{
			Logger: logger,
		},
		configuration: configuration,
	}, nil
}

func (r *rabbitMQ) Initialize() error {
	return nil
}

func (r *rabbitMQ) createBrokerResources() error {

	var err error
	r.brokerConn, err = amqp.Dial(r.configuration.RabbitMQ.URI)
	if err != nil {
		return errors.Wrapf(err, "Failed to create connection. URI %s", r.configuration.RabbitMQ.URI)
	}

	go func() {
		r.Logger.Infof("closing: %s", <-r.brokerConn.NotifyClose(make(chan *amqp.Error)))
	}()

	r.Logger.Debugln("Connected to broker", "brokerUrl", r.configuration.RabbitMQ.URI)
	r.brokerChannel, err = r.brokerConn.Channel()
	if err != nil {
		return errors.Wrapf(err, "Failed to create channel. URI %s", r.configuration.RabbitMQ.URI)
	}

	r.Logger.Debug("Created broker channel")
	if err = r.brokerChannel.ExchangeDeclare(
		r.configuration.RabbitMQ.ExchangeName,
		"topic",
		true,  // durable
		false, // delete when complete
		false, // internal
		false, // noWait
		nil,   // arguments
	); err != nil {
		return errors.Wrap(err, "Failed to declare exchange")
	}

	r.Logger.Debugln("Declared exchange", "exchangeName", r.configuration.RabbitMQ.ExchangeName)
	queue, err := r.brokerChannel.QueueDeclare(
		r.configuration.RabbitMQ.QueueName,
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		return errors.Wrap(err, "Failed to declare queue")
	}

	r.Logger.Debugln("Declared queue", "queueName", queue)

	for _, topic := range r.configuration.RabbitMQ.Topics {
		err = r.brokerChannel.QueueBind(
			r.brokerQueue.Name, // queue name
			topic,              // routing key
			r.configuration.RabbitMQ.ExchangeName, // exchange
			false,
			nil)
		if err != nil {
			return errors.Wrap(err, "Failed to bind to queue")
		}

		r.Logger.Debugln("Bound queue to topic",
			"queueName", r.brokerQueue.Name,
			"topic", topic,
			"exchangeName", r.configuration.RabbitMQ.ExchangeName)

	}

	r.brokerInputMessagesChannel, err = r.brokerChannel.Consume(
		queue.Name, // name
		"",         // consumerTag,
		false,      // noAck
		false,      // exclusive
		false,      // noLocal
		false,      // noWait
		nil,        // arguments
	)
	if err != nil {
		return errors.Wrap(err, "Failed to start consuming messages")
	}

	r.Logger.Debugln("Starting consumption from queue", "queueName", r.brokerQueue.Name)
	return err
}

func (r *rabbitMQ) setEmptyParameters() {
	r.configuration.RabbitMQ.ExchangeName = "amq.topic"
	if r.configuration.RabbitMQ.ConsumerTag == "" {
		r.configuration.RabbitMQ.ConsumerTag = DefaultConsumerTag
	}

	if len(r.configuration.RabbitMQ.Topics) == 0 {
		r.configuration.RabbitMQ.Topics = []string{"*"}
	}
}

// Start as door of trigger
func (r *rabbitMQ) Start() error {

	r.setEmptyParameters()
	err := r.createBrokerResources()
	if err != nil {
		return err
	}
	// start listening for published messages
	go r.handleBrokerMessages()

	return nil
}

func (r *rabbitMQ) Stop() error {
	if err := r.brokerChannel.Cancel(r.configuration.RabbitMQ.ConsumerTag, true); err != nil {
		return fmt.Errorf("Consumer cancel failed: %s", err)
	}

	if err := r.brokerConn.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}

	defer r.Logger.Infoln("RabbitMQ shutdown OK")
	return <-r.done
}

func (r *rabbitMQ) handleBrokerMessages() {
	for message := range r.brokerInputMessagesChannel {
		job := jobMessagePool.Get().(*JobMessage)
		json.Unmarshal(message.Body, &job)
		message.Ack(false)
		r.Logger.Infof("handle job %+v", job)
		jobMessagePool.Put(job)
	}
	r.done <- nil
}
