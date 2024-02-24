package schema

import "time"

type Account struct {
	ID      string    // UUIDv7
	Name    string    // User defined name
	Product string    // Foreign key to Product.ID
	Opened  time.Time // Account opening date
	Closed  time.Time // Account closing date, nil if not closed
	CL      int       // Credit limit, nil if charge card
	Note    string    // User defined note
}

func (a *Account) Schema() string {
	return `CREATE TABLE IF NOT EXISTS account (
	id VARCHAR(36) PRIMARY KEY,
	name TEXT NOT NULL,
	product VARCHAR(36) NOT NULL,
	opened DATETIME NOT NULL,
	closed DATETIME,
	cl INTEGER,
	note TEXT
);`
}
