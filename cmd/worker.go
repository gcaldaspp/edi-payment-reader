package cmd

import (
	"encoding/json"
	"log/slog"

	"github.com/PicPay/ms-edi-wrk-payment-reader-go/config"
	"github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/kafka"
	"github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/payment"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var WorkerCmd = &cobra.Command{
	Use:   "wrk",
	Short: "Start worker kafka",
	Long:  "Start worker kafka that listen an specific topic",
	Run: func(cmd *cobra.Command, args []string) {
		var topic = viper.GetString("KAFKA_TOPIC")

		db, err := config.InitMongoDB()
		if err != nil {
			slog.Error("There are an error on connect database", "error", err)
			return
		}

		repository := payment.NewPaymentRepository(db)
		service := payment.NewPaymentService(repository)
		consumer := kafka.NewPaymentConsumer(service)

		config.StartKafkaConsumer(topic, func(evt []byte) {
			var event kafka.PaymentEvent
			err := json.Unmarshal(evt, &event)
			if err == nil {
				consumer.Process(event)
				slog.Info("Received event with success", "branchPaymentId", event.Payload.BranchPaymentId)
			}
		})
	},
}
