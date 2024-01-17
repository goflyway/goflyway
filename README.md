# GO flyway

![GitHub License](https://img.shields.io/github/license/jiangliuhong/go-flyway)
[![Static Badge](https://img.shields.io/badge/go.dev-reference-blue?style=flat)](https://pkg.go.dev/github.com/jiangliuhong/go-flyway)

## 安装

```shell
go get -u github.com/jiangliuhong/go-flyway
```

## 快速开始

```go
package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/jiangliuhong/go-flyway"
	"github.com/jiangliuhong/go-flyway/database"
)

func main() {
	db, _ := sql.Open("sqlite3", "./flyway_test.db")
	// use database.T_SQLITE 、 database.T_MYSQL or "sqlite","mysql"
	f, _ := flyway.Open(database.T_SQLITE, db, &flyway.Config{Locations: []string{"db_migration"}})
	f.Migrate()
}
```

## 支持的数据库

- sqlite
- mysql

## 使用配置

示例:

```go
&flyway.Config{...}
```

配置项：

 名称                | 类型       | 默认值              | 说明                                                                                          
-------------------|----------|------------------|---------------------------------------------------------------------------------------------
 Locations         | []string | ["db_migration"] | 要递归扫描迁移的位置                                                                                  
 BaselineOnMigrate | bool     | false            | 是否在对没有模式历史表的非空模式执行迁移时自动调用基线。在执行迁移之前，将使用baselineVersion对该模式进行基线化。只有baselinversion之上的迁移才会被应用。 
