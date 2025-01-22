package cmd

import (
	"encoding/json"
	"log/slog"

	"github.com/PicPay/ms-edi-wrk-payment-reader-go/config"
	"github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/kafka"
	"github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/payment"
	"github.com/spf13/cobra"
)

var WorkerCmd = &cobra.Command{
	Use:   "wrk",
	Short: "Start worker kafka",
	Long:  "Start worker kafka that listen an specific topic",
	Run: func(cmd *cobra.Command, args []string) {
		var topic, _ = cmd.Flags().GetString("topic")

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
			json.Unmarshal(evt, &event)
			consumer.Process(event)
			slog.Info("Received event with success", "event", evt)
		})
	},
}

func init() {
	WorkerCmd.Flags().StringP("topic", "t", "", "Topic name")
	WorkerCmd.MarkFlagRequired("topic")
}
