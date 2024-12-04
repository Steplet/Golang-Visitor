package visitor

import (
	"antlr4/lab/parser"
	"antlr4/lab/util"
	"fmt"
	"strings"
)

type CalcVisitor struct {
	*parser.BaseCalcVisitor
	Variables     map[string]string
	StringBuilder strings.Builder
}

func NewCalcVisitor() *CalcVisitor {
	return &CalcVisitor{Variables: make(map[string]string)}
}

func (v *CalcVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {

	for i := 0; i < ctx.GetChildCount(); i++ {
		if state := ctx.Statement(i); state != nil {
			state.Accept(v)
		}
	}

	v.StringBuilder.WriteString("ebreak")

	return 0
}

func (v *CalcVisitor) VisitStatement(ctx *parser.StatementContext) any {
	if ctx.Assigment() != nil {

		return ctx.Assigment().Accept(v)
	}

	fmt.Println("Expression Statment 0 in")

	ctx.Expression().Accept(v)
	fmt.Println("Expression Statment 0 out")

	return 0
}

func (v *CalcVisitor) VisitAssigment(ctx *parser.AssigmentContext) any {

	id := ctx.ID().GetText()
	value := ctx.Expression().Accept(v).(string)
	v.Variables[id] = value

	fmt.Printf("write to mem with id = %v and val = %v \n", id, value)

	riscString := util.AssigmentRiscFuncInit(id, value)

	v.StringBuilder.WriteString(riscString)

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

func (v *CalcVisitor) VisitAddSub(ctx *parser.AddSubContext) any {
	left := ""
	right := ""

	l := ctx.Expression(0).Accept(v)
	left = fmt.Sprintf("%v", l)

	r := ctx.Expression(1).Accept(v)
	right = fmt.Sprintf("%v", r)

	switch ctx.GetOp().GetTokenType() {

	case parser.CalcLexerADD:

		resId := util.MakeVarForResRisc()

		riscString := util.AddIdRiscFunc(resId, left, right)
		v.StringBuilder.WriteString(riscString)

		return resId

	case parser.CalcLexerSUB:

		resId := util.MakeVarForResRisc()

		riscString := util.SubIdRicsFunc(resId, left, right)
		v.StringBuilder.WriteString(riscString)

		return resId
	}

	return 0

}

func (v *CalcVisitor) VisitNumber(ctx *parser.NumberContext) any {

	id := util.MakeNumberVar(ctx.GetText())
	riscString := util.AssigmentRiscFunc(id, ctx.GetText())
	v.StringBuilder.WriteString(riscString)

	return id
}

func (v *CalcVisitor) VisitID(ctx *parser.IDContext) any {
	fmt.Println("IdFunc")
	id := ctx.GetText()

	if _, exists := v.Variables[id]; exists {
		return id
	}

	panic("no such var in memory!")
}
