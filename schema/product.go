package schema

import (
	"database/sql"
)

type Product struct {
	ID           int     `db:"id,key,auto" json:"id"`
	ProductAlias string  `db:"product_alias" json:"product_alias"`
	ProductName  string  `db:"product_name" json:"product_name"`
	Fee          float64 `db:"fee" json:"fee"`
	IssuingBank  string  `db:"issuing_bank" json:"issuing_bank"`
}

func (p *Product) Add(db *sql.DB) (sql.Result, error) {
	// get the next ID, if empty in db, set to 1 else increment by 1
	var id int
	err := db.QueryRow("SELECT COALESCE(MAX(id), 0) + 1 FROM product").Scan(&id)
	if err != nil {
		return nil, err
	}
	p.ID = id

	// foreign key constraint enforced at frontend
	stmt := "INSERT INTO product (id, product_alias, product_name, fee, issuing_bank) VALUES (?, ?, ?, ?, ?)"

	return db.Exec(stmt, p.ID, p.ProductAlias, p.ProductName, p.Fee, p.IssuingBank)
}

func (p *Product) Delete(db *sql.DB) (sql.Result, error) {
	stmt := "DELETE FROM product WHERE product_alias = ?"
	return db.Exec(stmt, p.ProductAlias)
}
