package cmd

import (
	"ritmotrack-backend/internal/boostrap"
	"ritmotrack-backend/internal/presentation/cli"

	"github.com/spf13/cobra"
)

var migrationUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Apply migration",
	Run: func(cmd *cobra.Command, args []string) {
		boostrap.GetContainer().Invoke(func(controller cli.MigrationCommandController) {
			controller.Upgrade()
		})
	},
}

func init() {
	migrationCmd.AddCommand(migrationUpCmd)
}
