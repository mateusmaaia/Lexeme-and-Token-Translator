package lexer_test

import (
	"testing"
)

func TestUnknownTokenError(t *testing.T) {
	err := UnknownTokenError{Literal: "test", Position: Position{Line: 0, Column: 1}}
	except := "1:2:UnknownTokenError: \"test\""

	if err.Error() != except {
		t.Errorf("excepted %#v but got %s", except, err.Error())
	}
}
