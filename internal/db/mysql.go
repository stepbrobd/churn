package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/samonzeweb/godb/adapters"
	adapter "github.com/samonzeweb/godb/adapters/mysql"
	"github.com/stepbrobd/churn/internal/config"
)

func mysqlInit(cfg *config.Config) (*sql.DB, adapters.Adapter, error) {
	dsn, err := mysql.ParseDSN(cfg.DB.DSN)
	if err != nil {
		return nil, nil, err
	}
	db := dsn.DBName
	dsn.DBName = ""

	conn, err := sql.Open(cfg.DB.Driver, dsn.FormatDSN())
	if err != nil {
		return nil, nil, err
	}

	conn.Exec("CREATE DATABASE IF NOT EXISTS " + db)
	conn.Exec("USE " + db)

	return conn, adapter.Adapter, nil
}
