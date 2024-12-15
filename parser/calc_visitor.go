// Code generated from Calc.g4 by ANTLR 4.10. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by CalcParser.
type CalcVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CalcParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by CalcParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by CalcParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by CalcParser#BasicExp.
	VisitBasicExp(ctx *BasicExpContext) interface{}

	// Visit a parse tree produced by CalcParser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) interface{}

	// Visit a parse tree produced by CalcParser#assigment.
	VisitAssigment(ctx *AssigmentContext) interface{}

	// Visit a parse tree produced by CalcParser#Number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by CalcParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by CalcParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by CalcParser#ID.
	VisitID(ctx *IDContext) interface{}
}
