// Code generated from Calc.g4 by ANTLR 4.10. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCalcListener is a complete listener for a parse tree produced by CalcParser.
type BaseCalcListener struct{}

var _ CalcListener = &BaseCalcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseCalcListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseCalcListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseCalcListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseCalcListener) ExitStatement(ctx *StatementContext) {}

// EnterPrintln is called when production println is entered.
func (s *BaseCalcListener) EnterPrintln(ctx *PrintlnContext) {}

// ExitPrintln is called when production println is exited.
func (s *BaseCalcListener) ExitPrintln(ctx *PrintlnContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseCalcListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseCalcListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseCalcListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseCalcListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterBasicExp is called when production BasicExp is entered.
func (s *BaseCalcListener) EnterBasicExp(ctx *BasicExpContext) {}

// ExitBasicExp is called when production BasicExp is exited.
func (s *BaseCalcListener) ExitBasicExp(ctx *BasicExpContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BaseCalcListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BaseCalcListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterElseStatement is called when production elseStatement is entered.
func (s *BaseCalcListener) EnterElseStatement(ctx *ElseStatementContext) {}

// ExitElseStatement is called when production elseStatement is exited.
func (s *BaseCalcListener) ExitElseStatement(ctx *ElseStatementContext) {}

// EnterAssigment is called when production assigment is entered.
func (s *BaseCalcListener) EnterAssigment(ctx *AssigmentContext) {}

// ExitAssigment is called when production assigment is exited.
func (s *BaseCalcListener) ExitAssigment(ctx *AssigmentContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseCalcListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseCalcListener) ExitNumber(ctx *NumberContext) {}

// EnterBool is called when production Bool is entered.
func (s *BaseCalcListener) EnterBool(ctx *BoolContext) {}

// ExitBool is called when production Bool is exited.
func (s *BaseCalcListener) ExitBool(ctx *BoolContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseCalcListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseCalcListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseCalcListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseCalcListener) ExitAddSub(ctx *AddSubContext) {}

// EnterID is called when production ID is entered.
func (s *BaseCalcListener) EnterID(ctx *IDContext) {}

// ExitID is called when production ID is exited.
func (s *BaseCalcListener) ExitID(ctx *IDContext) {}

// EnterLogicOp is called when production LogicOp is entered.
func (s *BaseCalcListener) EnterLogicOp(ctx *LogicOpContext) {}

// ExitLogicOp is called when production LogicOp is exited.
func (s *BaseCalcListener) ExitLogicOp(ctx *LogicOpContext) {}
