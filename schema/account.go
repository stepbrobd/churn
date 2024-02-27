package schema

import "time"

type Account struct {
	ID           string    `db:"id,key"`        // UUIDv7
	AccountAlias string    `db:"account_alias"` // User defined name
	ProductID    string    `db:"product_id"`    // Foreign key to product.id
	Opened       time.Time `db:"opened"`        // Account opening date
	Closed       time.Time `db:"closed"`        // Account closing date, nil if not closed
	CL           float32   `db:"cl"`            // Credit limit, nil if charge card
}
