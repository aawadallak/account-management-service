package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"
)

type kafkaProducer struct {
	writer *kafka.Writer
}

func (k *kafkaProducer) Close() {
	k.writer.Close()
}

func (k *kafkaProducer) Producer(value []byte, headerMessage string) error {

	err := k.writer.WriteMessages(context.Background(), kafka.Message{
		Value: value,
		Headers: []protocol.Header{
			{
				Key: headerMessage,
			},
		},
	})

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func NewKafkaProducer(topic string) *kafkaProducer {
	return &kafkaProducer{
		writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers: []string{"kafka:9093"},
			Topic:   topic,
		}),
	}
}
