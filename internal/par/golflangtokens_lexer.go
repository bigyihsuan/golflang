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
		"", "", "", "", "", "", "'true'", "'false'", "", "'\"'", "'('", "')'",
		"'['", "']'", "'{'", "'}'", "'.'", "','", "':'", "';'", "':='",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "NEWLINE", "INT", "DEC", "STR", "TRUE", "FALSE", "IDENT",
		"QUOTE", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "LBRACE", "RBRACE",
		"DOT", "COMMA", "COLON", "SEMICOLON", "ASSIGN",
	}
	staticData.RuleNames = []string{
		"WHITESPACE", "NEWLINE", "DIGITS", "INT", "DEC", "STR", "TRUE", "FALSE",
		"IDENT", "QUOTE", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "LBRACE",
		"RBRACE", "DOT", "COMMA", "COLON", "SEMICOLON", "ASSIGN",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 20, 120, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 1, 0, 4, 0, 45, 8, 0, 11, 0, 12, 0, 46, 1, 0, 1, 0, 1, 1, 4, 1, 52,
		8, 1, 11, 1, 12, 1, 53, 1, 1, 1, 1, 1, 2, 4, 2, 59, 8, 2, 11, 2, 12, 2,
		60, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 5, 5, 71, 8, 5, 10,
		5, 12, 5, 74, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 5, 8, 91, 8, 8, 10, 8, 12, 8, 94,
		9, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 72, 0, 21, 1, 1, 3, 2, 5, 0, 7, 3, 9,
		4, 11, 5, 13, 6, 15, 7, 17, 8, 19, 9, 21, 10, 23, 11, 25, 12, 27, 13, 29,
		14, 31, 15, 33, 16, 35, 17, 37, 18, 39, 19, 41, 20, 1, 0, 5, 2, 0, 9, 9,
		32, 32, 2, 0, 10, 10, 13, 13, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 3, 0,
		48, 57, 65, 90, 97, 122, 123, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0,
		0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0,
		0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 1, 44, 1, 0, 0, 0, 3, 51, 1,
		0, 0, 0, 5, 58, 1, 0, 0, 0, 7, 62, 1, 0, 0, 0, 9, 64, 1, 0, 0, 0, 11, 68,
		1, 0, 0, 0, 13, 77, 1, 0, 0, 0, 15, 82, 1, 0, 0, 0, 17, 88, 1, 0, 0, 0,
		19, 95, 1, 0, 0, 0, 21, 97, 1, 0, 0, 0, 23, 99, 1, 0, 0, 0, 25, 101, 1,
		0, 0, 0, 27, 103, 1, 0, 0, 0, 29, 105, 1, 0, 0, 0, 31, 107, 1, 0, 0, 0,
		33, 109, 1, 0, 0, 0, 35, 111, 1, 0, 0, 0, 37, 113, 1, 0, 0, 0, 39, 115,
		1, 0, 0, 0, 41, 117, 1, 0, 0, 0, 43, 45, 7, 0, 0, 0, 44, 43, 1, 0, 0, 0,
		45, 46, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 48, 1,
		0, 0, 0, 48, 49, 6, 0, 0, 0, 49, 2, 1, 0, 0, 0, 50, 52, 7, 1, 0, 0, 51,
		50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0,
		0, 54, 55, 1, 0, 0, 0, 55, 56, 6, 1, 0, 0, 56, 4, 1, 0, 0, 0, 57, 59, 7,
		2, 0, 0, 58, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 60,
		61, 1, 0, 0, 0, 61, 6, 1, 0, 0, 0, 62, 63, 3, 5, 2, 0, 63, 8, 1, 0, 0,
		0, 64, 65, 3, 5, 2, 0, 65, 66, 3, 33, 16, 0, 66, 67, 3, 5, 2, 0, 67, 10,
		1, 0, 0, 0, 68, 72, 3, 19, 9, 0, 69, 71, 9, 0, 0, 0, 70, 69, 1, 0, 0, 0,
		71, 74, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 73, 75, 1,
		0, 0, 0, 74, 72, 1, 0, 0, 0, 75, 76, 3, 19, 9, 0, 76, 12, 1, 0, 0, 0, 77,
		78, 5, 116, 0, 0, 78, 79, 5, 114, 0, 0, 79, 80, 5, 117, 0, 0, 80, 81, 5,
		101, 0, 0, 81, 14, 1, 0, 0, 0, 82, 83, 5, 102, 0, 0, 83, 84, 5, 97, 0,
		0, 84, 85, 5, 108, 0, 0, 85, 86, 5, 115, 0, 0, 86, 87, 5, 101, 0, 0, 87,
		16, 1, 0, 0, 0, 88, 92, 7, 3, 0, 0, 89, 91, 7, 4, 0, 0, 90, 89, 1, 0, 0,
		0, 91, 94, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 18,
		1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 95, 96, 5, 34, 0, 0, 96, 20, 1, 0, 0, 0,
		97, 98, 5, 40, 0, 0, 98, 22, 1, 0, 0, 0, 99, 100, 5, 41, 0, 0, 100, 24,
		1, 0, 0, 0, 101, 102, 5, 91, 0, 0, 102, 26, 1, 0, 0, 0, 103, 104, 5, 93,
		0, 0, 104, 28, 1, 0, 0, 0, 105, 106, 5, 123, 0, 0, 106, 30, 1, 0, 0, 0,
		107, 108, 5, 125, 0, 0, 108, 32, 1, 0, 0, 0, 109, 110, 5, 46, 0, 0, 110,
		34, 1, 0, 0, 0, 111, 112, 5, 44, 0, 0, 112, 36, 1, 0, 0, 0, 113, 114, 5,
		58, 0, 0, 114, 38, 1, 0, 0, 0, 115, 116, 5, 59, 0, 0, 116, 40, 1, 0, 0,
		0, 117, 118, 5, 58, 0, 0, 118, 119, 5, 61, 0, 0, 119, 42, 1, 0, 0, 0, 6,
		0, 46, 53, 60, 72, 92, 1, 6, 0, 0,
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
	GolflangTokensIDENT      = 8
	GolflangTokensQUOTE      = 9
	GolflangTokensLPAREN     = 10
	GolflangTokensRPAREN     = 11
	GolflangTokensLBRACKET   = 12
	GolflangTokensRBRACKET   = 13
	GolflangTokensLBRACE     = 14
	GolflangTokensRBRACE     = 15
	GolflangTokensDOT        = 16
	GolflangTokensCOMMA      = 17
	GolflangTokensCOLON      = 18
	GolflangTokensSEMICOLON  = 19
	GolflangTokensASSIGN     = 20
)
