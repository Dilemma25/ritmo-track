package cmd

import (
	"ritmotrack-backend/internal/boostrap"
	"ritmotrack-backend/internal/presentation/cli"

	"github.com/spf13/cobra"
)

var migrationDownCmd = &cobra.Command{
	Use:   "down",
	Short: "down migration",
	Run: func(cmd *cobra.Command, args []string) {
		boostrap.GetContainer().Invoke(func(controller cli.MigrationCommandController) {
			controller.Down()
		})
	},
}

func init() {
	migrationCmd.AddCommand(migrationDownCmd)
}
