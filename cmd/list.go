package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var sortType string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "show todo list.",
	Long: `pika
list`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("sorttype:%v\n", sortType)
	},
}

func init() {
	listCmd.PersistentFlags().StringVarP(&sortType, "sort", "s", "time", "choose a way to show todo list.")
}
