package swagger

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Config entidade da configuração
type Config struct {
	Server   string
	Database string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
