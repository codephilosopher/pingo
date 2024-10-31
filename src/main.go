package main

import (
	"log"
	"os"
	"pingo/src/lexer"
)

func main() {
	bytes, err := os.ReadFile("../examples/test.lang")
	if err != nil {
		log.Println(err)
	}
	tokens := lexer.Tokonize(string(bytes))
	for _, token := range tokens {
		token.Debug()
	}
}
