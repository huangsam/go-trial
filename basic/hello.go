package basic

import (
	"strings"
)

func World() string {
	return "Hello world"
}

func Name(content string) string {
	var builder strings.Builder
	builder.WriteString("Hello ")
	builder.WriteString(content)
	return builder.String()
}
