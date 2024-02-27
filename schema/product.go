package schema

type Product struct {
	ID          string  `db:"id,key"`       // UUIDv7
	ProductName string  `db:"product_name"` // Product name
	Fee         float32 `db:"fee"`          // Annual fee
	IssuingBank string  `db:"issuing_bank"` // Foreign key to bank.id
}
