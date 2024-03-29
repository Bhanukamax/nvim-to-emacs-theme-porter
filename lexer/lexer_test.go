package lexer

import (
	"testing"
)


func TestCreatingALexer(t* testing.T) {
	var test = struct {
		text string
		pos int
		ch byte
	} {
		"foo", 0, 0,
	}

	l := New("foo")

	if l.text != test.text {
		t.Errorf("text does not match")
	}
}
