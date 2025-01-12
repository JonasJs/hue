package internal

import (
	"strings"
	"unicode"
)

// Lex é responsável por tokenizar a entrada.
func Lex(input string) []Token {
	var tokens []Token
	input = strings.TrimSpace(input)

	for i := 0; i < len(input); i++ {
		ch := input[i]

		switch {
		case unicode.IsSpace(rune(ch)):
			// Ignora espaços
			continue
		case ch == '+':
			tokens = append(tokens, Token{Type: TokenPlus, Literal: string(ch)})
		case ch == '-':
			tokens = append(tokens, Token{Type: TokenMinus, Literal: string(ch)})
		case ch == '*':
			tokens = append(tokens, Token{Type: TokenAsterisk, Literal: string(ch)})
		case ch == '/':
			tokens = append(tokens, Token{Type: TokenSlash, Literal: string(ch)})
		case ch == '=':
			tokens = append(tokens, Token{Type: TokenAssign, Literal: string(ch)})
		case ch == '(':
			tokens = append(tokens, Token{Type: TokenLParen, Literal: string(ch)})
		case ch == ')':
			tokens = append(tokens, Token{Type: TokenRParen, Literal: string(ch)})
		case ch == 'l' && strings.HasPrefix(input[i:], "let"):
			tokens = append(tokens, Token{Type: TokenLet, Literal: "let"})
			i += 2 // Pula a palavra "let"
		case ch == 'p' && strings.HasPrefix(input[i:], "print"):
			tokens = append(tokens, Token{Type: TokenPrint, Literal: "print"})
			i += 4 // Pula a palavra "print"
		case unicode.IsDigit(rune(ch)):
			number := string(ch)
			for i+1 < len(input) && unicode.IsDigit(rune(input[i+1])) {
				i++
				number += string(input[i])
			}
			tokens = append(tokens, Token{Type: TokenNumber, Literal: number})
		case unicode.IsLetter(rune(ch)):
			identifier := string(ch)
			for i+1 < len(input) && (unicode.IsLetter(rune(input[i+1])) || unicode.IsDigit(rune(input[i+1]))) {
				i++
				identifier += string(input[i])
			}
			tokens = append(tokens, Token{Type: TokenIdentifier, Literal: identifier})
		default:
			tokens = append(tokens, Token{Type: TokenIllegal, Literal: string(ch)})
		}
	}

	tokens = append(tokens, Token{Type: TokenEOF, Literal: ""}) // Marca o fim
	return tokens
}
