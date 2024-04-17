package form

import (
	"strconv"
	"time"

	"github.com/charmbracelet/huh"
	"ysun.co/churn/internal/db"
	"ysun.co/churn/internal/validator"
	"ysun.co/churn/schema"
)

func FormTxAdd(tx *schema.Tx) error {
	var err error
	amount := func() string {
		if tx.Amount == 0 {
			return ""
		}
		return strconv.FormatFloat(tx.Amount, 'f', -1, 64)
	}()
	timestamp := func() string {
		if tx.TxTimestamp.IsZero() {
			return ""
		}
		return tx.TxTimestamp.Format("2006-01-02")
	}()
	var confirm bool

	db := db.Query()
	accounts := make([]*schema.Account, 0)
	err = db.SelectFrom("account").Do(&accounts)
	if err != nil {
		return err
	}

	products := make([]*schema.Product, 0)
	err = db.SelectFrom("product").Do(&products)
	if err != nil {
		return err
	}

	options := make([]huh.Option[int], 0)
	for _, account := range accounts {
		product := schema.Product{}
		for _, p := range products {
			if p.ID == account.ProductID {
				product = *p
				break
			}
		}

		options = append(options, huh.NewOption(product.ProductName, account.ID))
	}

	err = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Account").
				Options(
					options...,
				).
				Value(&tx.AccountID),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Amount").
				Value(&amount).
				Validate(validator.FloatConvertible),
			huh.NewSelect[string]().
				Title("Category").
				Options(
					huh.NewOption("Dining", "dining"),
					huh.NewOption("Gas", "gas"),
					huh.NewOption("Grocery", "grocery"),
					huh.NewOption("Mobile", "mobile"),
					huh.NewOption("Travel", "travel"),
					huh.NewOption("Other", "other"),
				).
				Value(&tx.Category),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Timestamp").
				Placeholder("YYYY-MM-DD").
				Value(&timestamp).
				Validate(validator.DateConvertible),
			huh.NewText().CharLimit(255).
				Title("Note").
				Value(&tx.Note.String),
		),
		huh.NewGroup(
			huh.NewConfirm().
				Title("Confirm").
				Description("Are you sure you want to add this transaction?").
				Value(&confirm),
		),
	).Run()

	if err != nil {
		return err
	}

	if !confirm {
		return nil
	}

	tx.Amount, _ = strconv.ParseFloat(amount, 64)
	tx.TxTimestamp, _ = time.Parse("2006-01-02", timestamp)

	return nil
}
