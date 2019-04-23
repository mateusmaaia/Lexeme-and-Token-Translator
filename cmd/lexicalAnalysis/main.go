package lexicalAnalysis

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/mateusmaaia/Lexeme-and-Token-Translator/pkg/lexer"
)
type (
	tokenMap struct {
		tokens map[string]token
		size int
	}

	token struct {
		name string
		tokenType lexer.TokenType
		positions []position
		tokenMapPosition int
	}

	position struct {
		column int
		line int
	}
)
// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Read(path string) bool {
	x := tokenMap{}
	x.tokens = make(map[string]token)

	fileByte, err := ioutil.ReadFile(path)
	fileName := getFileName(path)

	check(err)

	file := string(fileByte)

	lexerAnalysis := lexer.NewLexer(strings.NewReader(file))

	var fileContent string
	outputFile, err := os.Create("results/"+fileName)
	defer outputFile.Close()
	check(err)


	for true {
		lexerToken, err := lexerAnalysis.Scan()
		if err != nil {
			panic(err.Error())
		}
		if lexerToken == nil {
			break
		}

		if pointerHack, ok := x.tokens[lexerToken.Literal]; ok {
			pointerHack.positions = append(pointerHack.positions, position{
				lexerToken.Position.Column,
				lexerToken.Position.Line,
			})
			x.tokens[lexerToken.Literal] = pointerHack
		} else {
			currentPosition := []position{}
			finalPosition := append(currentPosition, position{
				lexerToken.Position.Column,
				lexerToken.Position.Line,
			})

			x.size++

			x.tokens[lexerToken.Literal] = token{
				lexerToken.Literal,
				lexerToken.Type,
				finalPosition,
				x.size,
			}
		}
	}

	for _, values := range x.tokens {
		fileContent += fmt.Sprintf("Name: %s, Type: %s, Positions(CxL): %v\n",
				values.name,
				values.tokenType,
				values.positions,
			)
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
