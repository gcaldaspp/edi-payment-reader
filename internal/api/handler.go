package api

import (
	"log/slog"
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

	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		slog.Error("Invalid payload", "error", err)
		return
	}

	slog.Info("Received payload", "event", event)
	pc.paymentConsumer.Process(event)
	ctx.JSON(http.StatusOK, event)
}
