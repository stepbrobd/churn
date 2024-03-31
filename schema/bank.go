package schema

import (
	"database/sql"
	"fmt"
)

type Bank struct {
	BankAlias        string `db:"bank_alias,key" json:"bank_alias"`
	BankName         string `db:"bank_name" json:"bank_name"`
	MaxAccount       int    `db:"max_account" json:"max_account"`
	MaxAccountPeriod int    `db:"max_account_period" json:"max_account_period"`
	MaxAccountScope  string `db:"max_account_scope" json:"max_account_scope"`
}

func (b *Bank) Add(db *sql.DB) (sql.Result, error) {
	// bank_alias and bank_name are required
	if b.BankAlias == "" || b.BankName == "" {
		return nil, fmt.Errorf("bank_alias and bank_name are required")
	}

	// if max_account is not set, insert NULL
	var maxAccount string
	if b.MaxAccount == 0 {
		maxAccount = "NULL"
	} else {
		maxAccount = fmt.Sprintf("%d", b.MaxAccount)
	}

	// if max_account_period is not set, insert NULL
	var maxAccountPeriod string
	if b.MaxAccountPeriod == 0 {
		maxAccountPeriod = "NULL"
	} else {
		maxAccountPeriod = fmt.Sprintf("%d", b.MaxAccountPeriod)
	}

	// if max_account_scope is not set, insert NULL
	var maxAccountScope string
	if b.MaxAccountScope == "" {
		maxAccountScope = "NULL"
	} else {
		maxAccountScope = fmt.Sprintf("'%s'", b.MaxAccountScope)
	}

	stmt := fmt.Sprintf(
		"INSERT INTO bank (bank_alias, bank_name, max_account, max_account_period, max_account_scope) VALUES ('%s', '%s', %s, %s, %s)",
		b.BankAlias, b.BankName, maxAccount, maxAccountPeriod, maxAccountScope,
	)

	return db.Exec(stmt)
}

func (b *Bank) Delete(db *sql.DB) (sql.Result, error) {
	if b.BankAlias == "" {
		return nil, fmt.Errorf("bank_alias is required")
	}

	stmt := fmt.Sprintf("DELETE FROM bank WHERE bank_alias = '%s'", b.BankAlias)
	return db.Exec(stmt)
}
