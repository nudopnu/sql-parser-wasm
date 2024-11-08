package utils

import (
	"strings"
)

func TrimWhiteSpaces(text string) (result string) {
	for _, line := range strings.Split(text, "\n") {
		result += strings.TrimSpace(line) + "\n"
	}
	return
}
