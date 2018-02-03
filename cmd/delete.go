package cmd

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a todo item.",
	Long: `pika
delete`,
	Run: func(cmd *cobra.Command, args []string) {
		println("pikachu, delete.")
	},
}

func init() {
	deleteCmd.PersistentFlags().IntVarP(&index, "index", "i", 0, "figure the index of the one you wanna delete.")
	deleteCmd.PersistentFlags().BoolVarP(&allFlag, "all", "a", false, "delete all item")
}
