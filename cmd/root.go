package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pika",
	Short: "pika is a command-line tool for managing your todo list.",
	Long: `pika
          pika`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
