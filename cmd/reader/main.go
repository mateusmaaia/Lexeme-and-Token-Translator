package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mateusmaaia/simple-go-compiler/pkg/lexer"
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
	fileName := "sum.txt"
	check(err)

	file := string(fileByte)

	lexerAnalysis := lexer.NewLexer(strings.NewReader(file))

	var fileContent string
	outputFile, err := os.Create("results/"+fileName)
	defer outputFile.Close()
	check(err)

	fileContent = "==========\n"

	for true {
		token, err := lexerAnalysis.Scan()
		if err != nil {
			panic(err.Error())
		}
		if token == nil {
			fileContent += ("==========")
			break
		}

		fileContent += fmt.Sprintf("line %2d, column %2d: %s: %s\n",
			token.Position.Line,
			token.Position.Column,
			token.Type,
			token.Literal)
	}

	_, err = outputFile.WriteString(fileContent)

}
