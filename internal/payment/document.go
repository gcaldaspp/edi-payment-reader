package payment

type PaymentDocument struct {
	Id              int64  `bson:"_id"`
	BranchPaymentId string `bson:"branchPaymentId"`
	Status          string `bson:"status"`
}

func FromDomain(payment *Payment) *PaymentDocument {
	return &PaymentDocument{
		Id:              payment.Id,
		BranchPaymentId: payment.BranchPaymentId,
		Status:          payment.Status,
	}
}
