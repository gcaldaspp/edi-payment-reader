package payment

type PaymentService struct {
	repo PaymentRepository
}

func NewPaymentService(r PaymentRepository) *PaymentService {
	return &PaymentService{
		repo: r,
	}
}

func (ps *PaymentService) Execute(payment *Payment) error {
	return ps.repo.Create(payment)
}
