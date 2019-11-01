package config

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

type TomlConfig struct {
	Config
}

func loadTomlConfig() {
	raw, err := ioutil.ReadFile("config.toml")
	if err != nil {
		panic(err)
	}

	_, err = toml.Decode(string(raw), &Conf)
	if err != nil {
		panic(err)
	}
}
