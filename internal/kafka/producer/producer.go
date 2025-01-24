package kafka_producer

import (
	"github.com/PicPay/ms-edi-wrk-payment-reader-go/config"
	"github.com/spf13/viper"
)

type PaymentProducer struct {
	topicName string
}

func NewPaymentProducer() *PaymentProducer {
	return &PaymentProducer{
		topicName: viper.GetString("KAFKA_COMPILER_TOPIC"),
	}
}

func (pp *PaymentProducer) Send(event any) error {
	return config.ProduceMessage(pp.topicName, event)
}
