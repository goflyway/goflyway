package flyway

import (
	"github.com/jiangliuhong/go-flyway/command"
	"github.com/jiangliuhong/go-flyway/database"
	"github.com/jiangliuhong/go-flyway/history"
	"github.com/jiangliuhong/go-flyway/location"
	"github.com/jiangliuhong/go-flyway/utils"
)

// buildFlyway 构建flyway对象
func buildCommandCtx(commandName string, f *flyway) (*command.Context, error) {
	ctx := &command.Context{
		Command: commandName,
	}
	d, err := database.New(f.databaseType, f.db)
	if err != nil {
		return ctx, err
	}
	ctx.Database = d
	err = schemaHandle(ctx.Database, f.config.Schemas, f.config.CreateSchemas)
	if err != nil {
		return nil, err
	}
	defaultSchema := f.config.DefaultSchema
	if len(f.config.Schemas) > 0 {
		defaultSchema = utils.StringIfNull(defaultSchema, f.config.Schemas[0])
	}
	ctx.SchemaHistory, err = history.New(ctx.Database, history.SchemaHistoryConfig{
		TableName:         f.config.Table,
		BaselineOnMigrate: f.config.BaselineOnMigrate,
		DefaultSchema:     defaultSchema,
	})
	var locations []location.Location
	for _, item := range f.config.Locations {
		ls, err2 := location.New(item)
		if err2 != nil {
			err = err2
			return ctx, err
		}
		locations = append(locations, ls...)
	}
	ctx.Options = &command.Options{
		Locations: locations,
	}
	return ctx, nil
}

// schemaHandle schema列表处理,createSchema为true时，判断schema是否存在，不存在则创建
func schemaHandle(d database.Database, schemas []string, createSchema bool) error {
	if !createSchema {
		return nil
	}
	for _, item := range schemas {
		schema, err := d.Schema(item)
		if err != nil {
			return err
		}
		exists, err := schema.Exists()
		if err != nil {
			return err
		}
		if !exists {
			err = schema.Create()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
