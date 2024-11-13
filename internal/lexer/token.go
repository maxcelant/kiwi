package lexer

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var (
	SEMICOLON   TokenType = ";"
	LEFT_PAREN  TokenType = "("
	RIGHT_PAREN TokenType = ")"
	LEFT_BRACE  TokenType = "{"
	RIGHT_BRACE TokenType = "}"
	PLUS        TokenType = "+"
	MINUS       TokenType = "-"
	STAR        TokenType = "*"
	DIV         TokenType = "/"
	SLASH       TokenType = "/"
	BANG        TokenType = "!"
	BANG_EQ     TokenType = "!="
	GREATER     TokenType = ">"
	GREATER_EQ  TokenType = ">="
	LESS        TokenType = "<"
	LESS_EQ     TokenType = "<="
	EQUAL       TokenType = "=="
	IDENTIFIER  TokenType = "IDENTIFIER"
	STRING      TokenType = "STRING"
	NUMBER      TokenType = "NUMBER"
	VAR         TokenType = "var"
	IF          TokenType = "if"
	OR          TokenType = "or"
	AND         TokenType = "and"
	FOR         TokenType = "for"
	WHILE       TokenType = "while"
	CLASS       TokenType = "class"
	FUNC        TokenType = "fn"
	EOF         TokenType = "EOF"
)
