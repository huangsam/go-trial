package main

import (
	"fmt"

	"sambyte.net/go-trial/hello"
)

func main() {
	fmt.Println(hello.World())
	fmt.Println(hello.Name("Peter"))
}
