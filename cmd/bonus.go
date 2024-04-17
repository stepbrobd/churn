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

// churn bonus --
var bonusCmd = &cobra.Command{
	Use:   "bonus",
	Short: "Manage bonuses (add, delete, edit)",
	Long:  "Open an interactive TUI to manage bonuses, add, delete, edit, etc.",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		db := db.Query()

		bonuses := make([]*schema.Bonus, 0)
		err = db.SelectFrom("bonus").Do(&bonuses)
		if err != nil {
			fmt.Println("Failed to fetch bonuses")
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
			{Title: "Account", Width: 25},
			{Title: "Type", Width: 15},
			{Title: "Spend", Width: 15},
			{Title: "Amount", Width: 15},
			{Title: "Unit", Width: 10},
			{Title: "Start", Width: 15},
			{Title: "End", Width: 15},
		}

		rows := make([]table.Row, 0)
		for _, bonus := range bonuses {
			account := schema.Account{}
			for _, a := range accounts {
				if a.ID == bonus.AccountID {
					account = *a
					break
				}
			}

			product := schema.Product{}
			for _, p := range products {
				if p.ID == account.ProductID {
					product = *p
					break
				}
			}

			rows = append(rows, table.Row{
				strconv.Itoa(bonus.ID),
				product.ProductName,
				bonus.BonusType,
				fmt.Sprintf("$%.2f", bonus.Spend),
				fmt.Sprintf("%.2f", bonus.BonusAmount),
				bonus.Unit,
				bonus.BonusStart.Format("2006-01-02"),
				bonus.BonusEnd.Format("2006-01-02"),
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

// churn bonus add --
var bonusAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a bonus",
	Long:  "Add a bonus interactively, this will create a new bonus, and associate it with an account",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		bonus := &schema.Bonus{}

		err := form.FormBonusAdd(bonus)
		if err != nil {
			fmt.Println("Failed to add bonus")
			os.Exit(1)
		}

		db, _ := db.Connect()
		bonus.Add(db)
	},
}

// churn bonus edit --
var bonusEditCmd = &cobra.Command{
	Use:   "edit <bonus id>",
	Short: "Edit a bonus",
	Long:  "Edit a bonus by its ID, this will update the bonus and its associated data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		q := db.Query()
		bonuses := make([]*schema.Bonus, 0)
		err := q.SelectFrom("bonus").Do(&bonuses)
		if err != nil {
			fmt.Println("Failed to fetch bonuses")
			os.Exit(1)
		}

		bonus := &schema.Bonus{}
		for _, b := range bonuses {
			if strconv.Itoa(b.ID) == args[0] {
				bonus = b
				break
			}
		}

		err = form.FormBonusAdd(bonus)
		if err != nil {
			fmt.Println("Failed to edit bonus")
			os.Exit(1)
		}

		db, _ := db.Connect()
		bonus.Update(db)
	},
}

// churn bonus delete --
var forceBonusDeletion bool
var bonusDeleteCmd = &cobra.Command{
	Use:   "delete <bonus id>",
	Short: "Delete a bonus",
	Long:  "Delete a bonus by its ID, this will remove the bonus and all its associated data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !forceBonusDeletion {
			fmt.Println("This will delete the bonus and all its associated data.")
			if !lib.Confirm() {
				return
			}
		}

		id, _ := strconv.Atoi(args[0])
		bonus := &schema.Bonus{
			ID: id,
		}

		db, _ := db.Connect()
		bonus.Delete(db)
	},
}

func init() {
	bonusDeleteCmd.Flags().BoolVarP(&forceBonusDeletion, "force", "f", false, "Force delete the bonus and all its associated data")

	rootCmd.AddCommand(bonusCmd)
	bonusCmd.AddCommand(bonusAddCmd)
	bonusCmd.AddCommand(bonusEditCmd)
	bonusCmd.AddCommand(bonusDeleteCmd)
}
