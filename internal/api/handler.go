package api

import (
	"net/http"

	"github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/kafka"
	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	paymentConsumer *kafka.PaymentConsumer
}

func NewPaymentController(consumer *kafka.PaymentConsumer) *PaymentController {
	return &PaymentController{
		paymentConsumer: consumer,
	}
}

func (pc *PaymentController) PaymentHandler(ctx *gin.Context) {
	var event kafka.PaymentEvent
	ctx.ShouldBindBodyWithJSON(&event)
	pc.paymentConsumer.Process(event)
	ctx.JSON(http.StatusOK, event)
}
