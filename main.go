//go:build js && wasm

package main

import (
	"encoding/json"
	"log"
	"syscall/js"
)

func ToJSON(v any) interface{} {
	result, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	obj := js.Global().Get("JSON").Call("parse", string(result))
	return js.ValueOf(obj)
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("parseMigrations", js.FuncOf(parseMigrationsFunc))
	js.Global().Set("parseSQL", js.FuncOf(processSqlFunc))
	<-c
}
