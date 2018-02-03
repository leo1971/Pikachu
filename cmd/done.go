package cmd

import (
	"github.com/spf13/cobra"
)

var index int
var allFlag bool
var undoneFlag bool
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "done a todo item.",
	Long: `pika
done`,
	Run: func(cmd *cobra.Command, args []string) {
		println("pikachu, done.")
	},
}

func init() {
	doneCmd.PersistentFlags().StringVarP(&itemName, "name", "n", "", "figure the thing you have done.")
	doneCmd.PersistentFlags().BoolVarP(&undoneFlag, "undone", "u", false, "undone a thing you had done.")
	doneCmd.PersistentFlags().IntVarP(&index, "index", "i", 99, "done a thing by index.")
	doneCmd.PersistentFlags().BoolVarP(&allFlag, "all", "a", false, "done all your thing in the list.")
}
