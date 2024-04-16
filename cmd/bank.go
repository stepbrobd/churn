package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"ysun.co/churn/internal/db"
	"ysun.co/churn/internal/lib"
	"ysun.co/churn/internal/ui"
	"ysun.co/churn/internal/ui/form"
	"ysun.co/churn/schema"
)

// churn bank --
var bankCmd = &cobra.Command{
	Use:   "bank",
	Short: "Manage banks (add, delete, edit)",
	Long:  "Open an interactive TUI to manage banks, add, delete, edit, etc.",
	Run: func(cmd *cobra.Command, args []string) {
		db := db.Query()
		banks := make([]*schema.Bank, 0)
		err := db.SelectFrom("bank").Do(&banks)
		if err != nil {
			fmt.Println("Failed to fetch banks")
			os.Exit(1)
		}

		columns := []table.Column{
			{Title: "Name", Width: 25},
			{Title: "Alias", Width: 15},
			// {Title: "Max Account", Width: 12},
			// {Title: "Max Account Period", Width: 18},
			// {Title: "Max Account Scope", Width: 17},
		}

		rows := make([]table.Row, 0)
		for _, bank := range banks {
			rows = append(rows, table.Row{
				bank.BankName,
				bank.BankAlias,
				// strconv.FormatInt(bank.MaxAccount.Int64, 10),
				// strconv.FormatInt(bank.MaxAccountPeriod.Int64, 10),
				// bank.MaxAccountScope.String,
			})
		}

		t := table.New(
			table.WithColumns(columns),
			table.WithRows(rows),
			table.WithFocused(true),
		)

		m := ui.ModelTable{Table: t}
		tea.NewProgram(m).Run()
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
			fmt.Println("Failed to add bank")
			os.Exit(1)
		}

		db, _ := db.Connect()
		bank.Add(db)
	},
}

// churn bank delete --
var forceBankDeletion bool
var bankDeleteCmd = &cobra.Command{
	Use:   "delete <bank alias>",
	Short: "Delete a bank", // bank alias
	Long:  "Delete a bank by its alias, note that this will also delete all products, accounts, transactions, etc. associated with the bank.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !forceBankDeletion {
			fmt.Println("This will delete the bank and all its associated data.")
			confirm := lib.Confirm()
			if !confirm {
				return
			}
		}

		bank := &schema.Bank{
			BankAlias: args[0],
		}

		db, _ := db.Connect()
		bank.Delete(db)

		fmt.Printf("Deleted bank %s\n", bank.BankAlias)
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
			fmt.Println("Failed to import bank(s)")
			os.Exit(1)
		}

		db, _ := db.Connect()
		for _, bank := range banks {
			bank.Add(db)
		}

		fmt.Printf("Imported %d bank(s)\n", len(banks))
	},
}

func init() {
	bankDeleteCmd.Flags().BoolVarP(&forceBankDeletion, "force", "f", false, "Force delete the bank and all its associated data")

	rootCmd.AddCommand(bankCmd)
	bankCmd.AddCommand(bankAddCmd)
	bankCmd.AddCommand(bankDeleteCmd)
	bankCmd.AddCommand(bankImportCmd)
}
