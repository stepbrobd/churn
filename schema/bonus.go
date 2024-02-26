package schema

import "time"

type Bonus struct {
	ID      string    `db:"id,key"`  // UUIDv7
	Type    string    `db:"type"`    // Bonus category: "signup", "retention", "spend"
	Spend   int       `db:"spend"`   // Minimum spend to get the bonus
	Bonus   int       `db:"bonus"`   // Bonus amount
	Unit    string    `db:"unit"`    // Bonus unit: "point", "dollar"
	Start   time.Time `db:"start"`   // Bonus start date
	End     time.Time `db:"end"`     // Bonus end date
	Account string    `db:"account"` // Foreign key to Account.ID
}
