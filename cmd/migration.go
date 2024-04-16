package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"ysun.co/churn/internal/migration"
)

// churn migration --
var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "Manage database migrations",
	Long:  "List and execute database migrations.",
	Args:  cobra.ExactArgs(0), // product alias
}

// churn migration list --
var migrationListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available migrations",
	Long:  "List all available migrations.",
	Run: func(cmd *cobra.Command, args []string) {
		names, err := migration.List()
		if err != nil {
			panic(err)
		}

		for i, name := range names {
			fmt.Printf("%d. %s\n", i+1, name)
		}
	},
}

// churn migration exec --
var migrationExecCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute specified migration",
	Long:  "Execute specified migration.",
	Args:  cobra.ExactArgs(1), // migration name
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		err := migration.Exec(name)
		if err != nil {
			fmt.Printf("Error executing migration '%s': %s\n", name, err)
		}
		fmt.Printf("Migration '%s' executed successfully\n", name)
	},
}

func init() {
	rootCmd.AddCommand(migrationCmd)
	migrationCmd.AddCommand(migrationListCmd)
	migrationCmd.AddCommand(migrationExecCmd)
}
