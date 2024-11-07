package main

import (
	"fmt"
	"log"
	"syscall/js"

	"github.com/nudopnu/sql-parser-wasm/internal/parsing"
	"github.com/nudopnu/sql-parser-wasm/internal/utils"
)

func parseMigrationsFunc(this js.Value, i []js.Value) interface{} {
	fileContent := i[0].String()
	migration, err := parsing.ParseMigrations(fileContent)
	if err != nil {
		log.Fatal(fmt.Errorf("error parsing migrations: %w", err))
	}
	return utils.ToJSON(migration)
}
