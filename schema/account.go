package schema

import (
	"database/sql"

	"github.com/guregu/null/v5"
	self "ysun.co/churn/internal/db"
)

type Account struct {
	ID        int       `db:"id,key,auto" json:"id"`
	ProductID int       `db:"product_id" json:"product_id"`
	Opened    null.Time `db:"opened" json:"opened"`
	Closed    null.Time `db:"closed" json:"closed"`
	CL        float64   `db:"cl" json:"cl"`
}

func (a *Account) Add(db *sql.DB) (sql.Result, error) {
	// get the next ID, if empty in db, set to 1 else increment by 1
	var id int
	err := db.QueryRow("SELECT COALESCE(MAX(id), 0) + 1 FROM account").Scan(&id)
	if err != nil {
		return nil, err
	}
	a.ID = id

	var opened string
	if !a.Opened.Valid {
		opened = "NULL"
	} else {
		opened = a.Opened.Time.Format("2006-01-02") // based on docs
	}

	var closed string
	if !a.Closed.Valid {
		closed = "NULL"
	} else {
		closed = a.Closed.Time.Format("2006-01-02") // based on docs
	}

	// foreign key constraint enforced at frontend
	stmt := "INSERT INTO account (id, product_id, opened, closed, cl) VALUES (?, ?, ?, ?, ?)"

	return self.ExecInTx(db, stmt, a.ID, a.ProductID, opened, closed, a.CL)
}

func (a *Account) Update(db *sql.DB) (sql.Result, error) {
	var opened string
	if !a.Opened.Valid {
		opened = "NULL"
	} else {
		opened = a.Opened.Time.Format("2006-01-02") // based on docs
	}

	var closed string
	if !a.Closed.Valid {
		closed = "NULL"
	} else {
		closed = a.Closed.Time.Format("2006-01-02") // based on docs
	}

	stmt := "UPDATE account SET product_id = ?, opened = ?, closed = ?, cl = ? WHERE id = ?"
	return self.ExecInTx(db, stmt, a.ProductID, opened, closed, a.CL, a.ID)
}

func (a *Account) Delete(db *sql.DB) (sql.Result, error) {
	stmt := "DELETE FROM account WHERE id = ?"
	return self.ExecInTx(db, stmt, a.ID)
}
