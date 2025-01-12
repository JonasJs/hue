package internal

import (
	"fmt"
	"strconv"
)

// Parser mantém os tokens e a posição atual
type Parser struct {
	tokens    []Token
	position  int
	variables map[string]int
}

// NewParser cria um novo parser com tokens
func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens, position: 0, variables: make(map[string]int)}
}

// Parse analisa os tokens e executa o parsing.
func (p *Parser) Parse() string {
	ast := ""

	// Processa todos os tokens
	for p.position < len(p.tokens) {
		token := p.tokens[p.position]

		// Processa o comando 'let'
		if token.Type == TokenLet {
			ast += "Comando: let\n"
			p.position++
			if p.tokens[p.position].Type == TokenIdentifier {
				varName := p.tokens[p.position].Literal
				p.position++
				if p.tokens[p.position].Type == TokenAssign {
					p.position++
					// Processa a expressão de atribuição
					ast += fmt.Sprintf("Variável %s = %d\n", varName, p.evaluateExpression())
				}
			}
		}

		p.position++
	}

	return ast
}

// evaluateExpression avalia uma expressão simples (números e operadores).
func (p *Parser) evaluateExpression() int {
	result := p.evaluateTerm()

	for p.position < len(p.tokens) && (p.tokens[p.position].Type == TokenPlus || p.tokens[p.position].Type == TokenMinus) {
		op := p.tokens[p.position]
		p.position++

		if op.Type == TokenPlus {
			result += p.evaluateTerm()
		} else if op.Type == TokenMinus {
			result -= p.evaluateTerm()
		}
	}

	return result
}

// evaluateTerm avalia um termo (números e multiplicação/divisão).
func (p *Parser) evaluateTerm() int {
	result := p.evaluateFactor()

	for p.position < len(p.tokens) && (p.tokens[p.position].Type == TokenAsterisk || p.tokens[p.position].Type == TokenSlash) {
		op := p.tokens[p.position]
		p.position++

		if op.Type == TokenAsterisk {
			result *= p.evaluateFactor()
		} else if op.Type == TokenSlash {
			result /= p.evaluateFactor()
		}
	}

	return result
}

func (p *Parser) evaluateFactor() int {
	token := p.tokens[p.position]
	p.position++

	if token.Type == TokenNumber {
		value, _ := strconv.Atoi(token.Literal)
		return value
	} else if token.Type == TokenLParen {
		// Processa sub-expressões entre parênteses
		result := p.evaluateExpression()
		p.position++ // Consome o TokenRParen
		return result
	}

	return 0
}
