package kafka

type PaymentEvent struct {
	EventType string         `json:"eventType"`
	Payload   PaymentPayload `json:"payload"`
}

type PaymentPayload struct {
	Id              int64  `json:"id"`
	BranchPaymentId string `json:"branchPaymentId"`
	Status          string `json:"status"`
}
