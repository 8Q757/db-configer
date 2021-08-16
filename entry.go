package configer

import (
	"github.com/pelletier/go-toml"
)

func GetDSN() string {
	config, err := toml.LoadFile("db.toml")
	if err != nil {
		panic("Unable to read configuration file: db.toml.")
	}

	keys := config.Keys()
	if len(keys) == 0 {
		panic("Configuration file error.")
	}

	for _, key := range keys {
		switch key {
		case PostgreSQL:
			//TODO Can toml be directly converted into a struct?
			return PostgresMapToDsn(config.Get(PostgreSQL).(*toml.Tree).ToMap())

		default:
			panic("This database is not currently supported.")
		}
	}
	return ""
}
