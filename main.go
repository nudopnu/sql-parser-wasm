package main

import (
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("parseMigrations", js.FuncOf(parseMigrationsFunc))
	js.Global().Set("parseSQL", js.FuncOf(processSqlFunc))
	<-c
}
