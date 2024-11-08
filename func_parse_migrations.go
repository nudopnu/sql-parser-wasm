//go:build js && wasm

package main

import (
	"fmt"
	"log"
	"syscall/js"

	"github.com/nudopnu/sql-parser-wasm/internal/parsing"
)

func parseMigrationsFunc(this js.Value, i []js.Value) interface{} {
	fileContent := i[0].String()
	migration, err := parsing.ParseMigrations(fileContent)
	if err != nil {
		log.Fatal(fmt.Errorf("error parsing migrations: %w", err))
	}
	return ToJSON(migration)
}
