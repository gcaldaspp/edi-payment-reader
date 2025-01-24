package kafka_consumer

import payment_domain "github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/payment/domain"

type PaymentConsumer struct {
	useCase payment_domain.CreatePaymentUseCase
}

func NewPaymentConsumer(useCase payment_domain.CreatePaymentUseCase) *PaymentConsumer {
	return &PaymentConsumer{
		useCase: useCase,
	}
}

func (pc *PaymentConsumer) Process(event PaymentEvent) {
	p := ToDomain(&event)
	pc.useCase.Execute(p)
}
