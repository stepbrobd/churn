package schema

import "time"

type Tx struct {
	ID          string    `db:"id,key"`       // UUIDv7
	TxTimestamp time.Time `db:"tx_timestamp"` // Transaction timestamp
	Amount      float32   `db:"amount"`       // Transaction amount
	Category    string    `db:"category"`     // Transaction category: "dining", "travel", "grocery", "fuel", "other"
	Note        string    `db:"note"`         // User defined description
	AccountID   string    `db:"account_id"`   // Foreign key to account.id
}
