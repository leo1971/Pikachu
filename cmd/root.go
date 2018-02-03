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
	rootCmd.AddCommand(bornCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(doneCmd)
	rootCmd.AddCommand(deleteCmd)
}

// Execute will be the only export interface for main.go
func Execute() {
	rootCmd.Execute()
}
