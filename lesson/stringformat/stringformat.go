package stringformat

import (
	"bytes"
	"fmt"
	"strings"
)

// FmtSprintf demonstrates using fmt.Sprintf for string formatting.
func FmtSprintf(name string, age int) string {
	return fmt.Sprintf("Name: %s, Age: %d", name, age)
}

// StringsBuilder demonstrates using strings.Builder for efficient concatenation.
func StringsBuilder(words []string) string {
	var builder strings.Builder
	for _, word := range words {
		if builder.Len() > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(word)
	}
	return builder.String()
}

// BytesBuffer demonstrates using bytes.Buffer for efficient concatenation.
func BytesBuffer(words []string) string {
	var buf bytes.Buffer
	for i, word := range words {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(word)
	}
	return buf.String()
}

// FmtPrint demonstrates using fmt.Print and fmt.Println.
func FmtPrint(name string, age int) string {
	var builder strings.Builder
	fmt.Fprint(&builder, "Name: ", name)
	fmt.Fprintf(&builder, ", Age: %d", age)
	return builder.String()
}

// StringConcat demonstrates simple string concatenation with + operator.
func StringConcat(words []string) string {
	result := ""
	for i, word := range words {
		if i > 0 {
			result += ", "
		}
		result += word
	}
	return result
}

// Join demonstrates using strings.Join for joining slices.
func Join(words []string) string {
	return strings.Join(words, ", ")
}

// PercentFormatting demonstrates C-style % formatting.
func PercentFormatting(format string, a interface{}, b interface{}) string {
	return fmt.Sprintf(format, a, b)
}

// MultilineStrings demonstrates multiline string literals.
func MultilineStrings() string {
	return `This is a
multiline string
in Go using backticks.`
}
