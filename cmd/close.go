package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// closeCmd represents the close command
var closeCmd = &cobra.Command{
	Use:   "close <name> on <date>",
	Short: "Mark an account as closed",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("close called")
	},
}

func init() {
	rootCmd.AddCommand(closeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// closeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// closeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
