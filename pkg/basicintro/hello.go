package basicintro

import "strings"

// GreetWorld returns a greeting for the world.
func GreetWorld() string {
	return GreetName("World")
}

// GreetName returns a greeting for the given name.
func GreetName(content string) string {
	var builder strings.Builder
	builder.WriteString("Hello ")
	builder.WriteString(content)
	return builder.String()
}
