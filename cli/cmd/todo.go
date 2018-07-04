package cmd

import (
	"errors"

	"github.com/leo1971/pikachu/cli/pikachu"
	"github.com/spf13/cobra"
)

var thing string
var doneall bool
var doneid string

var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "Todo list.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return pikachu.TodoShow()
	},
}

var todoAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add todo item.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if thing == "" {
			return errors.New("Thing is required")
		}
		pikachu.TodoAdd(thing)
		return pikachu.TodoShow()
	},
}

var todoDoneCmd = &cobra.Command{
	Use:   "done",
	Short: "done a todo item.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if doneall == true {
			pikachu.TodoDoneAll()
		} else {
			if doneid == "" {
				return errors.New("ID is required")
			}
			pikachu.TodoDone(doneid)
		}
		return pikachu.TodoShow()
	},
}

var todoInitCmd = &cobra.Command{
	Use:   "init",
	Short: "init todo cache file.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return pikachu.TodoInit()
	},
}

func init() {
	todoCmd.PersistentFlags().StringVarP(&thing,
		"thing", "t", "", "todo item.")
	todoDoneCmd.PersistentFlags().BoolVarP(&doneall, "all", "a", false, "done all things.")
	todoDoneCmd.PersistentFlags().StringVarP(&doneid, "id", "i", "", "done something.")
}
