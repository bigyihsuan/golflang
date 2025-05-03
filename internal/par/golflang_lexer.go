// Code generated from Golflang.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

type GolflangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var GolflangLexerLexerStaticData struct {
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

func golflanglexerLexerInit() {
	staticData := &GolflangLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "'true'", "'false'", "", "'\"'", "'('",
		"')'", "'['", "']'", "'{'", "'}'", "'.'", "','", "':'", "';'", "':='",
		"'\\'", "'=>'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "NEWLINE", "COMMENT", "INT", "DEC", "STR", "TRUE",
		"FALSE", "IDENT", "QUOTE", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET",
		"LBRACE", "RBRACE", "DOT", "COMMA", "COLON", "SEMICOLON", "ASSIGN",
		"BACKSLASH", "ARROW",
	}
	staticData.RuleNames = []string{
		"DIGITS", "WHITESPACE", "NEWLINE", "COMMENT", "INT", "DEC", "STR", "TRUE",
		"FALSE", "IDENT", "QUOTE", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET",
		"LBRACE", "RBRACE", "DOT", "COMMA", "COLON", "SEMICOLON", "ASSIGN",
		"BACKSLASH", "ARROW",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 23, 146, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 4, 0, 51, 8, 0, 11,
		0, 12, 0, 52, 1, 1, 4, 1, 56, 8, 1, 11, 1, 12, 1, 57, 1, 1, 1, 1, 1, 2,
		4, 2, 63, 8, 2, 11, 2, 12, 2, 64, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5,
		3, 73, 8, 3, 10, 3, 12, 3, 76, 9, 3, 1, 3, 1, 3, 3, 3, 80, 8, 3, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 5, 6, 92, 8, 6, 10,
		6, 12, 6, 95, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 5, 9, 112, 8, 9, 10, 9, 12, 9, 115,
		9, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 2,
		74, 93, 0, 24, 1, 0, 3, 1, 5, 2, 7, 3, 9, 4, 11, 5, 13, 6, 15, 7, 17, 8,
		19, 9, 21, 10, 23, 11, 25, 12, 27, 13, 29, 14, 31, 15, 33, 16, 35, 17,
		37, 18, 39, 19, 41, 20, 43, 21, 45, 22, 47, 23, 1, 0, 5, 1, 0, 48, 57,
		3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 2, 0, 65, 90, 97, 122,
		3, 0, 48, 57, 65, 90, 97, 122, 151, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0,
		0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1,
		0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37,
		1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0,
		45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 1, 50, 1, 0, 0, 0, 3, 55, 1, 0, 0, 0,
		5, 62, 1, 0, 0, 0, 7, 68, 1, 0, 0, 0, 9, 83, 1, 0, 0, 0, 11, 85, 1, 0,
		0, 0, 13, 89, 1, 0, 0, 0, 15, 98, 1, 0, 0, 0, 17, 103, 1, 0, 0, 0, 19,
		109, 1, 0, 0, 0, 21, 116, 1, 0, 0, 0, 23, 118, 1, 0, 0, 0, 25, 120, 1,
		0, 0, 0, 27, 122, 1, 0, 0, 0, 29, 124, 1, 0, 0, 0, 31, 126, 1, 0, 0, 0,
		33, 128, 1, 0, 0, 0, 35, 130, 1, 0, 0, 0, 37, 132, 1, 0, 0, 0, 39, 134,
		1, 0, 0, 0, 41, 136, 1, 0, 0, 0, 43, 138, 1, 0, 0, 0, 45, 141, 1, 0, 0,
		0, 47, 143, 1, 0, 0, 0, 49, 51, 7, 0, 0, 0, 50, 49, 1, 0, 0, 0, 51, 52,
		1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 2, 1, 0, 0, 0,
		54, 56, 7, 1, 0, 0, 55, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 55, 1,
		0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 60, 6, 1, 0, 0, 60,
		4, 1, 0, 0, 0, 61, 63, 7, 2, 0, 0, 62, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0,
		0, 64, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 67,
		6, 2, 0, 0, 67, 6, 1, 0, 0, 0, 68, 69, 5, 47, 0, 0, 69, 70, 5, 47, 0, 0,
		70, 74, 1, 0, 0, 0, 71, 73, 9, 0, 0, 0, 72, 71, 1, 0, 0, 0, 73, 76, 1,
		0, 0, 0, 74, 75, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 75, 79, 1, 0, 0, 0, 76,
		74, 1, 0, 0, 0, 77, 80, 3, 5, 2, 0, 78, 80, 5, 0, 0, 1, 79, 77, 1, 0, 0,
		0, 79, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 82, 6, 3, 0, 0, 82, 8, 1,
		0, 0, 0, 83, 84, 3, 1, 0, 0, 84, 10, 1, 0, 0, 0, 85, 86, 3, 1, 0, 0, 86,
		87, 3, 35, 17, 0, 87, 88, 3, 1, 0, 0, 88, 12, 1, 0, 0, 0, 89, 93, 3, 21,
		10, 0, 90, 92, 9, 0, 0, 0, 91, 90, 1, 0, 0, 0, 92, 95, 1, 0, 0, 0, 93,
		94, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 94, 96, 1, 0, 0, 0, 95, 93, 1, 0, 0,
		0, 96, 97, 3, 21, 10, 0, 97, 14, 1, 0, 0, 0, 98, 99, 5, 116, 0, 0, 99,
		100, 5, 114, 0, 0, 100, 101, 5, 117, 0, 0, 101, 102, 5, 101, 0, 0, 102,
		16, 1, 0, 0, 0, 103, 104, 5, 102, 0, 0, 104, 105, 5, 97, 0, 0, 105, 106,
		5, 108, 0, 0, 106, 107, 5, 115, 0, 0, 107, 108, 5, 101, 0, 0, 108, 18,
		1, 0, 0, 0, 109, 113, 7, 3, 0, 0, 110, 112, 7, 4, 0, 0, 111, 110, 1, 0,
		0, 0, 112, 115, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0,
		114, 20, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 116, 117, 5, 34, 0, 0, 117,
		22, 1, 0, 0, 0, 118, 119, 5, 40, 0, 0, 119, 24, 1, 0, 0, 0, 120, 121, 5,
		41, 0, 0, 121, 26, 1, 0, 0, 0, 122, 123, 5, 91, 0, 0, 123, 28, 1, 0, 0,
		0, 124, 125, 5, 93, 0, 0, 125, 30, 1, 0, 0, 0, 126, 127, 5, 123, 0, 0,
		127, 32, 1, 0, 0, 0, 128, 129, 5, 125, 0, 0, 129, 34, 1, 0, 0, 0, 130,
		131, 5, 46, 0, 0, 131, 36, 1, 0, 0, 0, 132, 133, 5, 44, 0, 0, 133, 38,
		1, 0, 0, 0, 134, 135, 5, 58, 0, 0, 135, 40, 1, 0, 0, 0, 136, 137, 5, 59,
		0, 0, 137, 42, 1, 0, 0, 0, 138, 139, 5, 58, 0, 0, 139, 140, 5, 61, 0, 0,
		140, 44, 1, 0, 0, 0, 141, 142, 5, 92, 0, 0, 142, 46, 1, 0, 0, 0, 143, 144,
		5, 61, 0, 0, 144, 145, 5, 62, 0, 0, 145, 48, 1, 0, 0, 0, 8, 0, 52, 57,
		64, 74, 79, 93, 113, 1, 6, 0, 0,
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

// GolflangLexerInit initializes any static state used to implement GolflangLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewGolflangLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func GolflangLexerInit() {
	staticData := &GolflangLexerLexerStaticData
	staticData.once.Do(golflanglexerLexerInit)
}

// NewGolflangLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewGolflangLexer(input antlr.CharStream) *GolflangLexer {
	GolflangLexerInit()
	l := new(GolflangLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &GolflangLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Golflang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GolflangLexer tokens.
const (
	GolflangLexerWHITESPACE = 1
	GolflangLexerNEWLINE    = 2
	GolflangLexerCOMMENT    = 3
	GolflangLexerINT        = 4
	GolflangLexerDEC        = 5
	GolflangLexerSTR        = 6
	GolflangLexerTRUE       = 7
	GolflangLexerFALSE      = 8
	GolflangLexerIDENT      = 9
	GolflangLexerQUOTE      = 10
	GolflangLexerLPAREN     = 11
	GolflangLexerRPAREN     = 12
	GolflangLexerLBRACKET   = 13
	GolflangLexerRBRACKET   = 14
	GolflangLexerLBRACE     = 15
	GolflangLexerRBRACE     = 16
	GolflangLexerDOT        = 17
	GolflangLexerCOMMA      = 18
	GolflangLexerCOLON      = 19
	GolflangLexerSEMICOLON  = 20
	GolflangLexerASSIGN     = 21
	GolflangLexerBACKSLASH  = 22
	GolflangLexerARROW      = 23
)
