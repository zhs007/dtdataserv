package dtdata

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Config - config
type Config struct {
	DTDataServAddr string

	AnkaDB struct {
		DBPath   string
		Engine   string
		HTTPAddr string
	}

	HTTPAddr string
}

// LoadConfig - load config
func LoadConfig(filename string) (*Config, error) {
	fi, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	err = yaml.Unmarshal(fd, cfg)
	if err != nil {
		return nil, err
	}

	err = checkConfig(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// checkConfig -
func checkConfig(cfg *Config) error {
	if cfg.DTDataServAddr == "" {
		return ErrNoServerAddress
	}

	if cfg.AnkaDB.DBPath == "" || cfg.AnkaDB.Engine == "" {
		return ErrNoAnkaDBConfig
	}

	if cfg.HTTPAddr == "" {
		return ErrNoHTTPServerAddr
	}

	return nil
}
