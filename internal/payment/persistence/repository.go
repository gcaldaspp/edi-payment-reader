package payment_persistence

import (
	"context"

	payment_domain "github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/payment/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentRepositoryImpl struct {
	db *mongo.Database
}

func NewPaymentRepository(db *mongo.Database) *PaymentRepositoryImpl {
	return &PaymentRepositoryImpl{
		db: db,
	}
}

func (pri *PaymentRepositoryImpl) Create(payment *payment_domain.Payment) error {
	d := FromDomain(payment)
	_, err := pri.db.Collection("acquirer_payments").InsertOne(context.TODO(), d)
	return err
}
