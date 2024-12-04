package main

import (
	"antlr4/lab/parser"
	"antlr4/lab/visitor"
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	// Setup the input
	if len(os.Args) < 2 {
		fmt.Println("Missing source file")
		return
	}

	input, err := antlr.NewFileStream(os.Args[1])

	if err != nil {
		fmt.Println("error reading source file: ", err)
		return
	}

	// Create the Lexer
	lexer := parser.NewCalcLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Parser
	p := parser.NewCalcParser(stream)
	tree := p.Program()

	// Visitor
	visitor := visitor.NewCalcVisitor()
	tree.Accept(visitor)

	// fmt.Println(visitor.Res)
	fmt.Println("Builder ---->")

	fmt.Println(visitor.StringBuilder.String())

	os.WriteFile("out.txt", []byte(fmt.Sprintf("%v", visitor.StringBuilder.String())), 0777)

}
