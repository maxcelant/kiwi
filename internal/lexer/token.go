package lexer

type TokenType int

type Token struct {
	Type    TokenType
	Literal any
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
	TRUE
	FALSE
	NIL
	VAR
	IF
	ELSE
	OR
	AND
	FOR
	PRINT
	WHILE
	RETURN
	CLASS
	FUNC
	EOF
)
