package cmd

import (
	"github.com/spf13/cobra"
)

// churn product --
var productCmd = &cobra.Command{
	Use:   "product",
	Short: "Manage product (add, delete, edit)",
	Long:  "Open a interactive TUI to manage product, add, delete, edit, etc.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// churn product add --
var productAddCmd = &cobra.Command{
	Use:   "add <product alias> by <bank alias>",
	Short: "Add a product",
	Long:  "Add a product by its alias, this will create a new product with the given alias.",
	Args:  cobra.ExactArgs(3), // product alias, by, bank alias
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// churn product delete --
var productDeleteCmd = &cobra.Command{
	Use:   "delete <product alias>",
	Short: "Delete a product", // product alias
	Long:  "Delete a product by its alias, note that this will also delete all accounts, rewards, transactions, etc. associated with the product.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// churn product import --
var productImportCmd = &cobra.Command{
	Use:   "import <uri>",
	Short: "Import product(s)", // either a file or a http(s) uri
	Long:  "Import product(s) from a local JSON file or a http(s) uri that returns a JSON file.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(productCmd)
	productCmd.AddCommand(productAddCmd)
	productCmd.AddCommand(productDeleteCmd)
	productCmd.AddCommand(productImportCmd)
}
