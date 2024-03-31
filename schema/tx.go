package schema

import "time"

type Tx struct {
	ID          int       `db:"id,key,auto" json:"id"`
	TxTimestamp time.Time `db:"tx_timestamp" json:"tx_timestamp"`
	Amount      float32   `db:"amount" json:"amount"`
	Category    string    `db:"category" json:"category"`
	Note        string    `db:"note" json:"note"`
	AccountID   int       `db:"account_id" json:"account_id"`
}
