// Code generated from Calc.g4 by ANTLR 4.10. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CalcListener is a complete listener for a parse tree produced by CalcParser.
type CalcListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterPrintln is called when entering the println production.
	EnterPrintln(c *PrintlnContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterBasicExp is called when entering the BasicExp production.
	EnterBasicExp(c *BasicExpContext)

	// EnterBlockStatement is called when entering the blockStatement production.
	EnterBlockStatement(c *BlockStatementContext)

	// EnterElseStatement is called when entering the elseStatement production.
	EnterElseStatement(c *ElseStatementContext)

	// EnterAssigment is called when entering the assigment production.
	EnterAssigment(c *AssigmentContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterBool is called when entering the Bool production.
	EnterBool(c *BoolContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterID is called when entering the ID production.
	EnterID(c *IDContext)

	// EnterLogicOp is called when entering the LogicOp production.
	EnterLogicOp(c *LogicOpContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitPrintln is called when exiting the println production.
	ExitPrintln(c *PrintlnContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitBasicExp is called when exiting the BasicExp production.
	ExitBasicExp(c *BasicExpContext)

	// ExitBlockStatement is called when exiting the blockStatement production.
	ExitBlockStatement(c *BlockStatementContext)

	// ExitElseStatement is called when exiting the elseStatement production.
	ExitElseStatement(c *ElseStatementContext)

	// ExitAssigment is called when exiting the assigment production.
	ExitAssigment(c *AssigmentContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitBool is called when exiting the Bool production.
	ExitBool(c *BoolContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitID is called when exiting the ID production.
	ExitID(c *IDContext)

	// ExitLogicOp is called when exiting the LogicOp production.
	ExitLogicOp(c *LogicOpContext)
}
