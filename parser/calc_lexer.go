// Code generated from Calc.g4 by ANTLR 4.10. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CalcLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var calclexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calclexerLexerInit() {
	staticData := &calclexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'if'", "'('", "')'", "'{'", "'}'", "'='", "'*'", "'/'", "'+'",
		"'-'", "'=='", "'>'", "'>='", "'<'", "'<='", "'!='",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "MUL", "DIV", "ADD", "SUB", "EQUAL", "GREATHER",
		"GREATHER_OR_EQUAL", "LESS", "LESS_OR_EQUAL", "NOT_EQUAL", "NUMBER",
		"ID", "NEWLINE", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "MUL", "DIV", "ADD",
		"SUB", "EQUAL", "GREATHER", "GREATHER_OR_EQUAL", "LESS", "LESS_OR_EQUAL",
		"NOT_EQUAL", "NUMBER", "ID", "NEWLINE", "WHITESPACE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 20, 102, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 15, 1, 16, 4, 16, 80, 8, 16, 11, 16, 12, 16, 81, 1, 17, 1, 17, 5,
		17, 86, 8, 17, 10, 17, 12, 17, 89, 9, 17, 1, 18, 3, 18, 92, 8, 18, 1, 18,
		1, 18, 1, 19, 4, 19, 97, 8, 19, 11, 19, 12, 19, 98, 1, 19, 1, 19, 0, 0,
		20, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21,
		11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39,
		20, 1, 0, 4, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57,
		65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 105, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		1, 41, 1, 0, 0, 0, 3, 44, 1, 0, 0, 0, 5, 46, 1, 0, 0, 0, 7, 48, 1, 0, 0,
		0, 9, 50, 1, 0, 0, 0, 11, 52, 1, 0, 0, 0, 13, 54, 1, 0, 0, 0, 15, 56, 1,
		0, 0, 0, 17, 58, 1, 0, 0, 0, 19, 60, 1, 0, 0, 0, 21, 62, 1, 0, 0, 0, 23,
		65, 1, 0, 0, 0, 25, 67, 1, 0, 0, 0, 27, 70, 1, 0, 0, 0, 29, 72, 1, 0, 0,
		0, 31, 75, 1, 0, 0, 0, 33, 79, 1, 0, 0, 0, 35, 83, 1, 0, 0, 0, 37, 91,
		1, 0, 0, 0, 39, 96, 1, 0, 0, 0, 41, 42, 5, 105, 0, 0, 42, 43, 5, 102, 0,
		0, 43, 2, 1, 0, 0, 0, 44, 45, 5, 40, 0, 0, 45, 4, 1, 0, 0, 0, 46, 47, 5,
		41, 0, 0, 47, 6, 1, 0, 0, 0, 48, 49, 5, 123, 0, 0, 49, 8, 1, 0, 0, 0, 50,
		51, 5, 125, 0, 0, 51, 10, 1, 0, 0, 0, 52, 53, 5, 61, 0, 0, 53, 12, 1, 0,
		0, 0, 54, 55, 5, 42, 0, 0, 55, 14, 1, 0, 0, 0, 56, 57, 5, 47, 0, 0, 57,
		16, 1, 0, 0, 0, 58, 59, 5, 43, 0, 0, 59, 18, 1, 0, 0, 0, 60, 61, 5, 45,
		0, 0, 61, 20, 1, 0, 0, 0, 62, 63, 5, 61, 0, 0, 63, 64, 5, 61, 0, 0, 64,
		22, 1, 0, 0, 0, 65, 66, 5, 62, 0, 0, 66, 24, 1, 0, 0, 0, 67, 68, 5, 62,
		0, 0, 68, 69, 5, 61, 0, 0, 69, 26, 1, 0, 0, 0, 70, 71, 5, 60, 0, 0, 71,
		28, 1, 0, 0, 0, 72, 73, 5, 60, 0, 0, 73, 74, 5, 61, 0, 0, 74, 30, 1, 0,
		0, 0, 75, 76, 5, 33, 0, 0, 76, 77, 5, 61, 0, 0, 77, 32, 1, 0, 0, 0, 78,
		80, 7, 0, 0, 0, 79, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 79, 1, 0, 0,
		0, 81, 82, 1, 0, 0, 0, 82, 34, 1, 0, 0, 0, 83, 87, 7, 1, 0, 0, 84, 86,
		7, 2, 0, 0, 85, 84, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0,
		87, 88, 1, 0, 0, 0, 88, 36, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 92, 5,
		13, 0, 0, 91, 90, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93,
		94, 5, 10, 0, 0, 94, 38, 1, 0, 0, 0, 95, 97, 7, 3, 0, 0, 96, 95, 1, 0,
		0, 0, 97, 98, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 100,
		1, 0, 0, 0, 100, 101, 6, 19, 0, 0, 101, 40, 1, 0, 0, 0, 5, 0, 81, 87, 91,
		98, 1, 6, 0, 0,
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

// CalcLexerInit initializes any static state used to implement CalcLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCalcLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalcLexerInit() {
	staticData := &calclexerLexerStaticData
	staticData.once.Do(calclexerLexerInit)
}

// NewCalcLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCalcLexer(input antlr.CharStream) *CalcLexer {
	CalcLexerInit()
	l := new(CalcLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &calclexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Calc.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalcLexer tokens.
const (
	CalcLexerT__0              = 1
	CalcLexerT__1              = 2
	CalcLexerT__2              = 3
	CalcLexerT__3              = 4
	CalcLexerT__4              = 5
	CalcLexerT__5              = 6
	CalcLexerMUL               = 7
	CalcLexerDIV               = 8
	CalcLexerADD               = 9
	CalcLexerSUB               = 10
	CalcLexerEQUAL             = 11
	CalcLexerGREATHER          = 12
	CalcLexerGREATHER_OR_EQUAL = 13
	CalcLexerLESS              = 14
	CalcLexerLESS_OR_EQUAL     = 15
	CalcLexerNOT_EQUAL         = 16
	CalcLexerNUMBER            = 17
	CalcLexerID                = 18
	CalcLexerNEWLINE           = 19
	CalcLexerWHITESPACE        = 20
)
