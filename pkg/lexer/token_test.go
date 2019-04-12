package lexer_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mateusmaaia/Lexeme-and-Token-Translator/pkg/lexer"
)

func ExampleNewRegexpTokenType() {
	const (
		NUMBER lexer.TokenID = iota
		OTHERS
	)

	lexerAnalysis := lexer.NewLexer(strings.NewReader("123this is test456"))

	lexerAnalysis.TokenTypes = []lexer.TokenType{
		lexer.NewRegexpTokenType(NUMBER, `[0-9]+`),
		lexer.NewRegexpTokenType(OTHERS, `[^0-9]+`),
	}

	for {
		token, _ := lexerAnalysis.Scan()
		if token == nil {
			break
		}

		if token.Type.GetID() == NUMBER {
			fmt.Printf("%s is number\n", token.Literal)
		}

		if token.Type.GetID() == OTHERS {
			fmt.Printf("%s is not number\n", token.Literal)
		}
	}

	// Output:
	// 123 is number
	// this is test is not number
	// 456 is number
}

func ExampleNewPatternTokenType() {
	const (
		HOGE lexer.TokenID = iota
		OTHERS
	)

	lexerAnalysis := lexer.NewLexer(strings.NewReader("this is hoge and HOGE or Hoge"))

	lexerAnalysis.TokenTypes = []lexer.TokenType{
		lexer.NewPatternTokenType(HOGE, []string{"hoge", "HOGE"}),
		lexer.NewRegexpTokenType(OTHERS, `[^ ]+`),
	}

	for {
		token, _ := lexerAnalysis.Scan()
		if token == nil {
			break
		}

		if token.Type.GetID() == HOGE {
			fmt.Printf("!!! %s !!!\n", token.Literal)
		}

		if token.Type.GetID() == OTHERS {
			fmt.Println(token.Literal)
		}
	}

	// Output:
	// this
	// is
	// !!! hoge !!!
	// and
	// !!! HOGE !!!
	// or
	// Hoge
}

func TestRegexpTokenType(t *testing.T) {
	tt := lexer.NewRegexpTokenType(1, `[0-9]+(\.[0-9]+)?`)

	if tok := tt.FindToken("not match 123", lexer.Position{}); tok != nil {
		t.Errorf("excepted nil but got %#v", tok)
	}

	pos := lexer.Position{Line: 1, Column: 2}

	if tok := tt.FindToken("123.1abc", pos); tok == nil {
		t.Errorf("excepted token but got nil")
	} else {
		if tok.Type != tt {
			t.Errorf("excepted token type is %#v but got %#v", &tt, &tok.Type)
		}
		if tok.Type.GetID() != 1 {
			t.Errorf("excepted token type ID is 1 but got %#v", tok.Type.GetID())
		}
		if tok.Literal != "123.1" {
			t.Errorf("excepted literal of token is 123.1 but got %#v", tok.Literal)
		}
		if len(tok.Submatches) != 1 || tok.Submatches[0] != ".1" {
			t.Errorf("excepted submatches of token is %#v but got %#v", []string{".1"}, tok.Submatches)
		}
		if tok.Position != pos {
			t.Errorf("excepted position of token is %#v but got %#v", pos, tok.Position)
		}
	}
}

func TestPatternTokenType(t *testing.T) {
	tt := lexer.NewPatternTokenType(1, []string{"abc", "def"})

	if tok := tt.FindToken("not match abc", lexer.Position{}); tok != nil {
		t.Errorf("excepted nil but got %#v", tok)
	}

	pos := lexer.Position{Line: 1, Column: 2}

	if tok := tt.FindToken("abc def", pos); tok == nil {
		t.Errorf("excepted token but got nil")
	} else {
		if tok.Type != tt {
			t.Errorf("excepted token type is %#v but got %#v", &tt, &tok.Type)
		}
		if tok.Type.GetID() != 1 {
			t.Errorf("excepted token type ID is 1 but got %#v", tok.Type.GetID())
		}
		if tok.Literal != "abc" {
			t.Errorf("excepted literal of token is abc but got %#v", tok.Literal)
		}
		if len(tok.Submatches) != 0 {
			t.Errorf("excepted submatches of token is empty but got %#v", tok.Submatches)
		}
		if tok.Position != pos {
			t.Errorf("excepted position of token is %#v but got %#v", pos, tok.Position)
		}
	}

	if tok := tt.FindToken("def", pos); tok == nil {
		t.Errorf("excepted token but got nil")
	} else {
		if tok.Type != tt {
			t.Errorf("excepted token type is %#v but got %#v", &tt, &tok.Type)
		}
		if tok.Type.GetID() != 1 {
			t.Errorf("excepted token type ID is 1 but got %#v", tok.Type.GetID())
		}
		if tok.Literal != "def" {
			t.Errorf("excepted literal of token is def but got %#v", tok.Literal)
		}
		if len(tok.Submatches) != 0 {
			t.Errorf("excepted submatches of token is empty but got %#v", tok.Submatches)
		}
		if tok.Position != pos {
			t.Errorf("excepted position of token is %#v but got %#v", pos, tok.Position)
		}
	}
}
