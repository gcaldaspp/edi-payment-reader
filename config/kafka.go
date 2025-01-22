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
		"security.protocol": viper.GetString("KAFKA_PROTOCOL"),
		"sasl.mechanisms":   viper.GetString("KAFKA_MECHANISMS"),
		"sasl.username":     viper.GetString("KAFKA_USER"),
		"sasl.password":     viper.GetString("KAFKA_PASSWORD"),
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
		slog.Error("There is an error on connect in kafka broker", "error", err)
		return
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
