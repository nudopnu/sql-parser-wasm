package main

import (
	"fmt"
	"log"
	"syscall/js"

	"github.com/nudopnu/sql-parser-wasm/internal/parsing"
	"github.com/nudopnu/sql-parser-wasm/internal/utils"
)

func processSqlFunc(this js.Value, i []js.Value) interface{} {
	sql := i[0].String()
	statement, err := parsing.ParseSQL(sql)
	if err != nil {
		log.Fatal(fmt.Errorf("error parsing sql: %w", err))
	}
	return utils.ToJSON(statement)
}
