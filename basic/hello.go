package basic

import (
	"strings"
)

func GreetWorld() string {
	return GreetName("World")
}

func GreetName(content string) string {
	var builder strings.Builder
	builder.WriteString("Hello ")
	builder.WriteString(content)
	return builder.String()
}
