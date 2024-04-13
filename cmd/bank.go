package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"ysun.co/churn/internal/db"
	"ysun.co/churn/internal/lib"
	"ysun.co/churn/internal/ui/form"
	"ysun.co/churn/schema"
)

// churn bank --
var bankCmd = &cobra.Command{
	Use:   "bank",
	Short: "Manage banks (add, delete, edit)",
	Long:  "Open a interactive TUI to manage banks, add, delete, edit, etc.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// churn bank add --
var bankAddCmd = &cobra.Command{
	Use:   "add <bank alias>",
	Short: "Add a bank",
	Long:  "Add a bank by its alias, this will create a new bank with the given alias.",
	Args:  cobra.ExactArgs(1), // bank alias
	Run: func(cmd *cobra.Command, args []string) {
		bank := &schema.Bank{
			BankAlias: args[0],
		}

		err := form.FormBankAdd(bank)
		if err != nil {
			panic(err)
		}

		db, _ := db.Connect()
		bank.Add(db)
	},
}

// churn bank delete --
var bankDeleteCmd = &cobra.Command{
	Use:   "delete <bank alias>",
	Short: "Delete a bank", // bank alias
	Long:  "Delete a bank by its alias, note that this will also delete all products, accounts, transactions, etc. associated with the bank.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bank := &schema.Bank{
			BankAlias: args[0],
		}

		db, _ := db.Connect()
		bank.Delete(db)
	},
}

// churn bank import --
var bankImportCmd = &cobra.Command{
	Use:   "import <uri>",
	Short: "Import bank(s)", // either a file or a http(s) uri
	Long:  "Import bank(s) from a local JSON file or a http(s) uri that returns a JSON file.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		uri := args[0]
		banks := make([]*schema.Bank, 0)

		err := lib.Import(uri, &banks)
		if err != nil {
			panic(err)
		}

		for _, bank := range banks {
			db, _ := db.Connect()
			bank.Add(db)
		}

		fmt.Printf("Imported %d bank(s)\n", len(banks))
	},
}

func init() {
	rootCmd.AddCommand(bankCmd)
	bankCmd.AddCommand(bankAddCmd)
	bankCmd.AddCommand(bankDeleteCmd)
	bankCmd.AddCommand(bankImportCmd)
}
