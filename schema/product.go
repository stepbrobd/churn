package schema

type Product struct {
	ProductAlias string  `db:"product_alias,key"` // Product alias
	ProductName  string  `db:"product_name"`      // Product name
	Fee          float32 `db:"fee"`               // Annual fee
	IssuingBank  string  `db:"issuing_bank"`      // Foreign key to bank.bank_alias
}
