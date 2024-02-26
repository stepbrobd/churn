package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stepbrobd/churn/internal/sqlite"
)

var openCmd = &cobra.Command{
	Use:   "open <name> on <date> with <limit>",
	Short: "Add a new account to the database",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := sqlite.Open()
		if err != nil {
			fmt.Println(err)
		}

		conn.Ping()
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
