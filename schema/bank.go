package schema

import (
	"database/sql"
	"fmt"

	"github.com/guregu/null/v5"
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

	// if max_account is not set, insert NULL
	var maxAccount string
	if !b.MaxAccount.Valid {
		maxAccount = "NULL"
	} else {
		maxAccount = fmt.Sprintf("%d", b.MaxAccount.Int64)
	}

	// if max_account_period is not set, insert NULL
	var maxAccountPeriod string
	if !b.MaxAccountPeriod.Valid {
		maxAccountPeriod = "NULL"
	} else {
		maxAccountPeriod = fmt.Sprintf("%d", b.MaxAccountPeriod.Int64)
	}

	// if max_account_scope is not set, insert NULL
	var maxAccountScope string
	if !b.MaxAccountScope.Valid {
		maxAccountScope = "NULL"
	} else {
		maxAccountScope = fmt.Sprintf("'%s'", b.MaxAccountScope.String)
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
