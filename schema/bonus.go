package schema

import "time"

type Bonus struct {
	ID          int       `db:"id,key,auto" json:"id"`
	BonusType   string    `db:"bonus_type" json:"bonus_type"`
	Spend       float32   `db:"spend" json:"spend"`
	BonusAmount float32   `db:"bonus_amount" json:"bonus_amount"`
	Unit        string    `db:"unit" json:"unit"`
	BonusStart  time.Time `db:"bonus_start" json:"bonus_start"`
	BonusEnd    time.Time `db:"bonus_end" json:"bonus_end"`
	AccountID   int       `db:"account_id" json:"account_id"`
}
