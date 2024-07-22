package lexer

import (
	"github.com/Lentanta/FantasyScript/token"
)

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
	l.readPosition += 1
}

func (l *Lexer) peek() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func newToken(tkType token.TokenType, char byte) token.Token {
	return token.Token{
		Type:    tkType,
		Literal: string(char),
	}
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' ||
		'A' <= char && char <= 'Z' ||
		char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (lex *Lexer) skipWhitespace() {
	for lex.char == ' ' ||
		lex.char == '\t' ||
		lex.char == '\n' ||
		lex.char == '\r' {
		lex.readChar()
	}
}

func (l *Lexer) readIndentifier() string {
	positon := l.positon
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[positon:l.positon]
}

func (l *Lexer) readNumber() string {
	positon := l.positon
	for isDigit(l.char) {
		l.readChar()
	}
	return l.input[positon:l.positon]
}

func New(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar()
	return lex
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.char {
	case '=':
		if l.peek() == '=' {
			char := l.char
			l.readChar()
			tok = token.Token{
				Type:    token.EQ,
				Literal: string(char) + string(l.char),
			}
		} else {
			tok = newToken(token.ASSIGN, l.char)
		}
	case '!':
		if l.peek() == '=' {
			char := l.char
			l.readChar()
			tok = token.Token{
				Type:    token.NOT_EQ,
				Literal: string(char) + string(l.char),
			}
		} else {
			tok = newToken(token.BANG, l.char)
		}
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '-':
		tok = newToken(token.MINUS, l.char)
	case '/':
		tok = newToken(token.SLASH, l.char)
	case '*':
		tok = newToken(token.ASTERISK, l.char)
	case '<':
		tok = newToken(token.LESS_THAN, l.char)
	case '>':
		tok = newToken(token.GREATER_THAN, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.char) {
			return token.Token{
				Literal: l.readIndentifier(),
				Type:    token.LookupIndent(tok.Literal),
			}
		} else if isDigit(l.char) {
			return token.Token{
				Literal: l.readNumber(),
				Type:    token.LookupIndent(tok.Literal),
			}
		} else {
			tok = newToken(token.ILLEGAL, l.char)
		}
	}

	l.readChar()
	return tok
}
