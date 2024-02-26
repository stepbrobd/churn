package config

import (
	"os"
	"path/filepath"

	"github.com/caarlos0/env/v10"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ConfigPath string   `env:"CONFIG_PATH" yaml:"config_path"`
	ConfigSrc  string   `env:"CONFIG_SRC" yaml:"config_src"`
	DB         DBConfig `envPrefix:"DB_" yaml:"db"`
}

type DBConfig struct {
	Driver string `env:"DRIVER" yaml:"driver"` // Currently supported: `mysql`, `sqlite3`
	DSN    string `env:"DSN" yaml:"dsn"`
}

func Default() *Config {
	return &Config{
		ConfigPath: os.Getenv("HOME") + "/.config/churn",
		ConfigSrc:  "churn.yaml",
		DB: DBConfig{
			Driver: "sqlite3",
			DSN:    "file://" + os.Getenv("HOME") + "/.local/share/churn/churn.db",
		},
	}
}

func (c *Config) Exists() bool {
	_, err := os.Stat(filepath.Join(c.ConfigPath, c.ConfigSrc))
	return !os.IsNotExist(err)
}

func (c *Config) Parse() error {
	f, err := os.Open(filepath.Join(c.ConfigPath, c.ConfigSrc))
	if err != nil {
		return err
	}
	defer f.Close()

	// Parse YAML first
	if err := yaml.NewDecoder(f).Decode(c); err != nil {
		return err
	}

	// Environment variables override YAML
	if err := env.ParseWithOptions(c, env.Options{Prefix: "CHURN_"}); err != nil {
		return err
	}

	return nil
}
