package cmd

import (
	"github.com/spf13/cobra"
)

var bornCmd = &cobra.Command{
	Use:   "born",
	Short: "init pikachu.",
	Long: `pika
born`,
	Run: func(cmd *cobra.Command, args []string) {
		println("pikachu, born.")
	},
}
