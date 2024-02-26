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
