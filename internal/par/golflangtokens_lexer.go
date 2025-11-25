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
		"'if'", "'then'", "'else'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "NEWLINE", "COMMENT", "INT", "DEC", "STR", "TRUE",
		"FALSE", "QUOTE", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "LBRACE",
		"RBRACE", "DOT", "COMMA", "COLON", "SEMICOLON", "ASSIGN", "BACKSLASH",
		"ARROW", "PLUS", "MINUS", "STAR", "SLASH", "DOUBLESTAR", "UNDERSCORE",
		"GT", "LT", "EQ", "GE", "LE", "NE", "DOUBLELT", "DOUBLEGTL", "DOUBLEGTA",
		"AMPERSAND", "PIPE", "CARET", "IF", "THEN", "ELSE", "IDENT",
	}
	staticData.RuleNames = []string{
		"WHITESPACE", "NEWLINE", "COMMENT", "DIGITS", "INT", "DEC", "STR", "TRUE",
		"FALSE", "QUOTE", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "LBRACE",
		"RBRACE", "DOT", "COMMA", "COLON", "SEMICOLON", "ASSIGN", "BACKSLASH",
		"ARROW", "PLUS", "MINUS", "STAR", "SLASH", "DOUBLESTAR", "UNDERSCORE",
		"GT", "LT", "EQ", "GE", "LE", "NE", "DOUBLELT", "DOUBLEGTL", "DOUBLEGTA",
		"AMPERSAND", "PIPE", "CARET", "IF", "THEN", "ELSE", "IDENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 44, 246, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 1, 0, 4, 0, 93, 8, 0, 11,
		0, 12, 0, 94, 1, 0, 1, 0, 1, 1, 4, 1, 100, 8, 1, 11, 1, 12, 1, 101, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 110, 8, 2, 10, 2, 12, 2, 113, 9,
		2, 1, 2, 1, 2, 3, 2, 117, 8, 2, 1, 2, 1, 2, 1, 3, 4, 3, 122, 8, 3, 11,
		3, 12, 3, 123, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 5, 6, 134,
		8, 6, 10, 6, 12, 6, 137, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26,
		1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1,
		31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34,
		1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 42,
		1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1,
		44, 5, 44, 242, 8, 44, 10, 44, 12, 44, 245, 9, 44, 2, 111, 135, 0, 45,
		1, 1, 3, 2, 5, 3, 7, 0, 9, 4, 11, 5, 13, 6, 15, 7, 17, 8, 19, 9, 21, 10,
		23, 11, 25, 12, 27, 13, 29, 14, 31, 15, 33, 16, 35, 17, 37, 18, 39, 19,
		41, 20, 43, 21, 45, 22, 47, 23, 49, 24, 51, 25, 53, 26, 55, 27, 57, 28,
		59, 29, 61, 30, 63, 31, 65, 32, 67, 33, 69, 34, 71, 35, 73, 36, 75, 37,
		77, 38, 79, 39, 81, 40, 83, 41, 85, 42, 87, 43, 89, 44, 1, 0, 5, 3, 0,
		9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 1, 0, 48, 57, 2, 0, 65, 90,
		97, 122, 3, 0, 48, 57, 65, 90, 97, 122, 251, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1,
		0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59,
		1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0,
		67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0,
		0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0,
		0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0,
		0, 0, 1, 92, 1, 0, 0, 0, 3, 99, 1, 0, 0, 0, 5, 105, 1, 0, 0, 0, 7, 121,
		1, 0, 0, 0, 9, 125, 1, 0, 0, 0, 11, 127, 1, 0, 0, 0, 13, 131, 1, 0, 0,
		0, 15, 140, 1, 0, 0, 0, 17, 145, 1, 0, 0, 0, 19, 151, 1, 0, 0, 0, 21, 153,
		1, 0, 0, 0, 23, 155, 1, 0, 0, 0, 25, 157, 1, 0, 0, 0, 27, 159, 1, 0, 0,
		0, 29, 161, 1, 0, 0, 0, 31, 163, 1, 0, 0, 0, 33, 165, 1, 0, 0, 0, 35, 167,
		1, 0, 0, 0, 37, 169, 1, 0, 0, 0, 39, 171, 1, 0, 0, 0, 41, 173, 1, 0, 0,
		0, 43, 176, 1, 0, 0, 0, 45, 178, 1, 0, 0, 0, 47, 181, 1, 0, 0, 0, 49, 183,
		1, 0, 0, 0, 51, 185, 1, 0, 0, 0, 53, 187, 1, 0, 0, 0, 55, 189, 1, 0, 0,
		0, 57, 192, 1, 0, 0, 0, 59, 194, 1, 0, 0, 0, 61, 196, 1, 0, 0, 0, 63, 198,
		1, 0, 0, 0, 65, 200, 1, 0, 0, 0, 67, 203, 1, 0, 0, 0, 69, 206, 1, 0, 0,
		0, 71, 209, 1, 0, 0, 0, 73, 212, 1, 0, 0, 0, 75, 216, 1, 0, 0, 0, 77, 220,
		1, 0, 0, 0, 79, 222, 1, 0, 0, 0, 81, 224, 1, 0, 0, 0, 83, 226, 1, 0, 0,
		0, 85, 229, 1, 0, 0, 0, 87, 234, 1, 0, 0, 0, 89, 239, 1, 0, 0, 0, 91, 93,
		7, 0, 0, 0, 92, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0,
		94, 95, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 97, 6, 0, 0, 0, 97, 2, 1, 0,
		0, 0, 98, 100, 7, 1, 0, 0, 99, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101,
		99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 104, 6,
		1, 0, 0, 104, 4, 1, 0, 0, 0, 105, 106, 5, 47, 0, 0, 106, 107, 5, 47, 0,
		0, 107, 111, 1, 0, 0, 0, 108, 110, 9, 0, 0, 0, 109, 108, 1, 0, 0, 0, 110,
		113, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112, 116,
		1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 114, 117, 3, 3, 1, 0, 115, 117, 5, 0,
		0, 1, 116, 114, 1, 0, 0, 0, 116, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0,
		118, 119, 6, 2, 0, 0, 119, 6, 1, 0, 0, 0, 120, 122, 7, 2, 0, 0, 121, 120,
		1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0,
		0, 0, 124, 8, 1, 0, 0, 0, 125, 126, 3, 7, 3, 0, 126, 10, 1, 0, 0, 0, 127,
		128, 3, 7, 3, 0, 128, 129, 3, 33, 16, 0, 129, 130, 3, 7, 3, 0, 130, 12,
		1, 0, 0, 0, 131, 135, 3, 19, 9, 0, 132, 134, 9, 0, 0, 0, 133, 132, 1, 0,
		0, 0, 134, 137, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0,
		136, 138, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 138, 139, 3, 19, 9, 0, 139,
		14, 1, 0, 0, 0, 140, 141, 5, 116, 0, 0, 141, 142, 5, 114, 0, 0, 142, 143,
		5, 117, 0, 0, 143, 144, 5, 101, 0, 0, 144, 16, 1, 0, 0, 0, 145, 146, 5,
		102, 0, 0, 146, 147, 5, 97, 0, 0, 147, 148, 5, 108, 0, 0, 148, 149, 5,
		115, 0, 0, 149, 150, 5, 101, 0, 0, 150, 18, 1, 0, 0, 0, 151, 152, 5, 34,
		0, 0, 152, 20, 1, 0, 0, 0, 153, 154, 5, 40, 0, 0, 154, 22, 1, 0, 0, 0,
		155, 156, 5, 41, 0, 0, 156, 24, 1, 0, 0, 0, 157, 158, 5, 91, 0, 0, 158,
		26, 1, 0, 0, 0, 159, 160, 5, 93, 0, 0, 160, 28, 1, 0, 0, 0, 161, 162, 5,
		123, 0, 0, 162, 30, 1, 0, 0, 0, 163, 164, 5, 125, 0, 0, 164, 32, 1, 0,
		0, 0, 165, 166, 5, 46, 0, 0, 166, 34, 1, 0, 0, 0, 167, 168, 5, 44, 0, 0,
		168, 36, 1, 0, 0, 0, 169, 170, 5, 58, 0, 0, 170, 38, 1, 0, 0, 0, 171, 172,
		5, 59, 0, 0, 172, 40, 1, 0, 0, 0, 173, 174, 5, 58, 0, 0, 174, 175, 5, 61,
		0, 0, 175, 42, 1, 0, 0, 0, 176, 177, 5, 92, 0, 0, 177, 44, 1, 0, 0, 0,
		178, 179, 5, 61, 0, 0, 179, 180, 5, 62, 0, 0, 180, 46, 1, 0, 0, 0, 181,
		182, 5, 43, 0, 0, 182, 48, 1, 0, 0, 0, 183, 184, 5, 45, 0, 0, 184, 50,
		1, 0, 0, 0, 185, 186, 5, 42, 0, 0, 186, 52, 1, 0, 0, 0, 187, 188, 5, 47,
		0, 0, 188, 54, 1, 0, 0, 0, 189, 190, 5, 42, 0, 0, 190, 191, 5, 42, 0, 0,
		191, 56, 1, 0, 0, 0, 192, 193, 5, 95, 0, 0, 193, 58, 1, 0, 0, 0, 194, 195,
		5, 62, 0, 0, 195, 60, 1, 0, 0, 0, 196, 197, 5, 60, 0, 0, 197, 62, 1, 0,
		0, 0, 198, 199, 5, 61, 0, 0, 199, 64, 1, 0, 0, 0, 200, 201, 5, 62, 0, 0,
		201, 202, 5, 61, 0, 0, 202, 66, 1, 0, 0, 0, 203, 204, 5, 60, 0, 0, 204,
		205, 5, 61, 0, 0, 205, 68, 1, 0, 0, 0, 206, 207, 5, 33, 0, 0, 207, 208,
		5, 61, 0, 0, 208, 70, 1, 0, 0, 0, 209, 210, 5, 60, 0, 0, 210, 211, 5, 60,
		0, 0, 211, 72, 1, 0, 0, 0, 212, 213, 5, 108, 0, 0, 213, 214, 5, 62, 0,
		0, 214, 215, 5, 62, 0, 0, 215, 74, 1, 0, 0, 0, 216, 217, 5, 97, 0, 0, 217,
		218, 5, 62, 0, 0, 218, 219, 5, 62, 0, 0, 219, 76, 1, 0, 0, 0, 220, 221,
		5, 38, 0, 0, 221, 78, 1, 0, 0, 0, 222, 223, 5, 124, 0, 0, 223, 80, 1, 0,
		0, 0, 224, 225, 5, 94, 0, 0, 225, 82, 1, 0, 0, 0, 226, 227, 5, 105, 0,
		0, 227, 228, 5, 102, 0, 0, 228, 84, 1, 0, 0, 0, 229, 230, 5, 116, 0, 0,
		230, 231, 5, 104, 0, 0, 231, 232, 5, 101, 0, 0, 232, 233, 5, 110, 0, 0,
		233, 86, 1, 0, 0, 0, 234, 235, 5, 101, 0, 0, 235, 236, 5, 108, 0, 0, 236,
		237, 5, 115, 0, 0, 237, 238, 5, 101, 0, 0, 238, 88, 1, 0, 0, 0, 239, 243,
		7, 3, 0, 0, 240, 242, 7, 4, 0, 0, 241, 240, 1, 0, 0, 0, 242, 245, 1, 0,
		0, 0, 243, 241, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 90, 1, 0, 0, 0,
		245, 243, 1, 0, 0, 0, 8, 0, 94, 101, 111, 116, 123, 135, 243, 1, 0, 1,
		0,
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
	GolflangTokensIF         = 41
	GolflangTokensTHEN       = 42
	GolflangTokensELSE       = 43
	GolflangTokensIDENT      = 44
)
