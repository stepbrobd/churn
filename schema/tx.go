package schema

import (
	"database/sql"
	"time"

	"github.com/guregu/null/v5"
	self "ysun.co/churn/internal/db"
)

type Tx struct {
	ID          int         `db:"id,key,auto" json:"id"`
	TxTimestamp time.Time   `db:"tx_timestamp" json:"tx_timestamp"`
	Amount      float64     `db:"amount" json:"amount"`
	Category    string      `db:"category" json:"category"`
	Note        null.String `db:"note" json:"note"`
	AccountID   int         `db:"account_id" json:"account_id"`
}

func (t *Tx) Add(db *sql.DB) (sql.Result, error) {
	var id int
	err := db.QueryRow("SELECT COALESCE(MAX(id), 0) + 1 FROM tx").Scan(&id)
	if err != nil {
		return nil, err
	}
	t.ID = id

	stmt := "INSERT INTO tx (id, tx_timestamp, amount, category, note, account_id) VALUES (?, ?, ?, ?, ?, ?)"

	return self.ExecInTx(db, stmt, t.ID, t.TxTimestamp, t.Amount, t.Category, t.Note.String, t.AccountID)
}

func (t *Tx) Update(db *sql.DB) (sql.Result, error) {
	stmt := "UPDATE tx SET tx_timestamp = ?, amount = ?, category = ?, note = ?, account_id = ? WHERE id = ?"
	return self.ExecInTx(db, stmt, t.TxTimestamp, t.Amount, t.Category, t.Note.String, t.AccountID, t.ID)
}

func (t *Tx) Delete(db *sql.DB) (sql.Result, error) {
	stmt := "DELETE FROM tx WHERE id = ?"
	return self.ExecInTx(db, stmt, t.ID)
}
