package cmd

import (
	"log/slog"
	"net/http"

	"github.com/PicPay/ms-edi-wrk-payment-reader-go/config"
	"github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/api"
	"github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/kafka"
	"github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/payment"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var ApiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start api server",
	Long:  "Start the api server to test in QA",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := config.InitMongoDB()
		if err != nil {
			slog.Error("There are an error on connect database", "error", err)
			return
		}

		repository := payment.NewPaymentRepository(db)
		service := payment.NewPaymentService(repository)
		consumer := kafka.NewPaymentConsumer(service)
		paymentController := api.NewPaymentController(consumer)

		r := gin.Default()

		r.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		})

		r.POST("/payment", paymentController.PaymentHandler)

		r.Run()
	},
}
