package internal

type TokenType string

const (
	TokenIllegal    TokenType = "ILLEGAL"
	TokenEOF        TokenType = "EOF"
	TokenPrint      TokenType = "PRINT"
	TokenLet        TokenType = "LET"
	TokenIdentifier TokenType = "IDENTIFIER"
	TokenNumber     TokenType = "NUMBER"
	TokenPlus       TokenType = "PLUS"
	TokenMinus      TokenType = "MINUS"
	TokenAsterisk   TokenType = "ASTERISK"
	TokenSlash      TokenType = "SLASH"
	TokenAssign     TokenType = "ASSIGN"
	TokenLParen     TokenType = "LPAREN"
	TokenRParen     TokenType = "RPAREN"
	TokenString     TokenType = "STRING"
)

type Token struct {
	Type    TokenType
	Literal string
}
