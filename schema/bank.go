package schema

type Bank struct {
	ID               string // UUIDv7
	Name             string // Bank name
	MaxAccount       int    // Maximum number of accounts within a certain period
	MaxAccountPeriod int    // Period of maximum number of accounts, in days
}

func (b *Bank) Schema() string {
	return `CREATE TABLE IF NOT EXISTS bank (
	id VARCHAR(36) PRIMARY KEY,
	name TEXT NOT NULL,
	max_account INTEGER NOT NULL,
	max_account_period INTEGER NOT NULL
);`
}
