package main

import (
	"fmt"
	lexicalAnalysis "github.com/mateusmaaia/simple-go-compiler/pkg"
	"io/ioutil"
	"strings"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	fileByte, err := ioutil.ReadFile("examples/math/sum.txt")
	check(err)

	file := string(fileByte)

	lexer := lexicalAnalysis.NewLexer(strings.NewReader(file))

	fmt.Println(file)
	fmt.Println("==========")

	for {
		token, err := lexer.Scan()
		if err != nil {
			panic(err.Error())
		}
		if token == nil {
			fmt.Println("==========")
			return
		}

		fmt.Printf("line %2d, column %2d: %s: %s\n",
			token.Position.Line,
			token.Position.Column,
			token.Type,
			token.Literal)
	}
}
