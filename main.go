package main

import (
	"fmt"

	"l-code/parser"
)

func main() {
	fmt.Println(parser.Parser.String())
	grammar, err := parser.ParseString(`study 'name'`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*grammar)
}
