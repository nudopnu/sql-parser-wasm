package utils

import (
	"encoding/json"
	"log"
	"strings"
	"syscall/js"
)

func TrimWhiteSpaces(text string) (result string) {
	for _, line := range strings.Split(text, "\n") {
		result += strings.TrimSpace(line) + "\n"
	}
	return
}

func ToJSON(v any) interface{} {
	result, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	obj := js.Global().Get("JSON").Call("parse", string(result))
	return js.ValueOf(obj)
}
