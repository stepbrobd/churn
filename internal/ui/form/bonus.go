package form

import (
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/huh"
	"ysun.co/churn/internal/db"
	"ysun.co/churn/internal/validator"
	"ysun.co/churn/schema"
)

func FormBonusAdd(bonus *schema.Bonus) error {
	var err error
	var spend string
	var amount string
	var start string
	var end string
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
				Value(&bonus.AccountID),
			huh.NewSelect[string]().
				Title("Bonus Type").
				Options(
					huh.NewOption("Sign-up Bonus", "sign-up"),
					huh.NewOption("Referral Bonus", "referral"),
					huh.NewOption("Spending Bonus", "spending"),
					huh.NewOption("Retention Bonus", "retention"),
					huh.NewOption("Other", "other"),
				).
				Value(&bonus.BonusType),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Spending Requirement").
				Value(&spend).
				Validate(validator.FloatConvertible),
			huh.NewInput().
				Title("Bonus Amount").
				Value(&amount).
				Validate(validator.FloatConvertible),
			huh.NewSelect[string]().
				Title("Bonus Unit").
				Options(
					huh.NewOption("Points", "points"),
					huh.NewOption("Dollars", "dollars"),
				).
				Value(&bonus.Unit),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Bonus Start Date").
				Value(&start).
				Validate(validator.DateConvertible),
			huh.NewInput().
				Title("Bonus End Date").
				Value(&end).
				Validate(validator.DateConvertible),
		),
		huh.NewGroup(
			huh.NewConfirm().
				Title("Confirm").
				Description("Are you sure you want to add this bonus?").
				Value(&confirm),
		),
	).Run()

	if err != nil {
		return err
	}

	if !confirm {
		os.Exit(0)
	}

	bonus.Spend, _ = strconv.ParseFloat(spend, 64)
	bonus.BonusAmount, _ = strconv.ParseFloat(amount, 64)
	bonus.BonusStart, _ = time.Parse("2006-01-02", start)
	bonus.BonusEnd, _ = time.Parse("2006-01-02", end)

	return nil
}
