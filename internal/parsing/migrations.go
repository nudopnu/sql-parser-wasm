package parsing

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"syscall/js"

	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"github.com/jvictor27/sql-migrate/sqlparse"
)

func preProcess(text string) (result string) {
	for _, line := range strings.Split(text, "\n") {
		result += strings.TrimSpace(line) + "\n"
	}
	return
}

func ParseMigrations(this js.Value, i []js.Value) interface{} {
	text := preProcess(i[0].String())
	fmt.Println(text)
	migration, err := sqlparse.ParseMigration(strings.NewReader(text))
	if err != nil {
		log.Fatal(fmt.Errorf("error parsing migrations: %w", err))
	}
	result, err := json.Marshal(migration)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(result))
	obj := js.Global().Get("JSON").Call("parse", string(result))
	return js.ValueOf(obj)
}

func processSql(sql string) {
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		log.Fatal(fmt.Errorf("error parsing sql statement: %w", err))
	}
	switch stmt := stmt.(type) {
	case *sqlparser.CreateTable:
		fmt.Println(stmt.NewName.Name)
		for _, col := range stmt.Columns {
			fmt.Printf("%s:%s\n", col.Name, col.Type)
		}
	case *sqlparser.DDL:
		fmt.Println(stmt.Action)
	default:
		fmt.Println(stmt)
	}
}
