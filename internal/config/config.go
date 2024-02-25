package config

import (
	"os"
	"path/filepath"

	"github.com/caarlos0/env/v10"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ConfigPath string `env:"CONFIG_PATH" yaml:"config_path"`
	ConfigSrc  string `env:"CONFIG_SRC" yaml:"config_src"`
	DataPath   string `env:"DATA_PATH" yaml:"data_path"`
	DataSrc    string `env:"DATA_SRC" yaml:"data_src"`
	// Encrypted SQLite
	// ...
}

func Default() *Config {
	return &Config{
		ConfigPath: os.Getenv("HOME") + "/.config/churn",
		ConfigSrc:  "churn.yaml",
		DataPath:   os.Getenv("HOME") + "/.local/share/churn",
		DataSrc:    "churn.db",
	}
}

func (c *Config) Exists() bool {
	_, err := os.Stat(c.ConfigurationPath())
	return !os.IsNotExist(err)
}

func (c *Config) Parse() error {
	f, err := os.Open(filepath.Join(c.ConfigPath, c.ConfigSrc))
	if err != nil {
		return err
	}
	defer f.Close()

	if err := yaml.NewDecoder(f).Decode(c); err != nil {
		return err
	}

	if err := env.ParseWithOptions(c, env.Options{Prefix: "CHURN_"}); err != nil {
		return err
	}

	return nil
}

func (c *Config) ConfigurationPath() string {
	return filepath.Join(c.ConfigPath, c.ConfigSrc)
}

func (c *Config) DatabasePath() string {
	return filepath.Join(c.DataPath, c.DataSrc)
}
