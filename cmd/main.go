package main

import (
    "fmt"
    "github.com/JonasJs/hue/internal"
    "io/ioutil"
    "os"
)

func main() {
    if len(os.Args) > 1 {
        file := os.Args[1]
        if err := runFile(file); err != nil {
            fmt.Printf("Error executing the file %s: %v\n", file, err)
        }
    } else {
        runRepl()
    }
}

func runFile(filename string) error {
    if ext := getFileExtension(filename); ext != ".hue" {
        return fmt.Errorf("invalid extension: %s (expected .hue)", ext)
    }

    content, err := ioutil.ReadFile(filename)
    if err != nil {
        return err
    }

    tokens := internal.Lex(string(content))
    internal.Parse(tokens)
    return nil
}

func getFileExtension(filename string) string {
    dotIndex := len(filename) - len(".hue")
    if dotIndex >= 0 && filename[dotIndex:] == ".hue" {
        return ".hue"
    }
    return ""
}

func runRepl() {
    fmt.Println("Welcome to the hue language REPL!")
    // Implementação anterior do REPL aqui.
}
