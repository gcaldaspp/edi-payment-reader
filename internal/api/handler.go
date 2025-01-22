package api

import (
	"net/http"

	"github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/kafka"
	"github.com/gin-gonic/gin"
)

type PaymentController struct{}

func NewPaymentController() *PaymentController {
	return &PaymentController{}
}

func (pc *PaymentController) PaymentHandler(ctx *gin.Context) {
	var event kafka.PaymentEvent
	ctx.ShouldBindBodyWithJSON(&event)
	ctx.JSON(http.StatusOK, event)
}
