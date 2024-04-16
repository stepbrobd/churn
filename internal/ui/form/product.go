package form

import (
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
	"ysun.co/churn/internal/db"
	"ysun.co/churn/internal/validator"
	"ysun.co/churn/schema"
)

func FormProductAdd(product *schema.Product) error {
	var af string
	var confirm bool

	db := db.Query()
	banks := make([]*schema.Bank, 0)
	err := db.SelectFrom("bank").Do(&banks)
	if err != nil {
		return err
	}
	options := make([]huh.Option[string], 0)
	for _, bank := range banks {
		options = append(options, huh.NewOption(bank.BankName, bank.BankAlias))
	}

	err = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Product Name").
				Value(&product.ProductName).
				Validate(validator.NotNull),
			huh.NewInput().
				Title("Product Alias").
				Value(&product.ProductAlias).
				Validate(validator.NotNull),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Annual Fee").
				Value(&af).
				Validate(validator.FloatConvertible),
			huh.NewSelect[string]().
				Title("Issuing Bank").
				Options(
					options...,
				).
				Value(&product.IssuingBank),
		),
		huh.NewGroup(
			huh.NewConfirm().
				Title("Confirm").
				Description("Are you sure you want to add this product?").
				Value(&confirm),
		),
	).Run()

	if err != nil {
		return err
	}

	if !confirm {
		os.Exit(0)
	}

	product.Fee, _ = strconv.ParseFloat(af, 64)

	return nil
}
