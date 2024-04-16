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

// churn reward --
var rewardCmd = &cobra.Command{
	Use:   "reward",
	Short: "Manage rewards (add, delete, edit)",
	Long:  "Open an interactive TUI to manage rewards, add, delete, edit, etc.",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		db := db.Query()

		rewards := make([]*schema.Reward, 0)
		err = db.SelectFrom("reward").Do(&rewards)
		if err != nil {
			fmt.Println("Failed to fetch rewards")
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
			{Title: "Category", Width: 15},
			{Title: "Reward", Width: 15},
			{Title: "Unit", Width: 10},
		}

		rows := make([]table.Row, 0)
		for _, reward := range rewards {
			product := schema.Product{}
			for _, p := range products {
				if p.ID == reward.ProductID {
					product = *p
				}
			}

			rows = append(rows, table.Row{
				strconv.Itoa(reward.ID),
				product.ProductName,
				reward.Category,
				fmt.Sprintf("%.2f", reward.Reward),
				reward.Unit,
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

// churn reward add --
var rewardAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a reward",
	Long:  "Add a reward interactively, this will create a new reward, and associate it with a product",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		reward := &schema.Reward{}

		err := form.FormRewardAdd(reward)
		if err != nil {
			fmt.Println("Failed to add reward")
			os.Exit(1)
		}

		db, _ := db.Connect()
		reward.Add(db)
	},
}

// churn reward delete --
var forceRewardDeletion bool
var rewardDeleteCmd = &cobra.Command{
	Use:   "delete <reward id>",
	Short: "Delete a reward",
	Long:  "Delete a reward by its ID, this will remove the reward and all its associated data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !forceRewardDeletion {
			fmt.Println("This will delete the reward and all its associated data.")
			if !lib.Confirm() {
				return
			}
		}

		id, _ := strconv.Atoi(args[0])
		reward := &schema.Reward{
			ID: id,
		}

		db, _ := db.Connect()
		reward.Delete(db)
	},
}

func init() {
	rewardDeleteCmd.Flags().BoolVarP(&forceRewardDeletion, "force", "f", false, "Force delete the reward and all its associated data")

	rootCmd.AddCommand(rewardCmd)
	rewardCmd.AddCommand(rewardAddCmd)
	rewardCmd.AddCommand(rewardDeleteCmd)
}
