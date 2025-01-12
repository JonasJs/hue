package internal

import (
	"strings"
)

type TokenType string

const (
	TokenPrint  TokenType = "PRINT"
	TokenString TokenType = "STRING"
	TokenLParen TokenType = "LPAREN" // (
	TokenRParen TokenType = "RPAREN" // )
)

type Token struct {
	Type    TokenType
	Literal string
}

func Lex(input string) []Token {
	var tokens []Token
	input = strings.TrimSpace(input)

	if strings.HasPrefix(input, "print(") && strings.HasSuffix(input, ")") {
		tokens = append(tokens, Token{Type: TokenPrint, Literal: "print"})

		content := strings.TrimSuffix(strings.TrimPrefix(input, "print("), ")")
		content = strings.Trim(content, "\"")

		tokens = append(tokens, Token{Type: TokenLParen, Literal: "("})
		tokens = append(tokens, Token{Type: TokenString, Literal: content})
		tokens = append(tokens, Token{Type: TokenRParen, Literal: ")"})
	}

	return tokens
}
