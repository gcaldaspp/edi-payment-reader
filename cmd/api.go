package cmd

import (
	"net/http"

	"github.com/PicPay/ms-edi-wrk-payment-reader-go/internal/api"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var ApiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start api server",
	Long:  "Start the api server to test in QA",
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()

		r.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		})

		paymentController := api.NewPaymentController()
		paymentroutes := r.Group("/payment")
		{
			paymentroutes.POST("", paymentController.PaymentHandler)
		}

		r.Run()
	},
}
