package lexer_test

import (
	"testing"

	"github.com/mateusmaaia/Lexeme-and-Token-Translator/pkg/lexer"
)

func TestUnknownTokenError(t *testing.T) {
	err := lexer.UnknownTokenError{Literal: "test", Position: lexer.Position{Line: 0, Column: 1}}
	except := "1:2:UnknownTokenError: \"test\""

	if err.Error() != except {
		t.Errorf("excepted %#v but got %s", except, err.Error())
	}
}
