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

// churn tx --
var txCmd = &cobra.Command{
	Use:   "tx",
	Short: "Manage transactions (add, delete, edit)",
	Long:  "Open an interactive TUI to manage transactions, add, delete, edit, etc.",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		db := db.Query()

		txs := make([]*schema.Tx, 0)
		err = db.SelectFrom("tx").Do(&txs)
		if err != nil {
			fmt.Println("Failed to fetch transactions")
			os.Exit(1)
		}

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
			{Title: "Timestamp", Width: 15},
			{Title: "Account ID", Width: 10},
			{Title: "Product", Width: 25},
			{Title: "Amount", Width: 15},
			{Title: "Category", Width: 15},
			// {Title: "Note", Width: 25},
		}

		rows := make([]table.Row, 0)
		for _, tx := range txs {
			account := schema.Account{}
			for _, a := range accounts {
				if a.ID == tx.AccountID {
					account = *a
				}
			}

			product := schema.Product{}
			for _, p := range products {
				if p.ID == account.ProductID {
					product = *p
				}
			}

			rows = append(rows, table.Row{
				strconv.Itoa(tx.ID),
				tx.TxTimestamp.Format("2006-01-02"),
				strconv.Itoa(tx.AccountID),
				product.ProductName,
				fmt.Sprintf("$%.2f", tx.Amount),
				tx.Category,
				// tx.Note.String,
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

// churn tx add --
var txAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a transaction",
	Long:  "Add a transaction interactively, this will create a new transaction, and associate it with an account",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		tx := &schema.Tx{}

		err := form.FormTxAdd(tx)
		if err != nil {
			fmt.Println("Failed to add transaction")
			os.Exit(1)
		}

		db, _ := db.Connect()
		tx.Add(db)
	},
}

// churn tx edit --
var txEditCmd = &cobra.Command{
	Use:   "edit <tx id>",
	Short: "Edit a transaction",
	Long:  "Edit a transaction by its ID, this will update the transaction with the given ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		q := db.Query()
		txs := make([]*schema.Tx, 0)
		err := q.SelectFrom("tx").Do(&txs)
		if err != nil {
			fmt.Println("Failed to fetch transactions")
			os.Exit(1)
		}

		tx := &schema.Tx{}
		for _, t := range txs {
			if strconv.Itoa(t.ID) == args[0] {
				tx = t
				break
			}
		}

		err = form.FormTxAdd(tx)
		if err != nil {
			fmt.Println("Failed to edit transaction")
			os.Exit(1)
		}

		db, _ := db.Connect()
		tx.Update(db)
	},
}

// churn tx delete --
var forceTxDeletion bool
var txDeleteCmd = &cobra.Command{
	Use:   "delete <tx id>",
	Short: "Delete a transaction",
	Long:  "Delete a transaction by its ID, this will remove the bonus and all its associated data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !forceTxDeletion {
			fmt.Println("This will delete the transaction and all its associated data.")
			if !lib.Confirm() {
				return
			}
		}

		id, _ := strconv.Atoi(args[0])
		tx := &schema.Tx{ID: id}

		db, _ := db.Connect()
		tx.Delete(db)
	},
}

func init() {
	txDeleteCmd.Flags().BoolVarP(&forceTxDeletion, "force", "f", false, "Force delete the transaction and all its associated data")

	rootCmd.AddCommand(txCmd)
	txCmd.AddCommand(txAddCmd)
	txCmd.AddCommand(txEditCmd)
	txCmd.AddCommand(txDeleteCmd)
}
