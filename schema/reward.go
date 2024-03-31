package schema

type Reward struct {
	ID        int     `db:"id,key,auto" json:"id"`
	Category  string  `db:"category" json:"category"`
	Unit      string  `db:"unit" json:"unit"`
	Reward    float32 `db:"reward" json:"reward"`
	ProductID int     `db:"product_id" json:"product_id"`
}
