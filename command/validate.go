package command

import (
	"errors"
	"fmt"
	"github.com/goflyway/goflyway/consts"
	"strings"
)

func init() {
	Registry(consts.CMD_NAME_MIGRATE, &Validate{})
}

type Validate struct {
}

func (v Validate) Execute(ctx *Context) error {
	_, err := beforeExecute(ctx)
	if err != nil {
		return err
	}
	var errs []error
	for _, l := range ctx.Options.Locations {
		for _, sql := range l.Sqls {
			schemaHistory := ctx.SchemaHistory
			sd, err := schemaHistory.SelectVersion(sql.Version)
			if err != nil {
				return err
			}
			checksum, err := sql.CheckSum()
			if sd != nil && checksum != sd.Checksum {
				errs = append(errs, errors.New(fmt.Sprintf("Flyway checksum mismatch error\n database: %d,local:%d", sd.Checksum, checksum)))
			}
		}
	}
	if len(errs) > 0 {
		var errMsg []string
		for _, e := range errs {
			errMsg = append(errMsg, e.Error())
		}
		ctx.Logger.Error(ctx, strings.Join(errMsg, "\n"))
		return errors.New("validate failed")
	}
	return nil
}
