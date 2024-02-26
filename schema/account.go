package schema

import "time"

type Account struct {
	ID      string    `db:"id,key"`  // UUIDv7
	Name    string    `db:"name"`    // User defined name
	Product string    `db:"product"` // Foreign key to Product.ID
	Opened  time.Time `db:"opened"`  // Account opening date
	Closed  time.Time `db:"closed"`  // Account closing date, nil if not closed
	CL      int       `db:"cl"`      // Credit limit, nil if charge card
}

func (a *Account) Schema() string {
	return `CREATE TABLE IF NOT EXISTS account (
	id VARCHAR(36) PRIMARY KEY,
	name TEXT NOT NULL,
	product VARCHAR(36) NOT NULL,
	opened DATETIME NOT NULL,
	closed DATETIME,
	cl INTEGER
);`
}
