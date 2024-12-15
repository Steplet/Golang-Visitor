package visitor

import (
	"antlr4/lab/parser"
	"antlr4/lab/util"
	"fmt"
	"strings"
)

// Rework repeated assigment
//TODO logical op and var; If else; while; comments; println;

var addresCounter int = 0
var IfStatementCounter int = 0
var localBuilder strings.Builder

type CalcVisitor struct {
	*parser.BaseCalcVisitor
	Variables     map[string]string
	StringBuilder strings.Builder
}

func NewCalcVisitor() *CalcVisitor {
	return &CalcVisitor{Variables: make(map[string]string)} //id = valueId => a = res0 || a = 2x3
}

func (v *CalcVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {

	fmt.Printf("Child counter = %d\n", ctx.GetChildCount())
	for i := 0; i < ctx.GetChildCount(); i++ {
		if state := ctx.Statement(i); state != nil {
			state.Accept(v)
		}
	}

	// fmt.Printf("adrre counter = %d\n", addresCounter)

	v.StringBuilder.WriteString("ebreak")

	return 0
}

func (v *CalcVisitor) VisitStatement(ctx *parser.StatementContext) any {
	if ctx.Assigment() != nil {

		return ctx.Assigment().Accept(v)
	}

	if ctx.IfStatement() != nil {
		return ctx.IfStatement().Accept(v)
	}

	// ctx.Expression().Accept(v)

	return 0
}
func (v *CalcVisitor) VisitIfStatement(ctx *parser.IfStatementContext) any {
	// var localBuilder strings.Builder

	v.StringBuilder.WriteString(localBuilder.String())
	localBuilder.Reset()

	beqStringWithOutOffset := ctx.IfExpression().Accept(v).(string)
	ifAdderStart := addresCounter
	v.StringBuilder.WriteString(beqStringWithOutOffset)

	ctx.BlockStatement().Accept(v)

	offset := fmt.Sprintf("%d\n", addresCounter-ifAdderStart)
	v.StringBuilder.WriteString(offset)

	v.StringBuilder.WriteString(localBuilder.String())
	localBuilder.Reset()

	return 0
}

func (v *CalcVisitor) VisitBasicExp(ctx *parser.BasicExpContext) any {

	l := ctx.Expression(0).Accept(v)
	left := fmt.Sprintf("%v", l)

	r := ctx.Expression(1).Accept(v)
	right := fmt.Sprintf("%v", r)

	addresCounter++

	// fmt.Printf("adrre counter = %d\n", addresCounter)
	if ctx.GetOp().GetTokenType() == parser.CalcLexerEQUAL {

		riscString := util.IfBasicExpressionEqual(left, right)

		v.StringBuilder.WriteString(localBuilder.String())

		localBuilder.Reset()

		return riscString
	}

	if ctx.GetOp().GetTokenType() == parser.CalcLexerGREATHER {
		riscString := util.IfBasicExpressionGreather(left, right)

		v.StringBuilder.WriteString(localBuilder.String())

		localBuilder.Reset()

		return riscString
	}
	if ctx.GetOp().GetTokenType() == parser.CalcLexerLESS {
		riscString := util.IfBasicExpressionLess(left, right)

		v.StringBuilder.WriteString(localBuilder.String())

		localBuilder.Reset()

		return riscString
	}
	if ctx.GetOp().GetTokenType() == parser.CalcLexerNOT_EQUAL {
		riscString := util.IfBasicExpressionNotEqual(left, right)

		v.StringBuilder.WriteString(localBuilder.String())

		localBuilder.Reset()

		return riscString
	}

	return ""
}

func (v *CalcVisitor) VisitBlockStatement(ctx *parser.BlockStatementContext) any {

	for i := 0; i < ctx.GetChildCount(); i++ {
		if state := ctx.Statement(i); state != nil {
			state.Accept(v)
		}
	}

	return 0
}

func (v *CalcVisitor) VisitAssigment(ctx *parser.AssigmentContext) any {
	addresCounter++

	id := ctx.ID().GetText()
	// if val, ok := v.Variables[id]; ok {
	// 	valueId := ctx.Expression().Accept(v).(string)
	// 	resultId := util.MakeVarForResRisc()

	// }
	valueId := ctx.Expression().Accept(v).(string)
	v.Variables[id] = valueId

	riscString := util.AssigmentRiscFuncInit(id, valueId)

	localBuilder.WriteString(riscString)
	fmt.Printf("valueid = %s\n", valueId)

	return valueId

}

func (v *CalcVisitor) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
	addresCounter++

	left := ""
	right := ""

	l := ctx.Expression(0).Accept(v)
	left = fmt.Sprintf("%v", l)

	r := ctx.Expression(1).Accept(v)
	right = fmt.Sprintf("%v", r)

	switch ctx.GetOp().GetTokenType() {

	case parser.CalcLexerMUL:
		fmt.Printf("Mul with l = %v and r = %v\n", left, right)

		resId := util.MakeVarForResRisc()

		riscString := util.MultiRiscFunc(resId, left, right)
		localBuilder.WriteString(riscString)

		return resId

	case parser.CalcLexerDIV:

		fmt.Printf("Div with l = %v and r = %v\n", left, right)

		resId := util.MakeVarForResRisc()

		riscString := util.DivRiscFunc(resId, left, right)
		localBuilder.WriteString(riscString)

		return resId
	}

	return 0

}

func (v *CalcVisitor) VisitAddSub(ctx *parser.AddSubContext) any {
	addresCounter++

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
		localBuilder.WriteString(riscString)

		return resId

	case parser.CalcLexerSUB:

		resId := util.MakeVarForResRisc()

		riscString := util.SubIdRicsFunc(resId, left, right)
		localBuilder.WriteString(riscString)

		return resId
	}

	return 0

}

func (v *CalcVisitor) VisitNumber(ctx *parser.NumberContext) any {
	addresCounter++
	id := util.MakeNumberVar(ctx.GetText())
	riscString := util.AssigmentRiscFunc(id, ctx.GetText())
	localBuilder.WriteString(riscString)

	return id
}

func (v *CalcVisitor) VisitID(ctx *parser.IDContext) any {
	id := ctx.GetText()

	if _, exists := v.Variables[id]; exists {
		return id
	}

	panic("no such var in memory!")
}
