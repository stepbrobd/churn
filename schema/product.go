package schema

type Product struct {
	ID     string `db:"id,key"` // UUIDv7
	Name   string `db:"name"`   // Product name
	Issuer string `db:"issuer"` // Foreign key to Bank.ID
	Fee    int    `db:"fee"`    // Annual fee
}
