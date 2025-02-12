package config

import (
	"github.com/pelletier/go-toml"
	"log"
)

type Config struct {
	Mysql Mysql
	Etcd  Etcd
}

func (c *Config) Init() {
	tree, err := toml.LoadFile("config.toml")
	if err != nil {
		log.Fatalf("Error loading TOML file: %v", err)
	}
	err = tree.Unmarshal(c)
	if err != nil {
		log.Fatalf("Error unmarshalling TOML into struct: %v", err)
	}
}
