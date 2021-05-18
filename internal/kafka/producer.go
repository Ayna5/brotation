package kafka

import (
	"errors"

	"github.com/Shopify/sarama"
)

// Producer is a wrapper for sarama sync producer.
type Producer struct {
	topic    string
	producer sarama.SyncProducer
}

// New creates new producer.
func New(broker string, topic string) (*Producer, error) {
	if broker == "" {
		return nil, errors.New("empty broker") //nolint:errcheck,govet
	}
	if topic == "" {
		return nil, errors.New("empty topic") //nolint:errcheck,govet
	}

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		return nil, errors.New("could not create kafka producer") //nolint:errcheck,govet
	}

	return &Producer{
		topic:    topic,
		producer: producer,
	}, nil
}

// Send sends arbitrary message.
func (kp *Producer) Send(msg []byte) error {
	kafkaMsg := &sarama.ProducerMessage{
		Topic: kp.topic,
		Value: sarama.ByteEncoder(msg),
	}

	_, _, err := kp.producer.SendMessage(kafkaMsg)
	if err != nil {
		return errors.New("could not send message") //nolint:errcheck,govet
	}

	return nil
}

// Close tries to close the producer.
func (kp *Producer) Close() error {
	return kp.producer.Close()
}
