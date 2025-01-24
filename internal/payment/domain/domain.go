package payment_domain

import (
	"time"

	"github.com/PicPay/ms-edi-wrk-payment-reader-go/pkg/types"
)

type Payment struct {
	Id                 *int64
	BranchPaymentId    *string
	BranchId           *int64
	ProductId          *int64
	ProductType        *string
	PaymentAmount      *float64
	PaymentDate        *types.DateOnly
	CurrencyIsoCode    *int64
	Status             *string
	PrimaryAccountName *string
	AcquirerId         *int64
	BranchName         *string
	BankAccount        *string
	BankAgencyCode     *string
	BankCode           *string
	BankIspb           *int64
	BankAccountType    *string
	BranchDocumentType *string
	BranchDocument     *string
	DocumentType       *string
	Document           *string
	PaymentType        *string
	MovementCounter    *int64
	Prepayment         *string
	PaymentOrigin      *string
	PaymentAccountId   *string
	EventType          *string
}

func (p *Payment) IsApproved() bool {
	if p.Status != nil {
		return *p.Status == "APPROVED"
	}
	return false
}

func (p *Payment) PaymentDateValue() *time.Time {
	if p.PaymentDate != nil {
		return &p.PaymentDate.Time
	}
	return nil
}

func (p *Payment) PaymentDateString() *string {
	if p.PaymentDate != nil {
		date := p.PaymentDate.ToString()
		return &date
	}
	return nil
}
