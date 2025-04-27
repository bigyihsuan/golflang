// Code generated from GolflangTokens.g4 by ANTLR 4.13.2. DO NOT EDIT.

package par

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type GolflangTokens struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var GolflangTokensLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func golflangtokensLexerInit() {
	staticData := &GolflangTokensLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "'true'", "'false'", "'\"'", "'('", "')'", "'['",
		"']'", "'{'", "'}'", "'.'", "','", "':'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "NEWLINE", "INT", "DEC", "STR", "TRUE", "FALSE", "QUOTE",
		"LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "LBRACE", "RBRACE", "DOT",
		"COMMA", "COLON", "SEMICOLON",
	}
	staticData.RuleNames = []string{
		"WHITESPACE", "NEWLINE", "DIGITS", "INT", "DEC", "STR", "TRUE", "FALSE",
		"QUOTE", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "LBRACE", "RBRACE",
		"DOT", "COMMA", "COLON", "SEMICOLON",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 18, 106, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 4, 0, 41, 8, 0,
		11, 0, 12, 0, 42, 1, 0, 1, 0, 1, 1, 4, 1, 48, 8, 1, 11, 1, 12, 1, 49, 1,
		1, 1, 1, 1, 2, 4, 2, 55, 8, 2, 11, 2, 12, 2, 56, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 5, 1, 5, 5, 5, 67, 8, 5, 10, 5, 12, 5, 70, 9, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 68, 0, 19, 1, 1, 3, 2, 5, 0, 7, 3, 9, 4, 11, 5, 13, 6, 15, 7, 17, 8,
		19, 9, 21, 10, 23, 11, 25, 12, 27, 13, 29, 14, 31, 15, 33, 16, 35, 17,
		37, 18, 1, 0, 3, 2, 0, 9, 9, 32, 32, 2, 0, 10, 10, 13, 13, 1, 0, 48, 57,
		108, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 1, 40, 1, 0, 0, 0,
		3, 47, 1, 0, 0, 0, 5, 54, 1, 0, 0, 0, 7, 58, 1, 0, 0, 0, 9, 60, 1, 0, 0,
		0, 11, 64, 1, 0, 0, 0, 13, 73, 1, 0, 0, 0, 15, 78, 1, 0, 0, 0, 17, 84,
		1, 0, 0, 0, 19, 86, 1, 0, 0, 0, 21, 88, 1, 0, 0, 0, 23, 90, 1, 0, 0, 0,
		25, 92, 1, 0, 0, 0, 27, 94, 1, 0, 0, 0, 29, 96, 1, 0, 0, 0, 31, 98, 1,
		0, 0, 0, 33, 100, 1, 0, 0, 0, 35, 102, 1, 0, 0, 0, 37, 104, 1, 0, 0, 0,
		39, 41, 7, 0, 0, 0, 40, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 40, 1,
		0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 45, 6, 0, 0, 0, 45,
		2, 1, 0, 0, 0, 46, 48, 7, 1, 0, 0, 47, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0,
		0, 49, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 52,
		6, 1, 0, 0, 52, 4, 1, 0, 0, 0, 53, 55, 7, 2, 0, 0, 54, 53, 1, 0, 0, 0,
		55, 56, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 6, 1, 0,
		0, 0, 58, 59, 3, 5, 2, 0, 59, 8, 1, 0, 0, 0, 60, 61, 3, 5, 2, 0, 61, 62,
		3, 31, 15, 0, 62, 63, 3, 5, 2, 0, 63, 10, 1, 0, 0, 0, 64, 68, 3, 17, 8,
		0, 65, 67, 9, 0, 0, 0, 66, 65, 1, 0, 0, 0, 67, 70, 1, 0, 0, 0, 68, 69,
		1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 69, 71, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0,
		71, 72, 3, 17, 8, 0, 72, 12, 1, 0, 0, 0, 73, 74, 5, 116, 0, 0, 74, 75,
		5, 114, 0, 0, 75, 76, 5, 117, 0, 0, 76, 77, 5, 101, 0, 0, 77, 14, 1, 0,
		0, 0, 78, 79, 5, 102, 0, 0, 79, 80, 5, 97, 0, 0, 80, 81, 5, 108, 0, 0,
		81, 82, 5, 115, 0, 0, 82, 83, 5, 101, 0, 0, 83, 16, 1, 0, 0, 0, 84, 85,
		5, 34, 0, 0, 85, 18, 1, 0, 0, 0, 86, 87, 5, 40, 0, 0, 87, 20, 1, 0, 0,
		0, 88, 89, 5, 41, 0, 0, 89, 22, 1, 0, 0, 0, 90, 91, 5, 91, 0, 0, 91, 24,
		1, 0, 0, 0, 92, 93, 5, 93, 0, 0, 93, 26, 1, 0, 0, 0, 94, 95, 5, 123, 0,
		0, 95, 28, 1, 0, 0, 0, 96, 97, 5, 125, 0, 0, 97, 30, 1, 0, 0, 0, 98, 99,
		5, 46, 0, 0, 99, 32, 1, 0, 0, 0, 100, 101, 5, 44, 0, 0, 101, 34, 1, 0,
		0, 0, 102, 103, 5, 58, 0, 0, 103, 36, 1, 0, 0, 0, 104, 105, 5, 59, 0, 0,
		105, 38, 1, 0, 0, 0, 5, 0, 42, 49, 56, 68, 1, 6, 0, 0,
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

// GolflangTokensInit initializes any static state used to implement GolflangTokens. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewGolflangTokens(). You can call this function if you wish to initialize the static state ahead
// of time.
func GolflangTokensInit() {
	staticData := &GolflangTokensLexerStaticData
	staticData.once.Do(golflangtokensLexerInit)
}

// NewGolflangTokens produces a new lexer instance for the optional input antlr.CharStream.
func NewGolflangTokens(input antlr.CharStream) *GolflangTokens {
	GolflangTokensInit()
	l := new(GolflangTokens)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &GolflangTokensLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "GolflangTokens.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GolflangTokens tokens.
const (
	GolflangTokensWHITESPACE = 1
	GolflangTokensNEWLINE    = 2
	GolflangTokensINT        = 3
	GolflangTokensDEC        = 4
	GolflangTokensSTR        = 5
	GolflangTokensTRUE       = 6
	GolflangTokensFALSE      = 7
	GolflangTokensQUOTE      = 8
	GolflangTokensLPAREN     = 9
	GolflangTokensRPAREN     = 10
	GolflangTokensLBRACKET   = 11
	GolflangTokensRBRACKET   = 12
	GolflangTokensLBRACE     = 13
	GolflangTokensRBRACE     = 14
	GolflangTokensDOT        = 15
	GolflangTokensCOMMA      = 16
	GolflangTokensCOLON      = 17
	GolflangTokensSEMICOLON  = 18
)
