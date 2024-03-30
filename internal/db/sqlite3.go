package db

import (
	"database/sql"
	"net/url"
	"os"
	"path/filepath"

	"github.com/samonzeweb/godb/adapters"
	adapter "github.com/samonzeweb/godb/adapters/sqlite"
	"ysun.co/churn/internal/config"
)

func sqlite3Init(cfg *config.Config) (*sql.DB, adapters.Adapter, error) {
	p := url.PathEscape(cfg.DB.DSN)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(p), os.ModePerm)
	}

	conn, err := sql.Open(cfg.DB.Driver, cfg.DB.DSN)
	if err != nil {
		return nil, nil, err
	}

	return conn, adapter.Adapter, nil
}
