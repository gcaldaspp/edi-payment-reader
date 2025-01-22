package kafka

type PaymentProducer struct {
	topicName string
}

func NewPaymentProducer(topicName string) *PaymentProducer {
	return &PaymentProducer{
		topicName: topicName,
	}
}

func (pp *PaymentProducer) Send(event *PaymentEvent) {
}
