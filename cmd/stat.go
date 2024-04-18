package cmd

import (
	"fmt"
	"os"

	"github.com/samonzeweb/godb"
	"github.com/spf13/cobra"
	"ysun.co/churn/internal/db"
	"ysun.co/churn/schema"
)

var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "Show statistics",
	Long:  "Show statistics of account, bank, ... related usage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var statAccountCmd = &cobra.Command{
	Use:   "account",
	Short: "Show statistics of accounts",
	Long:  "Show statistics of account related usage.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		q := db.Query()
		accounts := make([]*schema.Account, 0)
		err := q.SelectFrom("account").Do(&accounts)
		if err != nil {
			fmt.Println("Failed to fetch accounts")
			os.Exit(1)
		}

		var openAccounts, closedAccounts int
		var totalLimit float64
		for _, a := range accounts {
			if a.Closed.Valid {
				closedAccounts++
			} else {
				openAccounts++
			}

			totalLimit += a.CL
		}

		fmt.Println("Total accounts:", len(accounts))
		fmt.Println("Open accounts:", openAccounts)
		fmt.Println("Closed accounts:", closedAccounts)
		fmt.Println("Total limit:", totalLimit)
	},
}

var statRewardCmd = &cobra.Command{
	Use:   "reward",
	Short: "Show statistics of rewards",
	Long:  "Show statistics of reward related usage.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		q := db.Query()

		// total number of rewards
		numRewards, err := q.SelectFrom("reward").Count()
		if err != nil {
			fmt.Println("Failed to fetch rewards")
			os.Exit(1)
		}
		fmt.Println("Total rewards:", numRewards)

		// transaction tied to rewards
		numTx, err := q.SelectFrom("tx").
			LeftJoin("account", "account", godb.Q("tx.account_id = account.id")).
			LeftJoin("product", "product", godb.Q("account.product_id = product.id")).
			LeftJoin("reward", "reward", godb.Q("product.id = reward.product_id")).
			Count()
		if err != nil {
			fmt.Println("Failed to fetch transactions")
			os.Exit(1)
		}
		fmt.Println("Total transactions tied to rewards:", numTx)

		// total unique reward categories
		uniqueCategories := make([]*schema.Reward, 0)
		err = q.SelectFrom("reward").Columns("category").Distinct().Do(&uniqueCategories)
		if err != nil {
			fmt.Println("Failed to fetch unique reward categories")
			os.Exit(1)
		}
		fmt.Println("Total unique reward categories:", len(uniqueCategories))

		// total reward amount per account
		// accounts := make([]*schema.Account, 0)
		// err = q.SelectFrom("account").Do(&accounts)
		// if err != nil {
		// 	fmt.Println("Failed to fetch accounts")
		// 	os.Exit(1)
		// }
		// for _, a := range accounts {
		// 	totalReward := make([]*schema.Tx, 0)
		// 	err := q.SelectFrom("tx").
		// 		LeftJoin("account", "account", godb.Q("tx.account_id = account.id")).
		// 		LeftJoin("product", "product", godb.Q("account.product_id = product.id")).
		// 		LeftJoin("reward", "reward", godb.Q("product.id = reward.product_id")).
		// 		Where("account.id = ?", a.ID).
		// 		Do(&totalReward)
		// 	if err != nil {
		// 		fmt.Println("Failed to fetch total reward for account", a.ID)
		// 		os.Exit(1)
		// 	}
		// 	fmt.Println("Total reward for account", a.ID, ":", totalReward)
		// }
	},
}

var statBonusCmd = &cobra.Command{
	Use:   "bonus",
	Short: "Show statistics of bonuses",
	Long:  "Show statistics of bonus related usage.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		q := db.Query()

		// total number of bonuses
		numBonuses, err := q.SelectFrom("bonus").Count()
		if err != nil {
			fmt.Println("Failed to fetch bonuses")
			os.Exit(1)
		}
		fmt.Println("Total bonuses:", numBonuses)

		// accounts with bonuses
		accounts := make([]*schema.Account, 0)
		err = q.SelectFrom("account").Do(&accounts)
		if err != nil {
			fmt.Println("Failed to fetch accounts")
			os.Exit(1)
		}
		for _, a := range accounts {
			bonuses := make([]*schema.Bonus, 0)
			err := q.SelectFrom("bonus").Where("account_id = ?", a.ID).Do(&bonuses)
			if err != nil {
				fmt.Println("Failed to fetch bonuses for account", a.ID)
				os.Exit(1)
			}
			fmt.Println("Bonuses for account", a.ID)
			for _, b := range bonuses {
				// type + bonus amount + unit + spending requirement
				fmt.Printf("  - %s: %f %s, spending requirement: %f\n", b.BonusType, b.BonusAmount, b.Unit, b.Spend)
			}
		}
	},
}

var statTxCmd = &cobra.Command{
	Use:   "tx",
	Short: "Show statistics of transactions",
	Long:  "Show statistics of transaction related usage.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		q := db.Query()
		txs := make([]*schema.Tx, 0)
		err := q.SelectFrom("tx").Do(&txs)
		if err != nil {
			fmt.Println("Failed to fetch transactions")
			os.Exit(1)
		}

		var totalAmount float64
		for _, t := range txs {
			totalAmount += t.Amount
		}

		fmt.Println("Total transactions:", len(txs))
		fmt.Println("Total amount:", totalAmount)
	},
}

func init() {
	rootCmd.AddCommand(statCmd)
	statCmd.AddCommand(statAccountCmd)
	statCmd.AddCommand(statRewardCmd)
	statCmd.AddCommand(statBonusCmd)
	statCmd.AddCommand(statTxCmd)
}
