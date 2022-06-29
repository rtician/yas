package cmd

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"yas/server"
)

// rootCmd represents the base command when called without any subcommands
var serviceCmd = &cobra.Command{
	Use:   "yas",
	Short: "CLI to run YAS app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("calling service")
		serviceStart()
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)
}

func serviceStart() {
	if err := server.Serve(); err != nil {
		log.Fatal().Err(err).Msg("error starting server")
	}
}
