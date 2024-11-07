package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"github.com/jvictor27/sql-migrate/sqlparse"
)

func main() {
	migration, err := sqlparse.ParseMigration(strings.NewReader(`-- +migrate Up
CREATE TABLE refresh_tokens (
token TEXT PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	expires_at TIMESTAMP NOT NULL,
	revoked_at TIMESTAMP,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +migrate Down
DROP TABLE refresh_tokens;`))
	if err != nil {
		log.Fatal(err)
	}
	for _, sql := range migration.UpStatements {
		processStatement(sql)
	}
	for _, sql := range migration.DownStatements {
		processStatement(sql)
	}
}

func processStatement(sql string) {
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		panic(err)
	}
	switch stmt := stmt.(type) {
	case *sqlparser.CreateTable:
		fmt.Println(stmt.NewName.Name)
		for _, col := range stmt.Columns {
			fmt.Printf("%s:%s\n", col.Name, col.Type)
		}
	case *sqlparser.Delete:
		fmt.Println(stmt)
	default:
		fmt.Println(stmt)
	}
}
