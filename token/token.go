package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	Illegal = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	Indent = "IDENT" // add, foobar, x, y, ...
	Int    = "INT"   // 1343456

	// Operators
	Assign = "="
	Plus   = "+"

	// Delimiters
	Comma     = ","
	Semicolon = ";"
	Lparen    = "("
	Rparen    = ")"
	Lbrace    = "{"
	Rbrace    = "}"

	// Keywords
	Function = "FUNCTION"
	Let      = "LET"
)
