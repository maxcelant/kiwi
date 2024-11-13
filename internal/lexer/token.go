package lexer

type TokenType int

type Token struct {
	Type    TokenType
	Literal interface{}
	Lexeme  string
	Line    int64
}

const (
	SEMICOLON TokenType = iota
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	PLUS
	MINUS
	STAR
	DIV
	SLASH
	BANG
	BANG_EQ
	GREATER
	GREATER_EQ
	LESS
	LESS_EQ
	EQUAL
	EQUAL_EQUAL
	IDENTIFIER
	STRING
	NUMBER
	VAR
	IF
	OR
	AND
	FOR
	WHILE
	CLASS
	FUNC
	EOF
)
