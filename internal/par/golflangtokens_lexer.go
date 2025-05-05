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
		"", "", "", "", "", "", "", "'true'", "'false'", "'\"'", "'('", "')'",
		"'['", "']'", "'{'", "'}'", "'.'", "','", "':'", "';'", "':='", "'\\'",
		"'=>'", "'+'", "'-'", "'*'", "'/'", "'**'", "'_'", "'>'", "'<'", "'='",
		"'>='", "'<='", "'!='", "'<<'", "'l>>'", "'a>>'", "'&'", "'|'", "'^'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "NEWLINE", "COMMENT", "INT", "DEC", "STR", "TRUE",
		"FALSE", "QUOTE", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "LBRACE",
		"RBRACE", "DOT", "COMMA", "COLON", "SEMICOLON", "ASSIGN", "BACKSLASH",
		"ARROW", "PLUS", "MINUS", "STAR", "SLASH", "DOUBLESTAR", "UNDERSCORE",
		"GT", "LT", "EQ", "GE", "LE", "NE", "DOUBLELT", "DOUBLEGTL", "DOUBLEGTA",
		"AMPERSAND", "PIPE", "CARET", "IDENT",
	}
	staticData.RuleNames = []string{
		"WHITESPACE", "NEWLINE", "COMMENT", "DIGITS", "INT", "DEC", "STR", "TRUE",
		"FALSE", "QUOTE", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "LBRACE",
		"RBRACE", "DOT", "COMMA", "COLON", "SEMICOLON", "ASSIGN", "BACKSLASH",
		"ARROW", "PLUS", "MINUS", "STAR", "SLASH", "DOUBLESTAR", "UNDERSCORE",
		"GT", "LT", "EQ", "GE", "LE", "NE", "DOUBLELT", "DOUBLEGTL", "DOUBLEGTA",
		"AMPERSAND", "PIPE", "CARET", "IDENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 41, 227, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 1, 0, 4, 0, 87, 8, 0, 11, 0, 12, 0, 88, 1, 0, 1, 0, 1, 1, 4, 1, 94,
		8, 1, 11, 1, 12, 1, 95, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 104,
		8, 2, 10, 2, 12, 2, 107, 9, 2, 1, 2, 1, 2, 3, 2, 111, 8, 2, 1, 2, 1, 2,
		1, 3, 4, 3, 116, 8, 3, 11, 3, 12, 3, 117, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 5, 6, 128, 8, 6, 10, 6, 12, 6, 131, 9, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28,
		1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1,
		33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1,
		40, 1, 41, 1, 41, 5, 41, 223, 8, 41, 10, 41, 12, 41, 226, 9, 41, 2, 105,
		129, 0, 42, 1, 1, 3, 2, 5, 3, 7, 0, 9, 4, 11, 5, 13, 6, 15, 7, 17, 8, 19,
		9, 21, 10, 23, 11, 25, 12, 27, 13, 29, 14, 31, 15, 33, 16, 35, 17, 37,
		18, 39, 19, 41, 20, 43, 21, 45, 22, 47, 23, 49, 24, 51, 25, 53, 26, 55,
		27, 57, 28, 59, 29, 61, 30, 63, 31, 65, 32, 67, 33, 69, 34, 71, 35, 73,
		36, 75, 37, 77, 38, 79, 39, 81, 40, 83, 41, 1, 0, 5, 3, 0, 9, 10, 13, 13,
		32, 32, 2, 0, 10, 10, 13, 13, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 3, 0,
		48, 57, 65, 90, 97, 122, 232, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0,
		0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0,
		0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1,
		0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53,
		1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0,
		61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0,
		0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0,
		0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0,
		0, 0, 1, 86, 1, 0, 0, 0, 3, 93, 1, 0, 0, 0, 5, 99, 1, 0, 0, 0, 7, 115,
		1, 0, 0, 0, 9, 119, 1, 0, 0, 0, 11, 121, 1, 0, 0, 0, 13, 125, 1, 0, 0,
		0, 15, 134, 1, 0, 0, 0, 17, 139, 1, 0, 0, 0, 19, 145, 1, 0, 0, 0, 21, 147,
		1, 0, 0, 0, 23, 149, 1, 0, 0, 0, 25, 151, 1, 0, 0, 0, 27, 153, 1, 0, 0,
		0, 29, 155, 1, 0, 0, 0, 31, 157, 1, 0, 0, 0, 33, 159, 1, 0, 0, 0, 35, 161,
		1, 0, 0, 0, 37, 163, 1, 0, 0, 0, 39, 165, 1, 0, 0, 0, 41, 167, 1, 0, 0,
		0, 43, 170, 1, 0, 0, 0, 45, 172, 1, 0, 0, 0, 47, 175, 1, 0, 0, 0, 49, 177,
		1, 0, 0, 0, 51, 179, 1, 0, 0, 0, 53, 181, 1, 0, 0, 0, 55, 183, 1, 0, 0,
		0, 57, 186, 1, 0, 0, 0, 59, 188, 1, 0, 0, 0, 61, 190, 1, 0, 0, 0, 63, 192,
		1, 0, 0, 0, 65, 194, 1, 0, 0, 0, 67, 197, 1, 0, 0, 0, 69, 200, 1, 0, 0,
		0, 71, 203, 1, 0, 0, 0, 73, 206, 1, 0, 0, 0, 75, 210, 1, 0, 0, 0, 77, 214,
		1, 0, 0, 0, 79, 216, 1, 0, 0, 0, 81, 218, 1, 0, 0, 0, 83, 220, 1, 0, 0,
		0, 85, 87, 7, 0, 0, 0, 86, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 86,
		1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 91, 6, 0, 0, 0,
		91, 2, 1, 0, 0, 0, 92, 94, 7, 1, 0, 0, 93, 92, 1, 0, 0, 0, 94, 95, 1, 0,
		0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 98,
		6, 1, 0, 0, 98, 4, 1, 0, 0, 0, 99, 100, 5, 47, 0, 0, 100, 101, 5, 47, 0,
		0, 101, 105, 1, 0, 0, 0, 102, 104, 9, 0, 0, 0, 103, 102, 1, 0, 0, 0, 104,
		107, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 106, 110,
		1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 108, 111, 3, 3, 1, 0, 109, 111, 5, 0,
		0, 1, 110, 108, 1, 0, 0, 0, 110, 109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0,
		112, 113, 6, 2, 0, 0, 113, 6, 1, 0, 0, 0, 114, 116, 7, 2, 0, 0, 115, 114,
		1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0,
		0, 0, 118, 8, 1, 0, 0, 0, 119, 120, 3, 7, 3, 0, 120, 10, 1, 0, 0, 0, 121,
		122, 3, 7, 3, 0, 122, 123, 3, 33, 16, 0, 123, 124, 3, 7, 3, 0, 124, 12,
		1, 0, 0, 0, 125, 129, 3, 19, 9, 0, 126, 128, 9, 0, 0, 0, 127, 126, 1, 0,
		0, 0, 128, 131, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0,
		130, 132, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 132, 133, 3, 19, 9, 0, 133,
		14, 1, 0, 0, 0, 134, 135, 5, 116, 0, 0, 135, 136, 5, 114, 0, 0, 136, 137,
		5, 117, 0, 0, 137, 138, 5, 101, 0, 0, 138, 16, 1, 0, 0, 0, 139, 140, 5,
		102, 0, 0, 140, 141, 5, 97, 0, 0, 141, 142, 5, 108, 0, 0, 142, 143, 5,
		115, 0, 0, 143, 144, 5, 101, 0, 0, 144, 18, 1, 0, 0, 0, 145, 146, 5, 34,
		0, 0, 146, 20, 1, 0, 0, 0, 147, 148, 5, 40, 0, 0, 148, 22, 1, 0, 0, 0,
		149, 150, 5, 41, 0, 0, 150, 24, 1, 0, 0, 0, 151, 152, 5, 91, 0, 0, 152,
		26, 1, 0, 0, 0, 153, 154, 5, 93, 0, 0, 154, 28, 1, 0, 0, 0, 155, 156, 5,
		123, 0, 0, 156, 30, 1, 0, 0, 0, 157, 158, 5, 125, 0, 0, 158, 32, 1, 0,
		0, 0, 159, 160, 5, 46, 0, 0, 160, 34, 1, 0, 0, 0, 161, 162, 5, 44, 0, 0,
		162, 36, 1, 0, 0, 0, 163, 164, 5, 58, 0, 0, 164, 38, 1, 0, 0, 0, 165, 166,
		5, 59, 0, 0, 166, 40, 1, 0, 0, 0, 167, 168, 5, 58, 0, 0, 168, 169, 5, 61,
		0, 0, 169, 42, 1, 0, 0, 0, 170, 171, 5, 92, 0, 0, 171, 44, 1, 0, 0, 0,
		172, 173, 5, 61, 0, 0, 173, 174, 5, 62, 0, 0, 174, 46, 1, 0, 0, 0, 175,
		176, 5, 43, 0, 0, 176, 48, 1, 0, 0, 0, 177, 178, 5, 45, 0, 0, 178, 50,
		1, 0, 0, 0, 179, 180, 5, 42, 0, 0, 180, 52, 1, 0, 0, 0, 181, 182, 5, 47,
		0, 0, 182, 54, 1, 0, 0, 0, 183, 184, 5, 42, 0, 0, 184, 185, 5, 42, 0, 0,
		185, 56, 1, 0, 0, 0, 186, 187, 5, 95, 0, 0, 187, 58, 1, 0, 0, 0, 188, 189,
		5, 62, 0, 0, 189, 60, 1, 0, 0, 0, 190, 191, 5, 60, 0, 0, 191, 62, 1, 0,
		0, 0, 192, 193, 5, 61, 0, 0, 193, 64, 1, 0, 0, 0, 194, 195, 5, 62, 0, 0,
		195, 196, 5, 61, 0, 0, 196, 66, 1, 0, 0, 0, 197, 198, 5, 60, 0, 0, 198,
		199, 5, 61, 0, 0, 199, 68, 1, 0, 0, 0, 200, 201, 5, 33, 0, 0, 201, 202,
		5, 61, 0, 0, 202, 70, 1, 0, 0, 0, 203, 204, 5, 60, 0, 0, 204, 205, 5, 60,
		0, 0, 205, 72, 1, 0, 0, 0, 206, 207, 5, 108, 0, 0, 207, 208, 5, 62, 0,
		0, 208, 209, 5, 62, 0, 0, 209, 74, 1, 0, 0, 0, 210, 211, 5, 97, 0, 0, 211,
		212, 5, 62, 0, 0, 212, 213, 5, 62, 0, 0, 213, 76, 1, 0, 0, 0, 214, 215,
		5, 38, 0, 0, 215, 78, 1, 0, 0, 0, 216, 217, 5, 124, 0, 0, 217, 80, 1, 0,
		0, 0, 218, 219, 5, 94, 0, 0, 219, 82, 1, 0, 0, 0, 220, 224, 7, 3, 0, 0,
		221, 223, 7, 4, 0, 0, 222, 221, 1, 0, 0, 0, 223, 226, 1, 0, 0, 0, 224,
		222, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 84, 1, 0, 0, 0, 226, 224, 1,
		0, 0, 0, 8, 0, 88, 95, 105, 110, 117, 129, 224, 1, 0, 1, 0,
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
	GolflangTokensCOMMENT    = 3
	GolflangTokensINT        = 4
	GolflangTokensDEC        = 5
	GolflangTokensSTR        = 6
	GolflangTokensTRUE       = 7
	GolflangTokensFALSE      = 8
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
	GolflangTokensBACKSLASH  = 21
	GolflangTokensARROW      = 22
	GolflangTokensPLUS       = 23
	GolflangTokensMINUS      = 24
	GolflangTokensSTAR       = 25
	GolflangTokensSLASH      = 26
	GolflangTokensDOUBLESTAR = 27
	GolflangTokensUNDERSCORE = 28
	GolflangTokensGT         = 29
	GolflangTokensLT         = 30
	GolflangTokensEQ         = 31
	GolflangTokensGE         = 32
	GolflangTokensLE         = 33
	GolflangTokensNE         = 34
	GolflangTokensDOUBLELT   = 35
	GolflangTokensDOUBLEGTL  = 36
	GolflangTokensDOUBLEGTA  = 37
	GolflangTokensAMPERSAND  = 38
	GolflangTokensPIPE       = 39
	GolflangTokensCARET      = 40
	GolflangTokensIDENT      = 41
)
