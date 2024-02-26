package sqlite

import (
	"database/sql"
	"os"
	"path/filepath"
	"sync"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
	"github.com/samonzeweb/godb"
	adapter "github.com/samonzeweb/godb/adapters/sqlite"
)

var instance = &sqlite{}

type sqlite struct {
	conn  *sql.DB
	query *godb.DB
	once  sync.Once
}

func Init(db string) (*sql.DB, *godb.DB, error) {
	if _, err := os.Stat(db); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(db), os.ModePerm)
	}

	var err error
	instance.once.Do(func() {
		instance.conn, err = sql.Open("sqlite3", db)
		if err != nil {
			return
		}
		err = instance.conn.Ping()
		if err != nil {
			return
		}

		instance.query, err = godb.Open(adapter.Adapter, db)
		if err != nil {
			return
		}
		err = instance.query.CurrentDB().Ping()
		if err != nil {
			return
		}
	})

	return instance.conn, instance.query, nil
}

func Conn() (*sql.DB, error) {
	if instance.conn == nil {
		return nil, sql.ErrConnDone
	}
	return instance.conn, nil
}

func Query() (*godb.DB, error) {
	if instance.query == nil {
		return nil, sql.ErrConnDone
	}
	return instance.query, nil
}

func Close() error {
	var err error

	err = instance.conn.Close()
	if err != nil {
		return err
	}

	err = instance.query.Close()
	if err != nil {
		return err
	}

	return nil
}
