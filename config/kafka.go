package config

import (
	"encoding/json"
	"log/slog"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/viper"
)

func kafkaConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": viper.GetString("KAFKA_BROKERS"),
		"group.id":          viper.GetString("CONSUMER_GROUP"),
		"auto.offset.reset": "earliest",
	}
}

func ProduceMessage(topic string, event any) error {
	p, err := kafka.NewProducer(kafkaConfig())
	if err != nil {
		return err
	}
	defer p.Close()

	value, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: value,
	}, nil)
}

func StartKafkaConsumer(topic string, consumer func([]byte)) {
	c, err := kafka.NewConsumer(kafkaConfig())

	if err != nil {
		panic(err)
	}
	defer c.Close()

	c.SubscribeTopics([]string{topic}, nil)
	slog.Info("Subscribed in " + topic)

	for {
		slog.Info("Wait to receive events")
		evt, err := c.ReadMessage(-1)
		if err == nil {
			consumer(evt.Value)
		}
	}
}
