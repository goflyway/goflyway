package flyway

import (
	"database/sql"
	"github.com/jiangliuhong/go-flyway/cmds"
	"github.com/jiangliuhong/go-flyway/consts"
	"github.com/jiangliuhong/go-flyway/database"
	"github.com/jiangliuhong/go-flyway/history"
	_ "github.com/jiangliuhong/go-flyway/init"
	"github.com/jiangliuhong/go-flyway/location"
)

type flyway struct {
	databaseType database.Type
	config       Config
	db           *sql.DB
}

func (f flyway) buildExecuteParam() (d database.Database, h *history.SchemaHistory, o *cmds.Options, err error) {
	d, err = database.New(f.databaseType, f.db)
	if err != nil {
		return
	}
	h, err = history.New(d, f.config.Table)
	var locations []location.Location
	fls := f.config.Locations
	if len(fls) == 0 {
		// set default location
		fls = append(fls, "db_migration")
	}
	for _, item := range fls {
		ls, err2 := location.New(item)
		if err2 != nil {
			err = err2
			return
		}
		locations = append(locations, ls...)
	}
	o = &cmds.Options{
		Locations: locations,
	}
	return
}

func (f flyway) Migrate() error {
	d, h, o, err := f.buildExecuteParam()
	if err != nil {
		return err
	}
	return cmds.Execute(consts.CMD_NAME_MIGRATE, d, h, o)
}

func Open(databaseType database.Type, db *sql.DB, config *Config) (*flyway, error) {
	f := &flyway{
		databaseType: databaseType,
		config:       *config,
		db:           db,
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
