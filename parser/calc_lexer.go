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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"MUL", "DIV", "ADD", "SUB", "LOGICAL_OR", "LOGICAL_AND", "EQUAL", "GREATHER",
		"GREATHER_OR_EQUAL", "LESS", "LESS_OR_EQUAL", "NOT_EQUAL", "NUMBER",
		"BOOL", "ID", "NEWLINE", "COMMENT", "LINE_COMMENT", "WHITESPACE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 28, 179, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21,
		4, 21, 121, 8, 21, 11, 21, 12, 21, 122, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 134, 8, 22, 1, 23, 1, 23, 5, 23,
		138, 8, 23, 10, 23, 12, 23, 141, 9, 23, 1, 24, 3, 24, 144, 8, 24, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 152, 8, 25, 10, 25, 12, 25, 155,
		9, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 5,
		26, 166, 8, 26, 10, 26, 12, 26, 169, 9, 26, 1, 26, 1, 26, 1, 27, 4, 27,
		174, 8, 27, 11, 27, 12, 27, 175, 1, 27, 1, 27, 1, 153, 0, 28, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12,
		25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21,
		43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 1, 0, 5, 1, 0,
		48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97,
		122, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13, 32, 32, 185, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0,
		0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0,
		0, 0, 1, 57, 1, 0, 0, 0, 3, 65, 1, 0, 0, 0, 5, 67, 1, 0, 0, 0, 7, 69, 1,
		0, 0, 0, 9, 75, 1, 0, 0, 0, 11, 78, 1, 0, 0, 0, 13, 80, 1, 0, 0, 0, 15,
		82, 1, 0, 0, 0, 17, 87, 1, 0, 0, 0, 19, 89, 1, 0, 0, 0, 21, 91, 1, 0, 0,
		0, 23, 93, 1, 0, 0, 0, 25, 95, 1, 0, 0, 0, 27, 97, 1, 0, 0, 0, 29, 100,
		1, 0, 0, 0, 31, 103, 1, 0, 0, 0, 33, 106, 1, 0, 0, 0, 35, 108, 1, 0, 0,
		0, 37, 111, 1, 0, 0, 0, 39, 113, 1, 0, 0, 0, 41, 116, 1, 0, 0, 0, 43, 120,
		1, 0, 0, 0, 45, 133, 1, 0, 0, 0, 47, 135, 1, 0, 0, 0, 49, 143, 1, 0, 0,
		0, 51, 147, 1, 0, 0, 0, 53, 161, 1, 0, 0, 0, 55, 173, 1, 0, 0, 0, 57, 58,
		5, 112, 0, 0, 58, 59, 5, 114, 0, 0, 59, 60, 5, 105, 0, 0, 60, 61, 5, 110,
		0, 0, 61, 62, 5, 116, 0, 0, 62, 63, 5, 108, 0, 0, 63, 64, 5, 110, 0, 0,
		64, 2, 1, 0, 0, 0, 65, 66, 5, 40, 0, 0, 66, 4, 1, 0, 0, 0, 67, 68, 5, 41,
		0, 0, 68, 6, 1, 0, 0, 0, 69, 70, 5, 119, 0, 0, 70, 71, 5, 104, 0, 0, 71,
		72, 5, 105, 0, 0, 72, 73, 5, 108, 0, 0, 73, 74, 5, 101, 0, 0, 74, 8, 1,
		0, 0, 0, 75, 76, 5, 105, 0, 0, 76, 77, 5, 102, 0, 0, 77, 10, 1, 0, 0, 0,
		78, 79, 5, 123, 0, 0, 79, 12, 1, 0, 0, 0, 80, 81, 5, 125, 0, 0, 81, 14,
		1, 0, 0, 0, 82, 83, 5, 101, 0, 0, 83, 84, 5, 108, 0, 0, 84, 85, 5, 115,
		0, 0, 85, 86, 5, 101, 0, 0, 86, 16, 1, 0, 0, 0, 87, 88, 5, 61, 0, 0, 88,
		18, 1, 0, 0, 0, 89, 90, 5, 42, 0, 0, 90, 20, 1, 0, 0, 0, 91, 92, 5, 47,
		0, 0, 92, 22, 1, 0, 0, 0, 93, 94, 5, 43, 0, 0, 94, 24, 1, 0, 0, 0, 95,
		96, 5, 45, 0, 0, 96, 26, 1, 0, 0, 0, 97, 98, 5, 124, 0, 0, 98, 99, 5, 124,
		0, 0, 99, 28, 1, 0, 0, 0, 100, 101, 5, 38, 0, 0, 101, 102, 5, 38, 0, 0,
		102, 30, 1, 0, 0, 0, 103, 104, 5, 61, 0, 0, 104, 105, 5, 61, 0, 0, 105,
		32, 1, 0, 0, 0, 106, 107, 5, 62, 0, 0, 107, 34, 1, 0, 0, 0, 108, 109, 5,
		62, 0, 0, 109, 110, 5, 61, 0, 0, 110, 36, 1, 0, 0, 0, 111, 112, 5, 60,
		0, 0, 112, 38, 1, 0, 0, 0, 113, 114, 5, 60, 0, 0, 114, 115, 5, 61, 0, 0,
		115, 40, 1, 0, 0, 0, 116, 117, 5, 33, 0, 0, 117, 118, 5, 61, 0, 0, 118,
		42, 1, 0, 0, 0, 119, 121, 7, 0, 0, 0, 120, 119, 1, 0, 0, 0, 121, 122, 1,
		0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 44, 1, 0, 0,
		0, 124, 125, 5, 84, 0, 0, 125, 126, 5, 114, 0, 0, 126, 127, 5, 117, 0,
		0, 127, 134, 5, 101, 0, 0, 128, 129, 5, 70, 0, 0, 129, 130, 5, 97, 0, 0,
		130, 131, 5, 108, 0, 0, 131, 132, 5, 115, 0, 0, 132, 134, 5, 101, 0, 0,
		133, 124, 1, 0, 0, 0, 133, 128, 1, 0, 0, 0, 134, 46, 1, 0, 0, 0, 135, 139,
		7, 1, 0, 0, 136, 138, 7, 2, 0, 0, 137, 136, 1, 0, 0, 0, 138, 141, 1, 0,
		0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 48, 1, 0, 0, 0,
		141, 139, 1, 0, 0, 0, 142, 144, 5, 13, 0, 0, 143, 142, 1, 0, 0, 0, 143,
		144, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 146, 5, 10, 0, 0, 146, 50,
		1, 0, 0, 0, 147, 148, 5, 47, 0, 0, 148, 149, 5, 42, 0, 0, 149, 153, 1,
		0, 0, 0, 150, 152, 9, 0, 0, 0, 151, 150, 1, 0, 0, 0, 152, 155, 1, 0, 0,
		0, 153, 154, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 154, 156, 1, 0, 0, 0, 155,
		153, 1, 0, 0, 0, 156, 157, 5, 42, 0, 0, 157, 158, 5, 47, 0, 0, 158, 159,
		1, 0, 0, 0, 159, 160, 6, 25, 0, 0, 160, 52, 1, 0, 0, 0, 161, 162, 5, 47,
		0, 0, 162, 163, 5, 47, 0, 0, 163, 167, 1, 0, 0, 0, 164, 166, 8, 3, 0, 0,
		165, 164, 1, 0, 0, 0, 166, 169, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 167,
		168, 1, 0, 0, 0, 168, 170, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 170, 171,
		6, 26, 0, 0, 171, 54, 1, 0, 0, 0, 172, 174, 7, 4, 0, 0, 173, 172, 1, 0,
		0, 0, 174, 175, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0,
		176, 177, 1, 0, 0, 0, 177, 178, 6, 27, 0, 0, 178, 56, 1, 0, 0, 0, 8, 0,
		122, 133, 139, 143, 153, 167, 175, 1, 6, 0, 0,
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
	CalcLexerT__6              = 7
	CalcLexerT__7              = 8
	CalcLexerT__8              = 9
	CalcLexerMUL               = 10
	CalcLexerDIV               = 11
	CalcLexerADD               = 12
	CalcLexerSUB               = 13
	CalcLexerLOGICAL_OR        = 14
	CalcLexerLOGICAL_AND       = 15
	CalcLexerEQUAL             = 16
	CalcLexerGREATHER          = 17
	CalcLexerGREATHER_OR_EQUAL = 18
	CalcLexerLESS              = 19
	CalcLexerLESS_OR_EQUAL     = 20
	CalcLexerNOT_EQUAL         = 21
	CalcLexerNUMBER            = 22
	CalcLexerBOOL              = 23
	CalcLexerID                = 24
	CalcLexerNEWLINE           = 25
	CalcLexerCOMMENT           = 26
	CalcLexerLINE_COMMENT      = 27
	CalcLexerWHITESPACE        = 28
)
