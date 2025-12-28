package cmd

import (
	"ritmotrack-backend/internal/boostrap"
	"ritmotrack-backend/internal/presentation/cli"

	"github.com/spf13/cobra"
)

var migrationResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "reset all migrations",
	Run: func(cmd *cobra.Command, args []string) {
		boostrap.GetContainer().Invoke(func(controller cli.MigrationCommandController) {
			controller.Reset()
		})
	},
}

func init() {
	migrationCmd.AddCommand(migrationResetCmd)
}
