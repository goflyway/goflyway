package flyway

import (
	"github.com/goflyway/goflyway/command"
	"github.com/goflyway/goflyway/database"
	"github.com/goflyway/goflyway/history"
	"github.com/goflyway/goflyway/location"
	"github.com/goflyway/goflyway/utils"
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
	err = schemaCleanHandle(ctx.Database, f.config.Schemas, f.config.CleanDisabled)
	if err != nil {
		return nil, err
	}
	defaultSchema := f.config.DefaultSchema
	if len(f.config.Schemas) > 0 {
		defaultSchema = utils.StringIfNull(defaultSchema, f.config.Schemas[0])
	}
	s, err := history.New(ctx.Database, history.SchemaHistoryConfig{
		TableName:         f.config.Table,
		BaselineOnMigrate: f.config.BaselineOnMigrate,
		DefaultSchema:     defaultSchema,
	})
	if err != nil {
		return ctx, err
	}
	ctx.SchemaHistory = s
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
		Locations:  locations,
		OutOfOrder: f.config.OutOfOrder,
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

// schemaCleanHandle 清空 schema ，先删除后新建
func schemaCleanHandle(d database.Database, schemas []string, cleanDisabled bool) error {
	if !cleanDisabled {
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
		// 如果存在，先删除再新建
		if exists {
			err = schema.Delete()
			if err != nil {
				return err
			}
			err = schema.Create()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
