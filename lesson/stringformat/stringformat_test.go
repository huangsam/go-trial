package stringformat

import (
	"fmt"
	"strings"
	"testing"
)

func TestFmtSprintf(t *testing.T) {
	result := FmtSprintf("Alice", 30)
	expected := "Name: Alice, Age: 30"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestStringsBuilder(t *testing.T) {
	words := []string{"apple", "banana", "cherry"}
	result := StringsBuilder(words)
	expected := "apple, banana, cherry"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestBytesBuffer(t *testing.T) {
	words := []string{"apple", "banana", "cherry"}
	result := BytesBuffer(words)
	expected := "apple, banana, cherry"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestFmtPrint(t *testing.T) {
	result := FmtPrint("Bob", 25)
	if !strings.Contains(result, "Name:") || !strings.Contains(result, "Age: 25") {
		t.Errorf("Unexpected result: %s", result)
	}
}

func TestStringConcat(t *testing.T) {
	words := []string{"apple", "banana", "cherry"}
	result := StringConcat(words)
	expected := "apple, banana, cherry"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestJoin(t *testing.T) {
	words := []string{"apple", "banana", "cherry"}
	result := Join(words)
	expected := "apple, banana, cherry"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestPercentFormatting(t *testing.T) {
	result := PercentFormatting("%s is %d years old", "Alice", 30)
	expected := "Alice is 30 years old"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestMultilineStrings(t *testing.T) {
	result := MultilineStrings()
	if !strings.Contains(result, "\n") {
		t.Error("Expected multiline string with newlines")
	}
	if !strings.Contains(result, "multiline string") {
		t.Errorf("Unexpected content: %s", result)
	}
}

func BenchmarkFmtSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FmtSprintf("Alice", 30)
	}
}

func BenchmarkStringsBuilder(b *testing.B) {
	words := []string{"apple", "banana", "cherry"}
	for i := 0; i < b.N; i++ {
		StringsBuilder(words)
	}
}

func BenchmarkBytesBuffer(b *testing.B) {
	words := []string{"apple", "banana", "cherry"}
	for i := 0; i < b.N; i++ {
		BytesBuffer(words)
	}
}

func BenchmarkStringConcat(b *testing.B) {
	words := make([]string, 100)
	for i := range words {
		words[i] = fmt.Sprintf("word%d", i)
	}
	for i := 0; i < b.N; i++ {
		StringConcat(words)
	}
}

func BenchmarkJoin(b *testing.B) {
	words := make([]string, 100)
	for i := range words {
		words[i] = fmt.Sprintf("word%d", i)
	}
	for i := 0; i < b.N; i++ {
		Join(words)
	}
}

func TestBuilderLength(t *testing.T) {
	var builder strings.Builder
	builder.WriteString("Hello")
	if builder.Len() != 5 {
		t.Errorf("Expected length 5, got %d", builder.Len())
	}
}

func TestBuilderReset(t *testing.T) {
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.Reset()
	if builder.Len() != 0 {
		t.Errorf("Expected length 0 after reset, got %d", builder.Len())
	}
}

func TestJoinEmptySlice(t *testing.T) {
	result := Join([]string{})
	if result != "" {
		t.Errorf("Expected empty string for empty slice, got '%s'", result)
	}
}
