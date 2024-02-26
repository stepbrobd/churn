package schema

type Product struct {
	ID     string `db:"id,key"` // UUIDv7
	Name   string `db:"name"`   // Product name
	Issuer string `db:"issuer"` // Foreign key to Bank.ID
	Fee    int    `db:"fee"`    // Annual fee
}

func (p *Product) Schema() string {
	return `CREATE TABLE IF NOT EXISTS product (
	id VARCHAR(36) PRIMARY KEY,
	name TEXT NOT NULL,
	issuer VARCHAR(36) NOT NULL,
	fee INTEGER NOT NULL
);`
}
