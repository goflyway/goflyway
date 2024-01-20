package flyway

import (
	"database/sql"
	"github.com/jiangliuhong/go-flyway/command"
	"github.com/jiangliuhong/go-flyway/consts"
	"github.com/jiangliuhong/go-flyway/database"
	_ "github.com/jiangliuhong/go-flyway/init"
)

type flyway struct {
	databaseType database.Type
	config       Config
	db           *sql.DB
}

func (f *flyway) Migrate() error {
	ctx, err := buildCommandCtx(consts.CMD_NAME_MIGRATE, f)
	if err != nil {
		return err
	}
	return command.Execute(ctx)
}

func Open(databaseType string, db *sql.DB, config *Config) (*flyway, error) {
	dbType, err := database.TypeValueOf(databaseType)
	if err != nil {
		return nil, err
	}
	f := &flyway{
		databaseType: dbType,
		config:       *config,
		db:           db,
	}
	return f, nil
}

type Config struct {
	Locations         []string
	Table             string
	BaselineOnMigrate bool     // 是否使用基线迁移
	Schemas           []string // 连接的模式列表
	CreateSchemas     bool     // 是否创建 Schemas 指定的模式
	DefaultSchema     string   // 默认的模式，为空时，默认为数据库连接的默认模式，如果指定了 Schemas 则取第一个为默认模式
	CleanDisabled     bool     // 为ture时，会清空 Schemas 下所有表
	OutOfOrder        bool     // 是否允许版本乱序运行，为ture时，如果已经应用了1.0和3.0版本，现在发现了2.0版本，那么它也将被应用，而不是被忽略。
}
