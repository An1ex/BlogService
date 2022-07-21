package config

import (
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

var SQL MySQL

func Init() error {
	_, err := toml.DecodeFile("config/config.toml", &SQL)
	if err != nil {
		return errors.Wrap(err, "load 'config/config.toml'")
	}
	return nil
}
