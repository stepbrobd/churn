package schema

type Bank struct {
	ID               string `db:"id,key"`             // UUIDv7
	Name             string `db:"name"`               // Bank name
	MaxAccount       int    `db:"max_account"`        // Maximum number of accounts within a certain period
	MaxAccountPeriod int    `db:"max_account_period"` // Period of maximum number of accounts, in days
}

func (b *Bank) Schema() string {
	return `CREATE TABLE IF NOT EXISTS bank (
	id VARCHAR(36) PRIMARY KEY,
	name TEXT NOT NULL,
	max_account INTEGER NOT NULL,
	max_account_period INTEGER NOT NULL
);`
}
