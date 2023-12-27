package go_flyway

import "database/sql"

type Flyway struct {
	config Config
	db     *sql.DB
}

func (f Flyway) Migrate() error {
	return nil
}

func Open(db *sql.DB, config *Config) (*Flyway, error) {
	f := &Flyway{
		config: *config,
		db:     db,
	}
	return f, nil
}

type Config struct {
	Locations         []string
	Table             string
	BaselineOnMigrate bool
	CleanDisabled     bool
	OutOfOrder        bool
}
