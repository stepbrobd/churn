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

// churn product --
var productCmd = &cobra.Command{
	Use:   "product",
	Short: "Manage product (add, delete, edit)",
	Long:  "Open an interactive TUI to manage product, add, delete, edit, etc.",
	Run: func(cmd *cobra.Command, args []string) {
		db := db.Query()
		products := make([]*schema.Product, 0)
		err := db.SelectFrom("product").Do(&products)
		if err != nil {
			fmt.Println("Failed to fetch products")
			os.Exit(1)
		}

		columns := []table.Column{
			{Title: "Name", Width: 25},
			{Title: "Alias", Width: 15},
			{Title: "Annual Fee", Width: 10},
			{Title: "Issuing Bank", Width: 12},
		}

		rows := make([]table.Row, 0)
		for _, product := range products {
			rows = append(rows, table.Row{
				product.ProductName,
				product.ProductAlias,
				fmt.Sprintf("$%.2f", product.Fee),
				product.IssuingBank,
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

// churn product add --
var productAddCmd = &cobra.Command{
	Use:   "add <product alias>",
	Short: "Add a product",
	Long:  "Add a product by its alias, this will create a new product with the given alias.",
	Args:  cobra.ExactArgs(1), // product alias, by, bank alias
	Run: func(cmd *cobra.Command, args []string) {
		product := &schema.Product{
			ProductAlias: args[0],
		}

		err := form.FormProductAdd(product)
		if err != nil {
			fmt.Println("Failed to add product")
			os.Exit(1)
		}

		db, _ := db.Connect()
		product.Add(db)
	},
}

// churn product edit --
var productEditCmd = &cobra.Command{
	Use:   "edit <product alias>",
	Short: "Edit a product", // product alias
	Long:  "Edit a product by its alias, this will update the product with the given alias.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		q := db.Query()
		products := make([]*schema.Product, 0)
		err := q.SelectFrom("product").Do(&products)
		if err != nil {
			fmt.Println("Failed to fetch products")
			os.Exit(1)
		}

		product := &schema.Product{}
		for _, p := range products {
			if p.ProductAlias == args[0] {
				product = p
				break
			}
		}

		err = form.FormProductAdd(product)
		if err != nil {
			fmt.Println("Failed to edit product")
			os.Exit(1)
		}

		db, _ := db.Connect()
		product.Update(db)
	},
}

// churn product delete --
var forceProductDeletion bool
var productDeleteCmd = &cobra.Command{
	Use:   "delete <product alias>",
	Short: "Delete a product", // product alias
	Long:  "Delete a product by its alias, note that this will also delete all accounts, rewards, transactions, etc. associated with the product.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !forceProductDeletion {
			fmt.Println("This will delete the product and all its associated data.")
			if !lib.Confirm() {
				return
			}
		}

		product := &schema.Product{
			ProductAlias: args[0],
		}

		db, _ := db.Connect()
		product.Delete(db)

		fmt.Printf("Deleted product %s\n", product.ProductAlias)
	},
}

// churn product import --
var productImportCmd = &cobra.Command{
	Use:   "import <uri>",
	Short: "Import product(s)", // either a file or a http(s) uri
	Long:  "Import product(s) from a local JSON file or a http(s) uri that returns a JSON file.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		uri := args[0]
		products := make([]*schema.Product, 0)

		err := lib.Import(uri, &products)
		if err != nil {
			fmt.Println("Failed to import product(s)")
			os.Exit(1)
		}

		db, _ := db.Connect()
		for _, product := range products {
			product.Add(db)
		}

		fmt.Printf("Imported %d products\n", len(products))
	},
}

func init() {
	productDeleteCmd.Flags().BoolVarP(&forceProductDeletion, "force", "f", false, "Force deletion the product and all its associated data")

	rootCmd.AddCommand(productCmd)
	productCmd.AddCommand(productAddCmd)
	productCmd.AddCommand(productEditCmd)
	productCmd.AddCommand(productDeleteCmd)
	productCmd.AddCommand(productImportCmd)
}
