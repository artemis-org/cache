package config

import (
	"os"
)

type Config struct {
	RedisUri string
	ListenPoolSize int // Amount of connections to be BLPOP'ing
	PublishPoolSize int // Amount of connections to be HMSET'ing
}

type Provider int

const(
	Toml Provider = iota
	EnvVars Provider = iota
)

var Conf Config

func (p *Provider) LoadConfig() {
	switch *p {
	case Toml:
		loadTomlConfig()
	case EnvVars:
		loadEnvVarConfig()
	}
}

func GetConfigProvider() Provider {
	tomlExists := false
	if _, err := os.Stat("config.toml"); err == nil {
		tomlExists = true
	}

	if tomlExists {
		return Toml
	} else {
		return EnvVars
	}
}
