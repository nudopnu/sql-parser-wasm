package parsing

import (
	"strings"

	"github.com/jvictor27/sql-migrate/sqlparse"
	"github.com/nudopnu/sql-parser-wasm/internal/utils"
)

func ParseMigrations(fileContent string) (*sqlparse.ParsedMigration, error) {
	fileContent = utils.TrimWhiteSpaces(fileContent)
	fileContent = strings.ReplaceAll(fileContent, "+goose", "+migrate")
	return sqlparse.ParseMigration(strings.NewReader(fileContent))
}
