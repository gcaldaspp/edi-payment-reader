package kafka_producer

type Producer interface {
	Send(event any) error
}
