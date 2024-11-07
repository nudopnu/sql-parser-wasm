package parsing

import (
	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"github.com/nudopnu/sql-parser-wasm/internal/utils"
)

func ParseSQL(sql string) (sqlparser.Statement, error) {
	sql = utils.TrimWhiteSpaces(sql)
	return sqlparser.Parse(sql)
}
