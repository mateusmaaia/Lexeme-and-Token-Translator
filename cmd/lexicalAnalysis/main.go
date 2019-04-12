package lexicalAnalysis

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/mateusmaaia/Lexeme-and-Token-Translator/pkg/lexer"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Read(path string) bool {

	fileByte, err := ioutil.ReadFile(path)
	fileName := getFileName(path)

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
	check(err)

	return true
}

func getFileName(path string) string {
	r,_ := regexp.Compile("[^/]+$")
	fileName := r.FindString(path)
	return fileName
}
