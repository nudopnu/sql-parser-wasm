package parsing

import (
	"github.com/nudopnu/sql-parser-wasm/internal/utils"
	"vitess.io/vitess/go/vt/sqlparser"
)

type Statement struct {
	sqlparser.Statement
	Type string
}

func ParseSQL(sql string) (Statement, error) {
	sql = utils.TrimWhiteSpaces(sql)
	parser, err := sqlparser.New(sqlparser.Options{})
	if err != nil {
		return Statement{}, nil
	}
	stmt, err := parser.Parse(sql)

	typus := ""
	switch stmt := stmt.(type) {
	case sqlparser.DDLStatement:
		typus = stmt.GetAction().ToString()
	}
	return Statement{Statement: stmt, Type: typus}, err
}
