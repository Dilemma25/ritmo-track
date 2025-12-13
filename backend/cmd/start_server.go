package cmd

import (
	"ritmotrack-backend/internal/boostrap"

	"github.com/spf13/cobra"
)

var startServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start server",
	Run: func(cmd *cobra.Command, args []string) {
		StartServer()
	},
}

func init() {
	startCmd.AddCommand(startServerCmd)
}

func StartServer() {
	boostrap.Start()
}
