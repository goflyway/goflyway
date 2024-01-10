package cmds

import (
	"errors"
	"fmt"
	"github.com/jiangliuhong/go-flyway/consts"
	"github.com/jiangliuhong/go-flyway/database"
	"github.com/jiangliuhong/go-flyway/history"
	"github.com/jiangliuhong/go-flyway/location"
	"time"
)

func init() {
	Registry(consts.CMD_NAME_MIGRATE, &Migrate{})
}

type Migrate struct {
}

func (m Migrate) Execute(database database.Database, schemaHistory *history.SchemaHistory, options *Options) error {
	exists, err := schemaHistory.Exists()
	if err != nil {
		return err
	}
	if !exists {
		err = schemaHistory.Create()
		if err != nil {
			return err
		}
	}
	err = schemaHistory.InitBaseLineRank()
	if err != nil {
		return err
	}
	for _, l := range options.Locations {
		for _, sql := range l.Sqls {
			err = m.invokeSql(database, schemaHistory, sql)
			if err != nil {
				return errors.New(fmt.Sprintf("Failed to execute the SQL file:%s\nerror:%s", sql.Path, err.Error()))
			}
		}
	}
	return nil
}

func (m Migrate) invokeSql(database database.Database, schemaHistory *history.SchemaHistory, sql location.SqlFile) error {
	sd, err := schemaHistory.SelectVersion(sql.Version)
	if err != nil {
		return err
	}
	checksum, err := sql.CheckSum()
	var rank int64
	if sd != nil {
		rank = sd.InstalledRank
		if checksum != sd.Checksum {
			return errors.New(fmt.Sprintf("Flyway checksum mismatch error\n database: %d,local:%d", sd.Checksum, checksum))
		}
		if !sd.Success {
			content, err2 := sql.Content()
			if err2 != nil {
				return err2
			}
			d, err2 := m.invokeSqlContent(database, content)
			if err2 != nil {
				return err2
			} else {
				err = schemaHistory.UpdateSuccessAndTime(rank, true, d.Microseconds())
				if err != nil {
					return err
				}
			}
		}
	} else {
		content, err2 := sql.Content()
		if err2 != nil {
			return err2
		}
		sd = &history.SchemaData{
			Version:       sql.Version,
			Description:   sql.Description,
			Type:          consts.SQL_TYPE,
			Script:        content,
			ExecutionTime: 0,
			Success:       false,
			Checksum:      checksum,
		}
		rank, err = schemaHistory.InsertData(*sd)
		if err != nil {
			return err
		}
		d, err2 := m.invokeSqlContent(database, content)
		if err2 != nil {
			return err2
		} else {
			err = schemaHistory.UpdateSuccessAndTime(rank, true, d.Microseconds())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (m Migrate) invokeSqlContent(database database.Database, content string) (time.Duration, error) {
	start := time.Now()
	err := database.Session().Exec(content)
	since := time.Since(start)
	return since, err
}
