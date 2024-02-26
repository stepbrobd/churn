package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var closeCmd = &cobra.Command{
	Use:   "close <name> on <date>",
	Short: "Mark an account as closed",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("close called")
	},
}

func init() {
	rootCmd.AddCommand(closeCmd)
}
