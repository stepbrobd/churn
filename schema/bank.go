package schema

type Bank struct {
	BankAlias        string `db:"bank_alias,key"`     // Unique bank alias
	BankName         string `db:"bank_name"`          // Bank name
	MaxAccount       int    `db:"max_account"`        // Maximum number of accounts within a certain period
	MaxAccountPeriod int    `db:"max_account_period"` // Period of maximum number of accounts, in days
	MaxAccountScope  string `db:"max_account_scope"`  // Scope of maximum number of accounts, either "all" or "bank"
}
