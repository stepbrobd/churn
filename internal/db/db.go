package db

import (
	"database/sql"
	"sync"

	"github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters"
	"ysun.co/churn/internal/config"
)

var instance = &db{}

type db struct {
	once sync.Once
	conn *sql.DB
	adpt adapters.Adapter
}

func Init(cfg *config.Config) error {
	var err error

	instance.once.Do(func() {
		switch cfg.DB.Driver {
		case "mysql":
			instance.conn, instance.adpt, err = mysqlInit(cfg)
		case "sqlite3":
			instance.conn, instance.adpt, err = sqlite3Init(cfg)
		default:
			err = sql.ErrConnDone
		}

		if err != nil {
			return
		}

		err = instance.conn.Ping()
		if err != nil {
			return
		}
	})

	return nil
}

func Connect() (*sql.DB, error) {
	if instance.conn == nil {
		return nil, sql.ErrConnDone
	}

	return instance.conn, nil
}

func Query() *godb.DB {
	return godb.Wrap(instance.adpt, instance.conn)
}

func Close() error {
	return instance.conn.Close()
}

func Tables() ([]string, error) {
	var query string
	switch instance.adpt.DriverName() {
	case "mysql":
		query = "SHOW TABLES;"
	case "sqlite3":
		query = "SELECT name FROM sqlite_master WHERE type = 'table';"
	}

	rows, err := instance.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var table string
		rows.Scan(&table)
		tables = append(tables, table)
	}

	return tables, nil
}
