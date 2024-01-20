# GO flyway

![GitHub License](https://img.shields.io/github/license/goflyway/goflyway)
[![Static Badge](https://img.shields.io/badge/go.dev-reference-blue?style=flat)](https://pkg.go.dev/github.com/goflyway/goflyway)

## Install

```shell
go get -u github.com/goflyway/goflyway
```

## Quick start

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

## Support database

- sqlite
- mysql

## Using configuration

Example:

```go
&flyway.Config{...}
```

Configuration item：

 Name              | Type     | Default          | Description                                                                                                                                                                                                                                                                       
-------------------|----------|------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
 Locations         | []string | ["db_migration"] | Comma-separated list of locations to scan recursively for migrations.                                                                                                                                                                                                             
 BaselineOnMigrate | bool     | false            | Whether to automatically call baseline when migrate is executed against a non-empty schema with no schema history table. This schema will then be baselined with the baselineVersion before executing the migrations. Only migrations above baselineVersion will then be applied. 
