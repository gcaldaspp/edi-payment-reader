package kafka_producer

import payment_domain "github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/payment/domain"

type CompilerEvent struct {
	PaymentId          *int64   `json:"paymentId"`
	BranchPaymentId    *string  `json:"branchPaymentId"`
	PaymentDate        *string  `json:"paymentDate"`
	BankAccount        *string  `json:"bankAccount"`
	BankAgencyCode     *string  `json:"bankAgencyCode"`
	BankCode           *string  `json:"bankCode"`
	BankAccountType    *string  `json:"bankAccountType"`
	BankIspb           *int64   `json:"bankIspb"`
	Status             *string  `json:"status"`
	PaymentTypeId      *string  `json:"paymentTypeId"`
	BranchDocument     *string  `json:"branchDocument"`
	BranchId           *int64   `json:"branchId"`
	ProductId          *int64   `json:"productId"`
	ProductType        *string  `json:"productType"`
	PaymentAmount      *float64 `json:"paymentAmount"`
	CurrencyIsoCode    *int64   `json:"currencyIsoCode"`
	PrimaryAccountName *string  `json:"primaryAccountName"`
	AcquirerId         *int64   `json:"acquirerId"`
	BranchName         *string  `json:"branchName"`
	BranchDocumentType *string  `json:"branchDocumentType"`
	DocumentType       *string  `json:"documentType"`
	Document           *string  `json:"document"`
	MovementCounter    *int64   `json:"movementCounter"`
	Prepayment         *string  `json:"prepayment"`
	PaymentOrigin      *string  `json:"paymentOrigin"`
	PaymentAccountId   *string  `json:"paymentAccountId"`
	EventType          *string  `json:"eventType"`
}

func FromDomain(p payment_domain.Payment) *CompilerEvent {
	return &CompilerEvent{
		PaymentId:          p.Id,
		BranchPaymentId:    p.BranchPaymentId,
		PaymentDate:        p.PaymentDateString(),
		BankAccount:        p.BankAccount,
		BankAgencyCode:     p.BankAgencyCode,
		BankCode:           p.BankCode,
		BankAccountType:    p.BankAccountType,
		BankIspb:           p.BankIspb,
		Status:             p.Status,
		PaymentTypeId:      p.PaymentType,
		BranchDocument:     p.BranchDocument,
		BranchId:           p.BranchId,
		ProductId:          p.ProductId,
		ProductType:        p.ProductType,
		PaymentAmount:      p.PaymentAmount,
		CurrencyIsoCode:    p.CurrencyIsoCode,
		PrimaryAccountName: p.PrimaryAccountName,
		AcquirerId:         p.AcquirerId,
		BranchName:         p.BranchName,
		BranchDocumentType: p.BranchDocumentType,
		DocumentType:       p.DocumentType,
		Document:           p.Document,
		MovementCounter:    p.MovementCounter,
		Prepayment:         p.Prepayment,
		PaymentOrigin:      p.PaymentOrigin,
		PaymentAccountId:   p.PaymentAccountId,
		EventType:          p.EventType,
	}
}
