package main

import (
	"syscall/js"

	"github.com/nudopnu/sql-parser-wasm/internal/parsing"
)

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("parseMigrations", js.FuncOf(parsing.ParseMigrations))
	js.Global().Set("parseSQL", js.FuncOf(parsing.ProcessSql))
	<-c
}
