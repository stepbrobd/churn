package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"ysun.co/churn/internal/db"
	"ysun.co/churn/internal/lib"
	"ysun.co/churn/internal/ui"
	"ysun.co/churn/internal/ui/form"
	"ysun.co/churn/schema"
)

// churn account --
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage accounts (add, delete, edit)",
	Long:  "Open an interactive TUI to manage account, add, delete, edit, etc.",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		db := db.Query()

		accounts := make([]*schema.Account, 0)
		err = db.SelectFrom("account").Do(&accounts)
		if err != nil {
			fmt.Println("Failed to fetch accounts")
			os.Exit(1)
		}

		products := make([]*schema.Product, 0)
		err = db.SelectFrom("product").Do(&products)
		if err != nil {
			fmt.Println("Failed to fetch products")
			os.Exit(1)
		}

		columns := []table.Column{
			{Title: "ID", Width: 5},
			{Title: "Product", Width: 25},
			{Title: "Limit", Width: 15},
			{Title: "Opened", Width: 15},
			{Title: "Closed", Width: 15},
		}

		rows := make([]table.Row, 0)
		for _, account := range accounts {
			product := schema.Product{}
			for _, p := range products {
				if p.ID == account.ProductID {
					product = *p
					break
				}
			}

			var opened string
			if account.Opened.Valid {
				opened = account.Opened.Time.Format("2006-01-02")
			} else {
				opened = "N/A"
			}

			var closed string
			if account.Closed.Valid {
				closed = account.Closed.Time.Format("2006-01-02")
			} else {
				closed = "N/A"
			}

			rows = append(rows, table.Row{
				strconv.Itoa(account.ID),
				product.ProductName,
				fmt.Sprintf("$%.2f", account.CL),
				opened,
				closed,
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

// churn account add --
var accountAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an account",
	Long:  "Add an account interactively, this will create a new account, and link it to the product.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		account := &schema.Account{}

		err := form.FormAccountAdd(account)
		if err != nil {
			fmt.Println("Failed to add account")
			os.Exit(1)
		}

		db, _ := db.Connect()
		account.Add(db)
	},
}

// churn account delete --
var forceAccountDeletion bool
var accountDeleteCmd = &cobra.Command{
	Use:   "delete <account id>",
	Short: "Delete an account", // by account id
	Long:  "Delete an account by its id, this will delete the account and all its associated data.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !forceAccountDeletion {
			fmt.Println("This will delete the account and all its associated data.")
			confirm := lib.Confirm()
			if !confirm {
				return
			}
		}

		id, _ := strconv.Atoi(args[0])
		account := &schema.Account{
			ID: id,
		}

		db, _ := db.Connect()
		account.Delete(db)
	},
}

func init() {
	accountDeleteCmd.Flags().BoolVarP(&forceAccountDeletion, "force", "f", false, "Force delete the account and all its associated data")

	rootCmd.AddCommand(accountCmd)
	accountCmd.AddCommand(accountAddCmd)
	accountCmd.AddCommand(accountDeleteCmd)
}
