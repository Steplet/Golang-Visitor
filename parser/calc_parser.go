// Code generated from Calc.g4 by ANTLR 4.10. DO NOT EDIT.

package parser // Calc

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CalcParser struct {
	*antlr.BaseParser
}

var calcParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calcParserInit() {
	staticData := &calcParserStaticData
	staticData.literalNames = []string{
		"", "'println'", "'('", "')'", "'while'", "'if'", "'{'", "'}'", "'else'",
		"'='", "'*'", "'/'", "'+'", "'-'", "'||'", "'&&'", "'=='", "'>'", "'>='",
		"'<'", "'<='", "'!='",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "MUL", "DIV", "ADD", "SUB",
		"LOGICAL_OR", "LOGICAL_AND", "EQUAL", "GREATHER", "GREATHER_OR_EQUAL",
		"LESS", "LESS_OR_EQUAL", "NOT_EQUAL", "NUMBER", "BOOL", "ID", "NEWLINE",
		"COMMENT", "LINE_COMMENT", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"program", "statement", "println", "whileStatement", "ifStatement",
		"ifExpression", "blockStatement", "elseStatement", "assigment", "expression",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 28, 110, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 5,
		0, 22, 8, 0, 10, 0, 12, 0, 25, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 31,
		8, 1, 1, 1, 1, 1, 3, 1, 35, 8, 1, 1, 1, 1, 1, 3, 1, 39, 8, 1, 1, 1, 1,
		1, 3, 1, 43, 8, 1, 1, 1, 1, 1, 3, 1, 47, 8, 1, 3, 1, 49, 8, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 3, 4, 68, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		5, 6, 76, 8, 6, 10, 6, 12, 6, 79, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 94, 8, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 105, 8, 9, 10, 9, 12, 9,
		108, 9, 9, 1, 9, 0, 1, 18, 10, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 4,
		3, 0, 16, 17, 19, 19, 21, 21, 1, 0, 10, 11, 1, 0, 12, 13, 1, 0, 14, 15,
		116, 0, 23, 1, 0, 0, 0, 2, 48, 1, 0, 0, 0, 4, 50, 1, 0, 0, 0, 6, 55, 1,
		0, 0, 0, 8, 61, 1, 0, 0, 0, 10, 69, 1, 0, 0, 0, 12, 73, 1, 0, 0, 0, 14,
		82, 1, 0, 0, 0, 16, 85, 1, 0, 0, 0, 18, 93, 1, 0, 0, 0, 20, 22, 3, 2, 1,
		0, 21, 20, 1, 0, 0, 0, 22, 25, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 23, 24,
		1, 0, 0, 0, 24, 26, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 26, 27, 5, 0, 0, 1,
		27, 1, 1, 0, 0, 0, 28, 30, 3, 16, 8, 0, 29, 31, 5, 25, 0, 0, 30, 29, 1,
		0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 49, 1, 0, 0, 0, 32, 34, 3, 8, 4, 0, 33,
		35, 5, 25, 0, 0, 34, 33, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 49, 1, 0,
		0, 0, 36, 38, 3, 12, 6, 0, 37, 39, 5, 25, 0, 0, 38, 37, 1, 0, 0, 0, 38,
		39, 1, 0, 0, 0, 39, 49, 1, 0, 0, 0, 40, 42, 3, 6, 3, 0, 41, 43, 5, 25,
		0, 0, 42, 41, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 49, 1, 0, 0, 0, 44, 46,
		3, 4, 2, 0, 45, 47, 5, 25, 0, 0, 46, 45, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0,
		47, 49, 1, 0, 0, 0, 48, 28, 1, 0, 0, 0, 48, 32, 1, 0, 0, 0, 48, 36, 1,
		0, 0, 0, 48, 40, 1, 0, 0, 0, 48, 44, 1, 0, 0, 0, 49, 3, 1, 0, 0, 0, 50,
		51, 5, 1, 0, 0, 51, 52, 5, 2, 0, 0, 52, 53, 3, 18, 9, 0, 53, 54, 5, 3,
		0, 0, 54, 5, 1, 0, 0, 0, 55, 56, 5, 4, 0, 0, 56, 57, 5, 2, 0, 0, 57, 58,
		3, 10, 5, 0, 58, 59, 5, 3, 0, 0, 59, 60, 3, 12, 6, 0, 60, 7, 1, 0, 0, 0,
		61, 62, 5, 5, 0, 0, 62, 63, 5, 2, 0, 0, 63, 64, 3, 10, 5, 0, 64, 65, 5,
		3, 0, 0, 65, 67, 3, 12, 6, 0, 66, 68, 3, 14, 7, 0, 67, 66, 1, 0, 0, 0,
		67, 68, 1, 0, 0, 0, 68, 9, 1, 0, 0, 0, 69, 70, 3, 18, 9, 0, 70, 71, 7,
		0, 0, 0, 71, 72, 3, 18, 9, 0, 72, 11, 1, 0, 0, 0, 73, 77, 5, 6, 0, 0, 74,
		76, 3, 2, 1, 0, 75, 74, 1, 0, 0, 0, 76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0,
		0, 77, 78, 1, 0, 0, 0, 78, 80, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 81,
		5, 7, 0, 0, 81, 13, 1, 0, 0, 0, 82, 83, 5, 8, 0, 0, 83, 84, 3, 12, 6, 0,
		84, 15, 1, 0, 0, 0, 85, 86, 5, 24, 0, 0, 86, 87, 5, 9, 0, 0, 87, 88, 3,
		18, 9, 0, 88, 17, 1, 0, 0, 0, 89, 90, 6, 9, -1, 0, 90, 94, 5, 22, 0, 0,
		91, 94, 5, 24, 0, 0, 92, 94, 5, 23, 0, 0, 93, 89, 1, 0, 0, 0, 93, 91, 1,
		0, 0, 0, 93, 92, 1, 0, 0, 0, 94, 106, 1, 0, 0, 0, 95, 96, 10, 6, 0, 0,
		96, 97, 7, 1, 0, 0, 97, 105, 3, 18, 9, 7, 98, 99, 10, 5, 0, 0, 99, 100,
		7, 2, 0, 0, 100, 105, 3, 18, 9, 6, 101, 102, 10, 4, 0, 0, 102, 103, 7,
		3, 0, 0, 103, 105, 3, 18, 9, 5, 104, 95, 1, 0, 0, 0, 104, 98, 1, 0, 0,
		0, 104, 101, 1, 0, 0, 0, 105, 108, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106,
		107, 1, 0, 0, 0, 107, 19, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 12, 23, 30,
		34, 38, 42, 46, 48, 67, 77, 93, 104, 106,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CalcParserInit initializes any static state used to implement CalcParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCalcParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalcParserInit() {
	staticData := &calcParserStaticData
	staticData.once.Do(calcParserInit)
}

// NewCalcParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCalcParser(input antlr.TokenStream) *CalcParser {
	CalcParserInit()
	this := new(CalcParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &calcParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Calc.g4"

	return this
}

// CalcParser tokens.
const (
	CalcParserEOF               = antlr.TokenEOF
	CalcParserT__0              = 1
	CalcParserT__1              = 2
	CalcParserT__2              = 3
	CalcParserT__3              = 4
	CalcParserT__4              = 5
	CalcParserT__5              = 6
	CalcParserT__6              = 7
	CalcParserT__7              = 8
	CalcParserT__8              = 9
	CalcParserMUL               = 10
	CalcParserDIV               = 11
	CalcParserADD               = 12
	CalcParserSUB               = 13
	CalcParserLOGICAL_OR        = 14
	CalcParserLOGICAL_AND       = 15
	CalcParserEQUAL             = 16
	CalcParserGREATHER          = 17
	CalcParserGREATHER_OR_EQUAL = 18
	CalcParserLESS              = 19
	CalcParserLESS_OR_EQUAL     = 20
	CalcParserNOT_EQUAL         = 21
	CalcParserNUMBER            = 22
	CalcParserBOOL              = 23
	CalcParserID                = 24
	CalcParserNEWLINE           = 25
	CalcParserCOMMENT           = 26
	CalcParserLINE_COMMENT      = 27
	CalcParserWHITESPACE        = 28
)

// CalcParser rules.
const (
	CalcParserRULE_program        = 0
	CalcParserRULE_statement      = 1
	CalcParserRULE_println        = 2
	CalcParserRULE_whileStatement = 3
	CalcParserRULE_ifStatement    = 4
	CalcParserRULE_ifExpression   = 5
	CalcParserRULE_blockStatement = 6
	CalcParserRULE_elseStatement  = 7
	CalcParserRULE_assigment      = 8
	CalcParserRULE_expression     = 9
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(CalcParserEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CalcParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CalcParserT__0)|(1<<CalcParserT__3)|(1<<CalcParserT__4)|(1<<CalcParserT__5)|(1<<CalcParserID))) != 0 {
		{
			p.SetState(20)
			p.Statement()
		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(26)
		p.Match(CalcParserEOF)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assigment() IAssigmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssigmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssigmentContext)
}

func (s *StatementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(CalcParserNEWLINE, 0)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) BlockStatement() IBlockStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *StatementContext) WhileStatement() IWhileStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *StatementContext) Println() IPrintlnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintlnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintlnContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CalcParserRULE_statement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(48)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalcParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.Assigment()
		}
		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CalcParserNEWLINE {
			{
				p.SetState(29)
				p.Match(CalcParserNEWLINE)
			}

		}

	case CalcParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(32)
			p.IfStatement()
		}
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CalcParserNEWLINE {
			{
				p.SetState(33)
				p.Match(CalcParserNEWLINE)
			}

		}

	case CalcParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(36)
			p.BlockStatement()
		}
		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CalcParserNEWLINE {
			{
				p.SetState(37)
				p.Match(CalcParserNEWLINE)
			}

		}

	case CalcParserT__3:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(40)
			p.WhileStatement()
		}
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CalcParserNEWLINE {
			{
				p.SetState(41)
				p.Match(CalcParserNEWLINE)
			}

		}

	case CalcParserT__0:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(44)
			p.Println()
		}
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CalcParserNEWLINE {
			{
				p.SetState(45)
				p.Match(CalcParserNEWLINE)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrintlnContext is an interface to support dynamic dispatch.
type IPrintlnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrintlnContext differentiates from other interfaces.
	IsPrintlnContext()
}

type PrintlnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintlnContext() *PrintlnContext {
	var p = new(PrintlnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_println
	return p
}

func (*PrintlnContext) IsPrintlnContext() {}

func NewPrintlnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintlnContext {
	var p = new(PrintlnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_println

	return p
}

func (s *PrintlnContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintlnContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrintlnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintlnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintlnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterPrintln(s)
	}
}

func (s *PrintlnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitPrintln(s)
	}
}

func (s *PrintlnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcVisitor:
		return t.VisitPrintln(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcParser) Println() (localctx IPrintlnContext) {
	this := p
	_ = this

	localctx = NewPrintlnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CalcParserRULE_println)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(CalcParserT__0)
	}
	{
		p.SetState(51)
		p.Match(CalcParserT__1)
	}
	{
		p.SetState(52)
		p.expression(0)
	}
	{
		p.SetState(53)
		p.Match(CalcParserT__2)
	}

	return localctx
}

// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_whileStatement
	return p
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) IfExpression() IIfExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfExpressionContext)
}

func (s *WhileStatementContext) BlockStatement() IBlockStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcVisitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcParser) WhileStatement() (localctx IWhileStatementContext) {
	this := p
	_ = this

	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CalcParserRULE_whileStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.Match(CalcParserT__3)
	}
	{
		p.SetState(56)
		p.Match(CalcParserT__1)
	}
	{
		p.SetState(57)
		p.IfExpression()
	}
	{
		p.SetState(58)
		p.Match(CalcParserT__2)
	}
	{
		p.SetState(59)
		p.BlockStatement()
	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IfExpression() IIfExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfExpressionContext)
}

func (s *IfStatementContext) BlockStatement() IBlockStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *IfStatementContext) ElseStatement() IElseStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseStatementContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcParser) IfStatement() (localctx IIfStatementContext) {
	this := p
	_ = this

	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CalcParserRULE_ifStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Match(CalcParserT__4)
	}
	{
		p.SetState(62)
		p.Match(CalcParserT__1)
	}
	{
		p.SetState(63)
		p.IfExpression()
	}
	{
		p.SetState(64)
		p.Match(CalcParserT__2)
	}
	{
		p.SetState(65)
		p.BlockStatement()
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CalcParserT__7 {
		{
			p.SetState(66)
			p.ElseStatement()
		}

	}

	return localctx
}

// IIfExpressionContext is an interface to support dynamic dispatch.
type IIfExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfExpressionContext differentiates from other interfaces.
	IsIfExpressionContext()
}

type IfExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfExpressionContext() *IfExpressionContext {
	var p = new(IfExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_ifExpression
	return p
}

func (*IfExpressionContext) IsIfExpressionContext() {}

func NewIfExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfExpressionContext {
	var p = new(IfExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_ifExpression

	return p
}

func (s *IfExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *IfExpressionContext) CopyFrom(ctx *IfExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *IfExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BasicExpContext struct {
	*IfExpressionContext
	op antlr.Token
}

func NewBasicExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BasicExpContext {
	var p = new(BasicExpContext)

	p.IfExpressionContext = NewEmptyIfExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IfExpressionContext))

	return p
}

func (s *BasicExpContext) GetOp() antlr.Token { return s.op }

func (s *BasicExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *BasicExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicExpContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BasicExpContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BasicExpContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(CalcParserEQUAL, 0)
}

func (s *BasicExpContext) LESS() antlr.TerminalNode {
	return s.GetToken(CalcParserLESS, 0)
}

func (s *BasicExpContext) GREATHER() antlr.TerminalNode {
	return s.GetToken(CalcParserGREATHER, 0)
}

func (s *BasicExpContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(CalcParserNOT_EQUAL, 0)
}

func (s *BasicExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterBasicExp(s)
	}
}

func (s *BasicExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitBasicExp(s)
	}
}

func (s *BasicExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcVisitor:
		return t.VisitBasicExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcParser) IfExpression() (localctx IIfExpressionContext) {
	this := p
	_ = this

	localctx = NewIfExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CalcParserRULE_ifExpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewBasicExpContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.expression(0)
	}
	{
		p.SetState(70)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*BasicExpContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CalcParserEQUAL)|(1<<CalcParserGREATHER)|(1<<CalcParserLESS)|(1<<CalcParserNOT_EQUAL))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*BasicExpContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(71)
		p.expression(0)
	}

	return localctx
}

// IBlockStatementContext is an interface to support dynamic dispatch.
type IBlockStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockStatementContext differentiates from other interfaces.
	IsBlockStatementContext()
}

type BlockStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStatementContext() *BlockStatementContext {
	var p = new(BlockStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_blockStatement
	return p
}

func (*BlockStatementContext) IsBlockStatementContext() {}

func NewBlockStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStatementContext {
	var p = new(BlockStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_blockStatement

	return p
}

func (s *BlockStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterBlockStatement(s)
	}
}

func (s *BlockStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitBlockStatement(s)
	}
}

func (s *BlockStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcVisitor:
		return t.VisitBlockStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcParser) BlockStatement() (localctx IBlockStatementContext) {
	this := p
	_ = this

	localctx = NewBlockStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CalcParserRULE_blockStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.Match(CalcParserT__5)
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CalcParserT__0)|(1<<CalcParserT__3)|(1<<CalcParserT__4)|(1<<CalcParserT__5)|(1<<CalcParserID))) != 0 {
		{
			p.SetState(74)
			p.Statement()
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(80)
		p.Match(CalcParserT__6)
	}

	return localctx
}

// IElseStatementContext is an interface to support dynamic dispatch.
type IElseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseStatementContext differentiates from other interfaces.
	IsElseStatementContext()
}

type ElseStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStatementContext() *ElseStatementContext {
	var p = new(ElseStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_elseStatement
	return p
}

func (*ElseStatementContext) IsElseStatementContext() {}

func NewElseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStatementContext {
	var p = new(ElseStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_elseStatement

	return p
}

func (s *ElseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStatementContext) BlockStatement() IBlockStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *ElseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterElseStatement(s)
	}
}

func (s *ElseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitElseStatement(s)
	}
}

func (s *ElseStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcVisitor:
		return t.VisitElseStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcParser) ElseStatement() (localctx IElseStatementContext) {
	this := p
	_ = this

	localctx = NewElseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CalcParserRULE_elseStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(CalcParserT__7)
	}
	{
		p.SetState(83)
		p.BlockStatement()
	}

	return localctx
}

// IAssigmentContext is an interface to support dynamic dispatch.
type IAssigmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssigmentContext differentiates from other interfaces.
	IsAssigmentContext()
}

type AssigmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssigmentContext() *AssigmentContext {
	var p = new(AssigmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_assigment
	return p
}

func (*AssigmentContext) IsAssigmentContext() {}

func NewAssigmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssigmentContext {
	var p = new(AssigmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_assigment

	return p
}

func (s *AssigmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssigmentContext) ID() antlr.TerminalNode {
	return s.GetToken(CalcParserID, 0)
}

func (s *AssigmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssigmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssigmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssigmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterAssigment(s)
	}
}

func (s *AssigmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitAssigment(s)
	}
}

func (s *AssigmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcVisitor:
		return t.VisitAssigment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcParser) Assigment() (localctx IAssigmentContext) {
	this := p
	_ = this

	localctx = NewAssigmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CalcParserRULE_assigment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(CalcParserID)
	}
	{
		p.SetState(86)
		p.Match(CalcParserT__8)
	}
	{
		p.SetState(87)
		p.expression(0)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NumberContext struct {
	*ExpressionContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CalcParserNUMBER, 0)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolContext struct {
	*ExpressionContext
}

func NewBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolContext {
	var p = new(BoolContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolContext) BOOL() antlr.TerminalNode {
	return s.GetToken(CalcParserBOOL, 0)
}

func (s *BoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterBool(s)
	}
}

func (s *BoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitBool(s)
	}
}

func (s *BoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcVisitor:
		return t.VisitBool(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulDivContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MulDivContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulDivContext) MUL() antlr.TerminalNode {
	return s.GetToken(CalcParserMUL, 0)
}

func (s *MulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(CalcParserDIV, 0)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

func (s *MulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcVisitor:
		return t.VisitMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AddSubContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddSubContext) ADD() antlr.TerminalNode {
	return s.GetToken(CalcParserADD, 0)
}

func (s *AddSubContext) SUB() antlr.TerminalNode {
	return s.GetToken(CalcParserSUB, 0)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type IDContext struct {
	*ExpressionContext
}

func NewIDContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IDContext {
	var p = new(IDContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IDContext) ID() antlr.TerminalNode {
	return s.GetToken(CalcParserID, 0)
}

func (s *IDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterID(s)
	}
}

func (s *IDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitID(s)
	}
}

func (s *IDContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcVisitor:
		return t.VisitID(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicOpContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewLogicOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicOpContext {
	var p = new(LogicOpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LogicOpContext) GetOp() antlr.Token { return s.op }

func (s *LogicOpContext) SetOp(v antlr.Token) { s.op = v }

func (s *LogicOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicOpContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicOpContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogicOpContext) LOGICAL_AND() antlr.TerminalNode {
	return s.GetToken(CalcParserLOGICAL_AND, 0)
}

func (s *LogicOpContext) LOGICAL_OR() antlr.TerminalNode {
	return s.GetToken(CalcParserLOGICAL_OR, 0)
}

func (s *LogicOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterLogicOp(s)
	}
}

func (s *LogicOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitLogicOp(s)
	}
}

func (s *LogicOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcVisitor:
		return t.VisitLogicOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *CalcParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, CalcParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(93)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalcParserNUMBER:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(90)
			p.Match(CalcParserNUMBER)
		}

	case CalcParserID:
		localctx = NewIDContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(91)
			p.Match(CalcParserID)
		}

	case CalcParserBOOL:
		localctx = NewBoolContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(92)
			p.Match(CalcParserBOOL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(104)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expression)
				p.SetState(95)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(96)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcParserMUL || _la == CalcParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(97)
					p.expression(7)
				}

			case 2:
				localctx = NewAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expression)
				p.SetState(98)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(99)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcParserADD || _la == CalcParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(100)
					p.expression(6)
				}

			case 3:
				localctx = NewLogicOpContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expression)
				p.SetState(101)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(102)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*LogicOpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcParserLOGICAL_OR || _la == CalcParserLOGICAL_AND) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*LogicOpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(103)
					p.expression(5)
				}

			}

		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

func (p *CalcParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CalcParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
