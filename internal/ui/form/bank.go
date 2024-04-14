package form

import (
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
	"ysun.co/churn/internal/validator"
	"ysun.co/churn/schema"
)

func FormBankAdd(bank *schema.Bank) error {
	var maxA string
	var maxAP string
	var confirm bool

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Bank Name").
				Value(&bank.BankName).
				Validate(validator.NotNull),
			huh.NewInput().
				Title("Bank Alias").
				Value(&bank.BankAlias).
				Validate(validator.NotNull),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Max Account").
				Value(&maxA).
				Validate(validator.IntConvertible),
			huh.NewInput().
				Title("Max Account Period").
				Value(&maxAP).
				Validate(validator.IntConvertible),
			huh.NewSelect[string]().
				Title("Max Account Scope").
				Options(
					huh.NewOption("All", "all"),
					huh.NewOption("Bank", "bank"),
				).
				Value(&bank.MaxAccountScope.String),
		),
		huh.NewGroup(
			huh.NewConfirm().
				Title("Confirm").
				Description("Are you sure you want to add this bank?").
				Value(&confirm),
		),
	).Run()

	if err != nil {
		return err
	}

	if !confirm {
		os.Exit(0)
	}

	bank.MaxAccount.Int64, _ = strconv.ParseInt(maxA, 10, 64)
	bank.MaxAccountPeriod.Int64, _ = strconv.ParseInt(maxAP, 10, 64)

	return nil
}
