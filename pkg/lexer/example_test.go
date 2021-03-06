package lexer_test

import (
	"fmt"
	"strings"

	"github.com/mateusmaaia/Lexeme-and-Token-Translator/pkg/lexer"
)

func Example() {
	input := "hello_world = \"hello world\"\nnumber = 1"
	lexerAnalysis := lexer.NewLexer(strings.NewReader(input))

	fmt.Println(input)
	fmt.Println("==========")

	for {
		token, err := lexerAnalysis.Scan()
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

	// Output:
	// hello_world = "hello world"
	// number = 1
	// ==========
	// line  0, column  0: IDENT: hello_world
	// line  0, column 12: OTHER: =
	// line  0, column 14: STRING: "hello world"
	// line  1, column  0: IDENT: number
	// line  1, column  7: OTHER: =
	// line  1, column  9: NUMBER: 1
	// ==========
}

func Example_positionInformation() {
	input := "this is a\ntest string\n"
	lexerAnalysis := lexer.NewLexer(strings.NewReader(input))

	for {
		token, err := lexerAnalysis.Scan()
		if err != nil {
			panic(err.Error())
		}
		if token == nil {
			break
		}

		fmt.Printf("%d: %s\n", token.Position.Line, lexerAnalysis.GetLastLine())
		fmt.Printf(" | %s%s\n\n",
			strings.Repeat(" ", token.Position.Column),
			strings.Repeat("=", len(token.Literal)))
	}

	// Output:
	// 0: this is a
	//  | ====
	//
	// 0: this is a
	//  |      ==
	//
	// 0: this is a
	//  |         =
	//
	// 1: test string
	//  | ====
	//
	// 1: test string
	//  |      ======
}

func Example_addOriginalTokenType() {
	const (
		SUBSITUATION lexer.TokenID = iota
		NEWLINE
	)

	input := "hello_world = \"hello world\"\nnumber = 1"
	lexerAnalysis := lexer.NewLexer(strings.NewReader(input))

	lexerAnalysis.Whitespace = lexer.NewPatternTokenType(-1, []string{"\t", " "})
	// lexer.Whitespace = simplexer.NewRegexpTokenType(-1, `[\t ]`)  // same mean above

	lexerAnalysis.TokenTypes = append([]lexer.TokenType{
		lexer.NewPatternTokenType(SUBSITUATION, []string{"="}),
		lexer.NewRegexpTokenType(NEWLINE, `^[\n\r]+`),
	}, lexerAnalysis.TokenTypes...)

	fmt.Println(input)
	fmt.Println("==========")

	for {
		token, err := lexerAnalysis.Scan()
		if err != nil {
			panic(err.Error())
		}
		if token == nil {
			fmt.Println("==========")
			return
		}

		fmt.Printf("%s: %#v\n", token.Type, token.Literal)
	}

	// Output:
	// hello_world = "hello world"
	// number = 1
	// ==========
	// IDENT: "hello_world"
	// UNKNOWN(0): "="
	// STRING: "\"hello world\""
	// UNKNOWN(1): "\n"
	// IDENT: "number"
	// UNKNOWN(0): "="
	// NUMBER: "1"
	// ==========
}
