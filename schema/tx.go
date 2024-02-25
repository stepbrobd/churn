package schema

import "time"

type Tx struct {
	ID          string    // UUIDv7
	Timestamp   time.Time // Transaction timestamp
	Account     string    // Foreign key to Account.ID
	Amount      int       // Transaction amount
	Category    string    // Transaction category: "dining", "travel", "grocery", "fuel", "other"
	Description string    // User defined description
}

func (t *Tx) Schema() string {
	return `CREATE TABLE IF NOT EXISTS tx (
	id VARCHAR(36) PRIMARY KEY,
	timestamp DATETIME NOT NULL,
	account VARCHAR(36) NOT NULL,
	amount INTEGER NOT NULL,
	category TEXT NOT NULL,
	description TEXT
);`
}
