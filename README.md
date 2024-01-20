# GO flyway

![GitHub License](https://img.shields.io/github/license/jiangliuhong/go-flyway)
[![Static Badge](https://img.shields.io/badge/go.dev-reference-blue?style=flat)](https://pkg.go.dev/github.com/goflyway/goflyway)

## 安装

```shell
go get -u github.com/goflyway/goflyway
```

## 快速开始

```go
package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/goflyway/goflyway"
	"github.com/goflyway/goflyway/database"
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
 Schemas           | []string | []               | 数据库连接的模式列表                                                                                  
 CreateSchemas     | bool     | false            | 是否创建 Schemas 指定的模式                                                                          
 DefaultSchema     | string   |                  | 默认的模式，为空时，默认为数据库连接的默认模式，如果指定了 Schemas 则取第一个为默认模式                                            
 CleanDisabled     | bool     | false            | 为ture时，会清空 Schemas 下所有表。注意：生产模式不要设置为true                                                    
 OutOfOrder        | bool     | false            | 是否允许版本乱序运行，为ture时，如果已经应用了1.0和3.0版本，现在发现了2.0版本，那么它也将被应用，而不是被忽略。                              
