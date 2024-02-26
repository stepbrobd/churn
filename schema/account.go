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
