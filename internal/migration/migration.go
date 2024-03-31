package migration

import (
	"embed"
	"io"

	"ysun.co/churn/internal/db"
)

//go:embed *.sql
var migrations embed.FS

func Exec() error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}

	files, err := migrations.ReadDir(".")
	if err != nil {
		return err
	}

	tx, err := conn.Begin()
	if err != nil {
		return err
	}

	for _, file := range files {
		f, err := migrations.Open(file.Name())
		if err != nil {
			tx.Rollback()
			return err
		}
		defer f.Close()

		q, err := io.ReadAll(f)
		if err != nil {
			tx.Rollback()
			return err
		}

		_, err = tx.Exec(string(q))
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}
