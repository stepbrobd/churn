package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/stepbrobd/churn/internal/config"
	"github.com/stepbrobd/churn/internal/db"
	"github.com/stepbrobd/churn/internal/migration"
)

func preRun(cmd *cobra.Command, args []string) error {
	cfg := config.Default()
	if cfg.Exists() {
		if err := cfg.Parse(); err != nil {
			return err
		}
	}

	err := db.Init(cfg)
	if err != nil {
		return err
	}

	err = migration.Exec()
	if err != nil {
		return err
	}

	return nil
}

func postRun(cmd *cobra.Command, args []string) error {
	return db.Close()
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
