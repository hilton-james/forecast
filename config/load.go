package config

import (
	"log"

	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
)

func New() Config {
	k := koanf.New(".")

	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("unable to load the default config: %v\n", err)
	}

	if err := k.Load(file.Provider("config.toml"), toml.Parser()); err != nil {
		log.Fatalf("unable to load the config.toml: %v\n", err)
	}

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		log.Fatalf("unable to unmarshal config: %v\n", err.Error())
	}

	return cfg
}
