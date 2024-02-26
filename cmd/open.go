package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open <name> on <date> with <limit>",
	Short: "Add a new account to the database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("open called")
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
