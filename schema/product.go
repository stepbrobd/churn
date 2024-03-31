package schema

type Product struct {
	ID           int     `db:"id,key,auto" json:"id"`
	ProductAlias string  `db:"product_alias" json:"product_alias"`
	ProductName  string  `db:"product_name" json:"product_name"`
	Fee          float32 `db:"fee" json:"fee"`
	IssuingBank  string  `db:"issuing_bank" json:"issuing_bank"`
}
