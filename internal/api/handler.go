package api

import (
	"log/slog"
	"net/http"

	kafka_consumer "github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/kafka/consumer"
	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	paymentConsumer *kafka_consumer.PaymentConsumer
}

func NewPaymentController(consumer *kafka_consumer.PaymentConsumer) *PaymentController {
	return &PaymentController{
		paymentConsumer: consumer,
	}
}

func (pc *PaymentController) PaymentHandler(ctx *gin.Context) {
	var event kafka_consumer.PaymentEvent

	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		slog.Error("Invalid payload", "error", err)
		return
	}

	slog.Info("Received payload", "branchPaymentId", *event.Payload.BranchPaymentId)
	pc.paymentConsumer.Process(event)
	ctx.JSON(http.StatusOK, event)
}
