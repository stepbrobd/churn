package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"ysun.co/churn/internal/config"
	"ysun.co/churn/internal/db"
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
