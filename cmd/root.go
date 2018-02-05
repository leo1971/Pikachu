package cmd

import (
	"github.com/leo1971/pikachu/pikachu"
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
var list pikachu.TodoList

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
