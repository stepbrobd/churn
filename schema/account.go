package schema

import "time"

type Account struct {
	ID        int       `db:"id,key,auto" json:"id"`
	ProductID int       `db:"product_id" json:"product_id"`
	Opened    time.Time `db:"opened" json:"opened"`
	Closed    time.Time `db:"closed" json:"closed"`
	CL        float32   `db:"cl" json:"cl"`
}
