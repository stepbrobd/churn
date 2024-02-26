package migration

import "github.com/stepbrobd/churn/internal/db"

func Exec() error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}

	tx, err := conn.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(Migration20240226183804())
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
