package main

import (
	"antlr4/lab/parser"
	"fmt"
	"os"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type calcVisitor struct {
	*parser.BaseCalcVisitor
	variables map[string]float64
	res       float64
}

func NewCalcVisitor() *calcVisitor {
	return &calcVisitor{variables: make(map[string]float64)}
}

func (v *calcVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	fmt.Println(ctx.GetChildCount())
	fmt.Println(ctx.GetText())
	for i := 0; i < ctx.GetChildCount(); i++ {
		if state := ctx.Statement(i); state != nil {
			state.Accept(v)
		}
	}

	return 0
}

func (v *calcVisitor) VisitStatement(ctx *parser.StatementContext) any {
	fmt.Println("Statemant scope")

	fmt.Println(ctx.GetText())

	if ctx.Assigment() != nil {
		fmt.Println("if scope for assigment")

		return ctx.Assigment().Accept(v)
	}
	fmt.Println("if scope for expression")

	v.res = ctx.Expression().Accept(v).(float64)

	return 0
}

func (v *calcVisitor) VisitAssigment(ctx *parser.AssigmentContext) any {
	fmt.Println("assigment scope")

	id := ctx.ID().GetText()
	value := ctx.Expression().Accept(v).(float64)
	v.variables[id] = value

	fmt.Printf("write to mem with id = %v and val = %v \n", id, value)

	return value

}

func (v *calcVisitor) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
	left := ctx.Expression(0).Accept(v).(float64)
	right := ctx.Expression(1).Accept(v).(float64)

	switch ctx.GetOp().GetTokenType() {

	case parser.CalcLexerMUL:
		fmt.Printf("Mul with l = %v and r = %v\n", left, right)
		return left * right

	case parser.CalcLexerDIV:
		if right == 0 {
			panic("division by zero!")
		}
		fmt.Printf("Div with l = %v and r = %v\n", left, right)

		return left / right
	}

	return 0

}

func (v *calcVisitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
	left := ctx.Expression(0).Accept(v).(float64)
	right := ctx.Expression(1).Accept(v).(float64)

	switch ctx.GetOp().GetTokenType() {

	case parser.CalcLexerADD:
		fmt.Printf("Add with l = %v and r = %v\n", left, right)

		return left + right

	case parser.CalcLexerSUB:
		fmt.Printf("Sub` with l = %v and r = %v\n", left, right)

		return left - right
	}

	return 0

}

func (v *calcVisitor) VisitNumber(ctx *parser.NumberContext) interface{} {
	val, _ := strconv.ParseFloat(ctx.GetText(), 64)

	return val
}

func (v *calcVisitor) VisitID(ctx *parser.IDContext) any {
	id := ctx.GetText()

	if value, exists := v.variables[id]; exists {
		return value
	}

	panic("no such var in memory!")
}

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
	visitor := NewCalcVisitor()
	tree.Accept(visitor)

	fmt.Println(visitor.res)

	os.WriteFile("out.txt", []byte(fmt.Sprintf("%v", visitor.res)), 0777)

}
