package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/steschwa/hopper-analytics-collector/cmd"

	"github.com/getsentry/sentry-go"
)

func main() {
	initSentry()

	rootCmd := &cobra.Command{}

	cmd.RegisterHoppersCmd(rootCmd)
	cmd.RegisterMarketsCmd(rootCmd)
	cmd.RegisterPricesCmd(rootCmd)
	cmd.RegisterFlySupplyCmd(rootCmd)
	cmd.RegisterVotesCmd(rootCmd)
	cmd.RegisterBaseSharesCmd(rootCmd)
	cmd.RegisterHopperHoldersCmd(rootCmd)

	rootCmd.Execute()
}

func initSentry() {
	env := os.Getenv("ENV")

	if env == "production" {
		err := sentry.Init(sentry.ClientOptions{
			Dsn:         "https://0997d0f7af464bd29da229b2c9f39c05@o1202748.ingest.sentry.io/6328151",
			Environment: "production",
		})
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Using Sentry for error reporting")
	}
}
