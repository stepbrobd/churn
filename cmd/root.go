package cmd

import (
	"os"
	"reflect"
	"slices"

	"github.com/spf13/cobra"
	"github.com/stepbrobd/churn/internal/config"
	"github.com/stepbrobd/churn/internal/sqlite"
	"github.com/stepbrobd/churn/schema"
)

func preRun(cmd *cobra.Command, args []string) error {
	cfg := config.Default()
	if cfg.Exists() {
		if err := cfg.Parse(); err != nil {
			return err
		}
	}

	conn, _, err := sqlite.Init(cfg.DatabasePath())
	if err != nil {
		return err
	}

	actual := make([]string, 0)
	expected := map[string]interface{}{
		"account": &schema.Account{},
		"bank":    &schema.Bank{},
		"bonus":   &schema.Bonus{},
		"product": &schema.Product{},
		"reward":  &schema.Reward{},
		"tx":      &schema.Tx{},
	}

	rows, err := conn.Query("SELECT name FROM sqlite_master WHERE type = 'table'")
	if err != nil {
		return err
	}
	for rows.Next() {
		var name string
		rows.Scan(&name)
		actual = append(actual, name)
	}

	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	for _, t := range reflect.ValueOf(expected).MapKeys() {
		if !slices.Contains(actual, t.String()) {
			query := reflect.ValueOf(expected).MapIndex(t).Elem().MethodByName("Schema").Call(nil)[0].String()
			_, err := tx.Exec(query)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	tx.Commit()

	return nil
}

func postRun(cmd *cobra.Command, args []string) error {
	return sqlite.Close()
}

var rootCmd = &cobra.Command{
	Use:                "churn",
	Short:              "Credit card churning management CLI",
	PersistentPreRunE:  preRun,
	PersistentPostRunE: postRun,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
}
