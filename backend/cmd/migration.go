package cmd

import (
	"github.com/spf13/cobra"
)

var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "Manage migration",
}

func init() {
	rootCmd.AddCommand(migrationCmd)
}
