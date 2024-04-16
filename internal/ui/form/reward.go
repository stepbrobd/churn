package form

import (
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
	"ysun.co/churn/internal/db"
	"ysun.co/churn/internal/validator"
	"ysun.co/churn/schema"
)

func FormRewardAdd(reward *schema.Reward) error {
	var amount string
	var confirm bool

	db := db.Query()
	products := make([]*schema.Product, 0)
	err := db.SelectFrom("product").Do(&products)
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
				Value(&reward.ProductID),
		),
		huh.NewGroup(
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
				Value(&reward.Category),
			huh.NewInput().
				Title("Reward Amount").
				Placeholder("If cashback, enter the percentage (if 3%, type \"3\"), otherwise enter the amount").
				Value(&amount).
				Validate(validator.FloatConvertible),
			huh.NewSelect[string]().
				Title("Unit").
				Options(
					huh.NewOption("Percentage", "percentage"),
					huh.NewOption("Other", "other"),
				).
				Value(&reward.Unit),
		),
		huh.NewGroup(
			huh.NewConfirm().
				Title("Confirm").
				Description("Are you sure you want to add this reward category?").
				Value(&confirm),
		),
	).Run()

	if err != nil {
		return err
	}

	if !confirm {
		os.Exit(0)
	}

	reward.Reward, _ = strconv.ParseFloat(amount, 64)

	return nil
}
