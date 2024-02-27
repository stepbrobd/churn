package schema

import "time"

type Bonus struct {
	ID          string    `db:"id,key"`       // UUIDv7
	BonusType   string    `db:"bonus_type"`   // Bonus category: "signup", "retention", "spend"
	Spend       float32   `db:"spend"`        // Minimum spend to get the bonus
	BonusAmount float32   `db:"bonus_amount"` // Bonus amount
	Unit        string    `db:"unit"`         // Bonus unit: "point", "dollar"
	BonusStart  time.Time `db:"bonus_start"`  // Bonus start date
	BonusEnd    time.Time `db:"bonus_end"`    // Bonus end date
	AccountID   string    `db:"account_id"`   // Foreign key to account.id
}
