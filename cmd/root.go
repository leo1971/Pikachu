package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pikachu",
	Short: "pikachu is a command-line tool for managing your todo list.",
	Long: `pika
pika`,
	Run: func(cmd *cobra.Command, args []string) {
		println("Hi, pikachu")
	},
}

func init() {
	rootCmd.AddCommand(todoCmd)

	todoCmd.AddCommand(todoAddCmd)
	todoCmd.AddCommand(todoDoneCmd)
	todoCmd.AddCommand(todoInitCmd)
}

// Execute will be the only export interface for main.go
func Execute() {
	rootCmd.Execute()
}
