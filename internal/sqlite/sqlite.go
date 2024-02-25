package sqlite

import (
	"database/sql"
	"os"
	"path/filepath"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var instance = &sqlite{}

type sqlite struct {
	conn *sql.DB
	once sync.Once
}

func Init(db string) (*sql.DB, error) {
	if _, err := os.Stat(db); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(db), os.ModePerm)
	}

	var err error
	instance.once.Do(func() {
		instance.conn, err = sql.Open("sqlite3", db)
	})
	return instance.conn, err
}

func Open() (*sql.DB, error) {
	if instance.conn == nil {
		return nil, sql.ErrConnDone
	}
	return instance.conn, nil
}

func Close() error {
	return instance.conn.Close()
}
