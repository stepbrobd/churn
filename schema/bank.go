package schema

import (
	"database/sql"
	"fmt"

	"github.com/guregu/null/v5"
	self "ysun.co/churn/internal/db"
)

type Bank struct {
	BankAlias        string      `db:"bank_alias,key" json:"bank_alias"`
	BankName         string      `db:"bank_name" json:"bank_name"`
	MaxAccount       null.Int64  `db:"max_account" json:"max_account"`
	MaxAccountPeriod null.Int64  `db:"max_account_period" json:"max_account_period"`
	MaxAccountScope  null.String `db:"max_account_scope" json:"max_account_scope"`
}

func (b *Bank) Add(db *sql.DB) (sql.Result, error) {
	// bank_alias and bank_name are required
	if b.BankAlias == "" || b.BankName == "" {
		return nil, fmt.Errorf("bank_alias and bank_name are required")
	}

	stmt := "INSERT INTO bank (bank_alias, bank_name, max_account, max_account_period, max_account_scope) VALUES (?, ?, ?, ?, ?)"

	return self.ExecInTx(db, stmt, b.BankAlias, b.BankName, b.MaxAccount.Int64, b.MaxAccountPeriod.Int64, b.MaxAccountScope.String)
}

func (b *Bank) Update(db *sql.DB) (sql.Result, error) {
	// bank_alias is required
	if b.BankAlias == "" {
		return nil, fmt.Errorf("bank_alias is required")
	}

	stmt := "UPDATE bank SET bank_name = ?, max_account = ?, max_account_period = ?, max_account_scope = ? WHERE bank_alias = ?"
	return self.ExecInTx(db, stmt, b.BankName, b.MaxAccount.Int64, b.MaxAccountPeriod.Int64, b.MaxAccountScope.String, b.BankAlias)
}

func (b *Bank) Delete(db *sql.DB) (sql.Result, error) {
	if b.BankAlias == "" {
		return nil, fmt.Errorf("bank_alias is required")
	}

	stmt := "DELETE FROM bank WHERE bank_alias = ?"
	return self.ExecInTx(db, stmt, b.BankAlias)
}
