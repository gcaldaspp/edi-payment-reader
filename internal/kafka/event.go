package kafka

import (
	"github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/payment"
	"github.com/PicPay/ms-edi-wrk-payment-reader-go/pkg/types"
)

type PaymentEvent struct {
	EventType *string         `json:"eventType"`
	Payload   *PaymentPayload `json:"payload"`
}

type PaymentPayload struct {
	Id                 *int64          `json:"id"`
	BranchPaymentId    *string         `json:"branchPaymentId"`
	Status             *string         `json:"status"`
	BranchId           *int64          `json:"branchId"`
	ProductId          *int64          `json:"productId"`
	ProductType        *string         `json:"productType"`
	PaymentAmount      *float64        `json:"paymentAmount"`
	PaymentDate        *types.DateOnly `json:"paymentDate"`
	CurrencyIsoCode    *int64          `json:"currencyIsoCode"`
	PrimaryAccountName *string         `json:"primaryAccountName"`
	AcquirerId         *int64          `json:"acquirerId"`
	BranchName         *string         `json:"branchName"`
	BankAccount        *string         `json:"bankAccount"`
	BankAgencyCode     *string         `json:"bankAgencyCode"`
	BankCode           *string         `json:"bankCode"`
	BankIspb           *int64          `json:"bankIspb"`
	BankAccountType    *string         `json:"bankAccountType"`
	BranchDocumentType *string         `json:"branchDocumentType"`
	BranchDocument     *string         `json:"branchDocument"`
	DocumentType       *string         `json:"documentType"`
	Document           *string         `json:"document"`
	PaymentType        *string         `json:"paymentType"`
	MovementCounter    *int64          `json:"movementCounter"`
	Prepayment         *string         `json:"prepayment"`
	PaymentOrigin      *string         `json:"paymentOrigin"`
	PaymentAccountId   *string         `json:"paymentAccountId"`
}

func ToDomain(p *PaymentEvent) *payment.Payment {
	return &payment.Payment{
		Id:                 p.Payload.Id,
		BranchPaymentId:    p.Payload.BranchPaymentId,
		BranchId:           p.Payload.BranchId,
		ProductId:          p.Payload.ProductId,
		ProductType:        p.Payload.ProductType,
		PaymentAmount:      p.Payload.PaymentAmount,
		PaymentDate:        p.Payload.PaymentDate,
		CurrencyIsoCode:    p.Payload.CurrencyIsoCode,
		Status:             p.Payload.Status,
		PrimaryAccountName: p.Payload.PrimaryAccountName,
		AcquirerId:         p.Payload.AcquirerId,
		BranchName:         p.Payload.BranchName,
		BankAccount:        p.Payload.BankAccount,
		BankAgencyCode:     p.Payload.BankAgencyCode,
		BankCode:           p.Payload.BankCode,
		BankIspb:           p.Payload.BankIspb,
		BankAccountType:    p.Payload.BankAccountType,
		BranchDocumentType: p.Payload.BranchDocumentType,
		BranchDocument:     p.Payload.BranchDocument,
		DocumentType:       p.Payload.DocumentType,
		Document:           p.Payload.Document,
		PaymentType:        p.Payload.PaymentType,
		MovementCounter:    p.Payload.MovementCounter,
		Prepayment:         p.Payload.Prepayment,
		PaymentOrigin:      p.Payload.PaymentOrigin,
		EventType:          p.EventType,
	}
}
