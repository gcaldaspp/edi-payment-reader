package payment_domain

type Writer interface {
	Create(payment *Payment) error
}

type PaymentRepository interface {
	Writer
}

type CreatePaymentUseCase interface {
	Execute(payment *Payment) error
}
