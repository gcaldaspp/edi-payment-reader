package payment

type Payment struct {
	Id              int64
	BranchPaymentId string
	Status          string
}

func (p *Payment) IsApproved() bool {
	return p.Status == "APPROVED"
}

type Writer interface {
	Create(payment *Payment) error
}

type PaymentRepository interface {
	Writer
}

type CreatePaymentUseCase interface {
	Execute(payment *Payment) error
}
