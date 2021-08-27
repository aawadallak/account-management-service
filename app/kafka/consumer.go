package kafka

import (
	"context"
	"latest/config"

	"github.com/segmentio/kafka-go"
)

type kafkaConsumer struct {
	reader *kafka.Reader
}

func (k *kafkaConsumer) Consumer() (string, error) {

	ctx := context.Background()

	v, err := k.reader.ReadMessage(ctx)

	if err != nil {
		return "", err
	}

	return string(v.Value), nil

}

func NewKafkaConsumer(topic string) *kafkaConsumer {
	return &kafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{config.GetConfig().KafkaBroker},
			GroupID: topic,
			Topic:   topic}),
	}
}
