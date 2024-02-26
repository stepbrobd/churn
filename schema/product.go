package schema

type Product struct {
	ID     string // UUIDv7
	Name   string // Product name
	Issuer string // Foreign key to Bank.ID
	Fee    int    // Annual fee
}

func (p *Product) Schema() string {
	return `CREATE TABLE IF NOT EXISTS product (
	id VARCHAR(36) PRIMARY KEY,
	name TEXT NOT NULL,
	issuer VARCHAR(36) NOT NULL,
	fee INTEGER NOT NULL
);`
}
