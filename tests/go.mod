module github.com/goflyway/goflyway/tests

go 1.20

require (
	github.com/goflyway/goflyway v0.0.0-20231227095744-97601a534699
	github.com/mattn/go-sqlite3 v1.14.19
)

require github.com/go-sql-driver/mysql v1.7.1

replace github.com/goflyway/goflyway => ../
