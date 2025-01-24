package cmd

import (
	"log/slog"
	"net/http"

	"github.com/PicPay/ms-edi-wrk-payment-reader-go/config"
	"github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/api"
	kafka_consumer "github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/kafka/consumer"
	kafka_producer "github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/kafka/producer"
	payment_persistence "github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/payment/persistence"
	payment_service "github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/payment/service"
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

		repository := payment_persistence.NewPaymentRepository(db)
		producer := kafka_producer.NewPaymentProducer()
		service := payment_service.NewPaymentService(repository, producer)
		consumer := kafka_consumer.NewPaymentConsumer(service)
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
