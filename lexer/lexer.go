package lexer

//import (
//	"fmt"
//)

type Lexer struct {
	text string
	pos int
	ch byte
}

func New(str string) *Lexer {
	l := Lexer{}
	l.text = str
	return &l
}
