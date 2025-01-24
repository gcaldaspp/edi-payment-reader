package payment_persistence

import (
	"time"

	payment_domain "github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/payment/domain"
)

type PaymentDocument struct {
	Id                 *int64     `bson:"_id"`
	BranchPaymentId    *string    `bson:"branchPaymentId"`
	BranchId           *int64     `bson:"branchId"`
	ProductId          *int64     `bson:"productId"`
	ProductType        *string    `bson:"productType"`
	PaymentAmount      *float64   `bson:"paymentAmount"`
	PaymentDate        *time.Time `bson:"paymentDate"`
	CurrencyIsoCode    *int64     `bson:"currencyIsoCode"`
	Status             *string    `bson:"status"`
	PrimaryAccountName *string    `bson:"primaryAccountName"`
	AcquirerId         *int64     `bson:"acquirerId"`
	BranchName         *string    `bson:"branchName"`
	BankAccount        *string    `bson:"bankAccount"`
	BankAgencyCode     *string    `bson:"bankAgencyCode"`
	BankCode           *string    `bson:"bankCode"`
	BankIspb           *int64     `bson:"bankIspb"`
	BankAccountType    *string    `bson:"bankAccountType"`
	BranchDocumentType *string    `bson:"branchDocumentType"`
	BranchDocument     *string    `bson:"branchDocument"`
	DocumentType       *string    `bson:"documentType"`
	Document           *string    `bson:"document"`
	PaymentType        *string    `bson:"paymentType"`
	MovementCounter    *int64     `bson:"movementCounter"`
	Prepayment         *string    `bson:"prepayment"`
	PaymentOrigin      *string    `bson:"paymentOrigin"`
	PaymentAccountId   *string    `bson:"paymentAccountId"`
	EventType          *string    `bson:"eventType"`
}

func FromDomain(payment *payment_domain.Payment) *PaymentDocument {
	return &PaymentDocument{
		Id:                 payment.Id,
		BranchPaymentId:    payment.BranchPaymentId,
		BranchId:           payment.BranchId,
		ProductId:          payment.ProductId,
		ProductType:        payment.ProductType,
		PaymentAmount:      payment.PaymentAmount,
		PaymentDate:        payment.PaymentDateValue(),
		CurrencyIsoCode:    payment.CurrencyIsoCode,
		Status:             payment.Status,
		PrimaryAccountName: payment.PrimaryAccountName,
		AcquirerId:         payment.AcquirerId,
		BranchName:         payment.BranchName,
		BankAccount:        payment.BankAccount,
		BankAgencyCode:     payment.BankAgencyCode,
		BankCode:           payment.BankCode,
		BankIspb:           payment.BankIspb,
		BankAccountType:    payment.BankAccountType,
		BranchDocumentType: payment.BranchDocumentType,
		BranchDocument:     payment.BranchDocument,
		DocumentType:       payment.DocumentType,
		Document:           payment.Document,
		PaymentType:        payment.PaymentType,
		MovementCounter:    payment.MovementCounter,
		Prepayment:         payment.Prepayment,
		PaymentOrigin:      payment.PaymentOrigin,
		EventType:          payment.EventType,
	}
}
