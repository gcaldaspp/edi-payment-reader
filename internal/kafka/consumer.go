package kafka

import "github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/payment"

type PaymentConsumer struct {
	useCase payment.CreatePaymentUseCase
}

func NewPaymentConsumer(useCase payment.CreatePaymentUseCase) *PaymentConsumer {
	return &PaymentConsumer{
		useCase: useCase,
	}
}

func (pc *PaymentConsumer) Process(event PaymentEvent) {
	p := ToDomain(&event)
	pc.useCase.Execute(p)
}
