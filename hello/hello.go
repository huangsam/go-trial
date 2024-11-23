package hello

import "fmt"

func World() string {
	return "Hello world"
}

func Name(content string) string {
	return fmt.Sprintf("Hello %s\n", content)
}
