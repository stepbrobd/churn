package migration

import (
	"embed"
	"io"

	"ysun.co/churn/internal/db"
)

//go:embed *.sql
var migrations embed.FS

func List() ([]string, error) {
	files, err := migrations.ReadDir(".")
	if err != nil {
		return nil, err
	}

	names := make([]string, 0)
	for _, file := range files {
		names = append(names, file.Name())
	}

	return names, nil
}

func Exec(name string) error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}

	f, err := migrations.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	q, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	_, err = conn.Exec(string(q))
	if err != nil {
		return err
	}

	return nil
}

func ExecAll() error {
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
