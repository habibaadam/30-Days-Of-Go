package greet

import (
	"strings"
)

func Hello(name string) string { // can be exported
	cleaned := normalizeName(name)

	return "Hello " + cleaned
}

func normalizeName(name string) string {
	n := strings.TrimSpace(name)

	if n == " " {
		return "Guest"
	}
	return strings.ToUpper(n)
}