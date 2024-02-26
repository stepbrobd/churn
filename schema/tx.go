package schema

import "time"

type Tx struct {
	ID          string    `db:"id,key"`      // UUIDv7
	Timestamp   time.Time `db:"timestamp"`   // Transaction timestamp
	Amount      int       `db:"amount"`      // Transaction amount
	Category    string    `db:"category"`    // Transaction category: "dining", "travel", "grocery", "fuel", "other"
	Description string    `db:"description"` // User defined description
	Account     string    `db:"account"`     // Foreign key to Account.ID
}

func (t *Tx) Schema() string {
	return `CREATE TABLE IF NOT EXISTS tx (
	id VARCHAR(36) PRIMARY KEY,
	timestamp DATETIME NOT NULL,
	amount INTEGER NOT NULL,
	category TEXT NOT NULL,
	description TEXT,
	account VARCHAR(36) NOT NULL,
	FOREIGN KEY (account) REFERENCES account (id)
);`
}
