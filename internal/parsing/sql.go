package parsing

import (
	"fmt"
	"syscall/js"

	"github.com/blastrain/vitess-sqlparser/sqlparser"
)

type Table struct {
	Name    string   `json:"name"`
	Columns []Column `json:"columns"`
}

type Column struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func ProcessSql(this js.Value, i []js.Value) interface{} {
	sql := preProcess(i[0].String())
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
	case *sqlparser.DDL:
		fmt.Println(stmt.Action)
	default:
		fmt.Println(stmt)
	}
	return js.ValueOf(":)")
}
