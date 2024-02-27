package schema

type Bank struct {
	ID               string `db:"id,key"`             // UUIDv7
	BankName         string `db:"bank_name"`          // Bank name
	MaxAccount       int    `db:"max_account"`        // Maximum number of accounts within a certain period
	MaxAccountPeriod int    `db:"max_account_period"` // Period of maximum number of accounts, in days
}
