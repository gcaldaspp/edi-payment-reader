package payment_service

import (
	kafka_producer "github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/kafka/producer"
	payment_domain "github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/payment/domain"
)

type PaymentService struct {
	repo     payment_domain.PaymentRepository
	producer kafka_producer.Producer
}

func NewPaymentService(r payment_domain.PaymentRepository, p kafka_producer.Producer) *PaymentService {
	return &PaymentService{
		repo:     r,
		producer: p,
	}
}

func (ps *PaymentService) Execute(payment *payment_domain.Payment) error {
	err := ps.repo.Create(payment)
	if err != nil {
		return err
	}
	return ps.producer.Send(payment)
}
