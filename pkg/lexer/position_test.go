package lexer_test

import (
	"testing"

	"github.com/mateusmaaia/simple-go-compiler/pkg/lexer"
)

func TestPositionString(t *testing.T) {
	if s := (lexer.Position{Line: 0, Column: 1}).String(); s != "[line:0, column:1]" {
		t.Errorf("failed convert to string: excepted [line:0, column:1] but got %#v", s)
	}

	if s := (lexer.Position{Line: 5, Column: 3}).String(); s != "[line:5, column:3]" {
		t.Errorf("failed convert to string: excepted [line:5, column:3] but got %#v", s)
	}
}

func TestPositionCompare(t *testing.T) {
	a := lexer.Position{Line: 0, Column: 0}
	a2 := lexer.Position{Line: 0, Column: 0}
	b := lexer.Position{Line: 0, Column: 5}
	c := lexer.Position{Line: 1, Column: 3}

	if a != a2 {
		t.Errorf("Position reports %v != %v", a, a2)
	}

	if !a.Before(b) {
		t.Errorf("Position reports %v is not before of %v", a, b)
	}

	if !a.Before(c) {
		t.Errorf("Position reports %v is not before of %v", a, c)
	}

	if !b.Before(c) {
		t.Errorf("Position reports %v is not before of %v", b, c)
	}

	if !b.After(a) {
		t.Errorf("Position reports %v is not after of %v", b, a)
	}

	if !c.After(a) {
		t.Errorf("Position reports %v is not after of %v", c, a)
	}

	if !c.After(b) {
		t.Errorf("Position reports %v is not after of %v", c, b)
	}
}
