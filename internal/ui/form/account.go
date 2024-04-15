package form

import (
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/guregu/null/v5"
	"ysun.co/churn/internal/db"
	"ysun.co/churn/internal/validator"
	"ysun.co/churn/schema"
)

func FormAccountAdd(account *schema.Account) error {
	var cl string
	var opened string
	var closed string
	var confirm bool

	db := db.Query()
	products := make([]*schema.Product, 0)
	err := db.Select(&products).Do()
	if err != nil {
		return err
	}
	options := make([]huh.Option[int], 0)
	for _, product := range products {
		options = append(options, huh.NewOption(product.ProductName, product.ID))
	}

	err = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Product").
				Options(
					options...,
				).
				Value(&account.ProductID),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Credit Limit").
				Value(&cl).
				Validate(validator.FloatConvertible),
			huh.NewInput().
				Title("Opened On").
				Value(&opened).
				Validate(validator.DateConvertibleNullable),
			huh.NewInput().
				Title("Closed On").
				Value(&closed).
				Validate(validator.DateConvertibleNullable),
		),
		huh.NewGroup(
			huh.NewConfirm().
				Title("Confirm").
				Description("Are you sure you want to add this account?").
				Value(&confirm),
		),
	).Run()

	if err != nil {
		return err
	}

	if !confirm {
		os.Exit(0)
	}

	account.CL, _ = strconv.ParseFloat(cl, 64)
	if opened != "" {
		t, _ := time.Parse("2006-01-02", opened)
		account.Opened = null.TimeFrom(t)
	}
	if closed != "" {
		t, _ := time.Parse("2006-01-02", closed)
		account.Closed = null.TimeFrom(t)
	}

	return nil
}
