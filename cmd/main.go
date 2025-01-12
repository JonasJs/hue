package main

import (
	"fmt"
	"hue/internal"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		file := os.Args[1]
		if err := runFile(file); err != nil {
			fmt.Printf("Erro ao executar o arquivo %s: %v\n", file, err)
		}
	} else {
		runRepl()
	}
}

// runFile executa o arquivo e faz o lexing e parsing.
func runFile(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	tokens := internal.Lex(string(content))
	parser := internal.NewParser(tokens)
	ast := parser.Parse()
	fmt.Println("AST gerado: \n", ast)
	return nil
}

func runRepl() {
	fmt.Println("Bem-vindo ao REPL da linguagem hue!")
	for {
		fmt.Print("hue> ")
		var line string
		_, err := fmt.Scanln(&line)
		if err != nil {
			fmt.Println("Erro ao ler entrada:", err)
			continue
		}

		tokens := internal.Lex(line)
		parser := internal.NewParser(tokens)
		ast := parser.Parse()
		fmt.Println("AST gerado: \n", ast)
	}
}
