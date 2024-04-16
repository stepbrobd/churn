package schema

import (
	"database/sql"
	"time"
)

type Bonus struct {
	ID          int       `db:"id,key,auto" json:"id"`
	BonusType   string    `db:"bonus_type" json:"bonus_type"`
	Spend       float64   `db:"spend" json:"spend"`
	BonusAmount float64   `db:"bonus_amount" json:"bonus_amount"`
	Unit        string    `db:"unit" json:"unit"`
	BonusStart  time.Time `db:"bonus_start" json:"bonus_start"`
	BonusEnd    time.Time `db:"bonus_end" json:"bonus_end"`
	AccountID   int       `db:"account_id" json:"account_id"`
}

func (b *Bonus) Add(db *sql.DB) (sql.Result, error) {
	var id int
	err := db.QueryRow("SELECT COALESCE(MAX(id), 0) + 1 FROM bonus").Scan(&id)
	if err != nil {
		return nil, err
	}
	b.ID = id

	// foreign key constraint enforced at frontend
	stmt := "INSERT INTO bonus (id, bonus_type, spend, bonus_amount, unit, bonus_start, bonus_end, account_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	return db.Exec(stmt, b.ID, b.BonusType, b.Spend, b.BonusAmount, b.Unit, b.BonusStart, b.BonusEnd, b.AccountID)
}

func (b *Bonus) Delete(db *sql.DB) (sql.Result, error) {
	stmt := "DELETE FROM bonus WHERE id = ?"
	return db.Exec(stmt, b.ID)
}
