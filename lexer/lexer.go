package lexer

import "github.com/Lentanta/FantasyScript/token"

type Lexer struct {
	input        string
	positon      int  // current positon in input
	readPosition int  // current reading position
	char         byte // current char under examination
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}

	l.positon = l.readPosition
}

func newToken(tkType token.TokenType, char byte) token.Token {
	return token.Token{
		Type:    tkType,
		Literal: string(char),
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.char {
	case '=':
		tok = newToken(token.Assign, l.char)
	case ';':
		tok = newToken(token.Semicolon, l.char)
	case '(':
		tok = newToken(token.Lparen, l.char)
	case ')':
		tok = newToken(token.Rparen, l.char)
	case ',':
		tok = newToken(token.Comma, l.char)
	case '+':
		tok = newToken(token.Plus, l.char)
	case '{':
		tok = newToken(token.Lbrace, l.char)
	case '}':
		tok = newToken(token.Rbrace, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}
