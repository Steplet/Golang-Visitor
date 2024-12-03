package visitor

import (
	"antlr4/lab/parser"
	"fmt"
	"strconv"
)

type CalcVisitor struct {
	*parser.BaseCalcVisitor
	Variables map[string]float64
	Res       float64
}

func NewCalcVisitor() *CalcVisitor {
	return &CalcVisitor{Variables: make(map[string]float64)}
}

func (v *CalcVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	fmt.Println(ctx.GetChildCount())
	fmt.Println(ctx.GetText())
	for i := 0; i < ctx.GetChildCount(); i++ {
		if state := ctx.Statement(i); state != nil {
			state.Accept(v)
		}
	}

	return 0
}

func (v *CalcVisitor) VisitStatement(ctx *parser.StatementContext) any {
	fmt.Println("Statemant scope")

	fmt.Println(ctx.GetText())

	if ctx.Assigment() != nil {
		fmt.Println("if scope for assigment")

		return ctx.Assigment().Accept(v)
	}
	fmt.Println("if scope for expression")

	v.Res = ctx.Expression().Accept(v).(float64)

	return 0
}

func (v *CalcVisitor) VisitAssigment(ctx *parser.AssigmentContext) any {
	fmt.Println("assigment scope")

	id := ctx.ID().GetText()
	value := ctx.Expression().Accept(v).(float64)
	v.Variables[id] = value

	fmt.Printf("write to mem with id = %v and val = %v \n", id, value)

	return value

}

func (v *CalcVisitor) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
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

func (v *CalcVisitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
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

func (v *CalcVisitor) VisitNumber(ctx *parser.NumberContext) interface{} {
	val, _ := strconv.ParseFloat(ctx.GetText(), 64)

	return val
}

func (v *CalcVisitor) VisitID(ctx *parser.IDContext) any {
	id := ctx.GetText()

	if value, exists := v.Variables[id]; exists {
		return value
	}

	panic("no such var in memory!")
}
