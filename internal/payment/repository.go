package payment

import (
	"context"

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

func (pri *PaymentRepositoryImpl) Create(payment *Payment) error {
	d := FromDomain(payment)
	_, err := pri.db.Collection("acquirer_payments").InsertOne(context.TODO(), d)
	return err
}
