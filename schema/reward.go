package schema

import "database/sql"

type Reward struct {
	ID        int     `db:"id,key,auto" json:"id"`
	Category  string  `db:"category" json:"category"`
	Unit      string  `db:"unit" json:"unit"`
	Reward    float64 `db:"reward" json:"reward"`
	ProductID int     `db:"product_id" json:"product_id"`
}

func (r *Reward) Add(db *sql.DB) (sql.Result, error) {
	var id int
	err := db.QueryRow("SELECT COALESCE(MAX(id), 0) + 1 FROM reward").Scan(&id)
	if err != nil {
		return nil, err
	}
	r.ID = id

	stmt := "INSERT INTO reward (id, category, unit, reward, product_id) VALUES (?, ?, ?, ?, ?)"

	return db.Exec(stmt, r.ID, r.Category, r.Unit, r.Reward, r.ProductID)
}

func (r *Reward) Delete(db *sql.DB) (sql.Result, error) {
	stmt := "DELETE FROM reward WHERE id = ?"
	return db.Exec(stmt, r.ID)
}
