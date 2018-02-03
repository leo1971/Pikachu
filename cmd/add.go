package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var itemName string
var itemLevel string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add todo item.",
	Long: `pika
add`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("name:%v\nlevel:%v\n", itemName, itemLevel)
	},
}

func init() {
	addCmd.PersistentFlags().StringVarP(&itemName,
		"name", "n", "", "the thing you want to add to the list.")
	addCmd.PersistentFlags().StringVarP(&itemLevel,
		"level", "l", "easy", "figure the level of the thing that you want to add to the list.")
}
