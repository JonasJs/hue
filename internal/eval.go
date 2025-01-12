package internal

func Eval(tokens []Token) int {
	parser := NewParser(tokens)
	return parser.evaluateExpression()
}
