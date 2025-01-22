package main

import (
	"github.com/PicPay/ms-edi-wrk-payment-reader-go/cmd"
	"github.com/PicPay/ms-edi-wrk-payment-reader-go/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wrk-payment-reader",
	Short: "Service that is responsible to lister payments",
	Long:  "The wrk-payment-reader is the service responsible to listener payment events comming from GLIC",
}

func main() {
	config.InitViper()
	rootCmd.AddCommand(cmd.ApiCmd)
	rootCmd.AddCommand(cmd.WorkerCmd)
	rootCmd.Execute()
}
