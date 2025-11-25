// Code generated from Golflang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package par // Golflang
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type GolflangParser struct {
	*antlr.BaseParser
}

var GolflangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func golflangParserInit() {
	staticData := &GolflangParserStaticData
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
		"prog", "stmt", "alias", "exprStmt", "exprList", "expr", "lambda", "identList",
		"ifThenElse", "literal", "literalList", "literalMap", "literalMapEntry",
		"literalPrimitive", "ident", "operator",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 44, 130, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 5, 0, 34, 8, 0, 10, 0, 12, 0, 37, 9, 0, 1, 1, 1, 1, 3, 1, 41, 8,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 4, 4, 52, 8, 4,
		11, 4, 12, 4, 53, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 61, 8, 5, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 5, 7, 71, 8, 7, 10, 7, 12, 7, 74,
		9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 3, 9,
		86, 8, 9, 1, 10, 1, 10, 3, 10, 90, 8, 10, 1, 10, 1, 10, 5, 10, 94, 8, 10,
		10, 10, 12, 10, 97, 9, 10, 1, 10, 3, 10, 100, 8, 10, 1, 10, 1, 10, 1, 11,
		1, 11, 3, 11, 106, 8, 11, 1, 11, 1, 11, 5, 11, 110, 8, 11, 10, 11, 12,
		11, 113, 9, 11, 1, 11, 3, 11, 116, 8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 0, 0, 16, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 0, 2, 1, 0, 4,
		8, 1, 0, 23, 40, 129, 0, 35, 1, 0, 0, 0, 2, 40, 1, 0, 0, 0, 4, 44, 1, 0,
		0, 0, 6, 48, 1, 0, 0, 0, 8, 51, 1, 0, 0, 0, 10, 60, 1, 0, 0, 0, 12, 62,
		1, 0, 0, 0, 14, 67, 1, 0, 0, 0, 16, 75, 1, 0, 0, 0, 18, 85, 1, 0, 0, 0,
		20, 87, 1, 0, 0, 0, 22, 103, 1, 0, 0, 0, 24, 119, 1, 0, 0, 0, 26, 123,
		1, 0, 0, 0, 28, 125, 1, 0, 0, 0, 30, 127, 1, 0, 0, 0, 32, 34, 3, 2, 1,
		0, 33, 32, 1, 0, 0, 0, 34, 37, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 36,
		1, 0, 0, 0, 36, 1, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 38, 41, 3, 4, 2, 0,
		39, 41, 3, 6, 3, 0, 40, 38, 1, 0, 0, 0, 40, 39, 1, 0, 0, 0, 41, 42, 1,
		0, 0, 0, 42, 43, 5, 19, 0, 0, 43, 3, 1, 0, 0, 0, 44, 45, 3, 28, 14, 0,
		45, 46, 5, 20, 0, 0, 46, 47, 3, 8, 4, 0, 47, 5, 1, 0, 0, 0, 48, 49, 3,
		8, 4, 0, 49, 7, 1, 0, 0, 0, 50, 52, 3, 10, 5, 0, 51, 50, 1, 0, 0, 0, 52,
		53, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 9, 1, 0, 0,
		0, 55, 61, 3, 16, 8, 0, 56, 61, 3, 28, 14, 0, 57, 61, 3, 30, 15, 0, 58,
		61, 3, 12, 6, 0, 59, 61, 3, 18, 9, 0, 60, 55, 1, 0, 0, 0, 60, 56, 1, 0,
		0, 0, 60, 57, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 60, 59, 1, 0, 0, 0, 61, 11,
		1, 0, 0, 0, 62, 63, 5, 21, 0, 0, 63, 64, 3, 14, 7, 0, 64, 65, 5, 22, 0,
		0, 65, 66, 3, 8, 4, 0, 66, 13, 1, 0, 0, 0, 67, 72, 3, 28, 14, 0, 68, 69,
		5, 17, 0, 0, 69, 71, 3, 28, 14, 0, 70, 68, 1, 0, 0, 0, 71, 74, 1, 0, 0,
		0, 72, 70, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 15, 1, 0, 0, 0, 74, 72,
		1, 0, 0, 0, 75, 76, 5, 41, 0, 0, 76, 77, 3, 8, 4, 0, 77, 78, 5, 42, 0,
		0, 78, 79, 3, 8, 4, 0, 79, 80, 5, 43, 0, 0, 80, 81, 3, 8, 4, 0, 81, 17,
		1, 0, 0, 0, 82, 86, 3, 20, 10, 0, 83, 86, 3, 22, 11, 0, 84, 86, 3, 26,
		13, 0, 85, 82, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 84, 1, 0, 0, 0, 86,
		19, 1, 0, 0, 0, 87, 89, 5, 12, 0, 0, 88, 90, 3, 10, 5, 0, 89, 88, 1, 0,
		0, 0, 89, 90, 1, 0, 0, 0, 90, 95, 1, 0, 0, 0, 91, 92, 5, 17, 0, 0, 92,
		94, 3, 10, 5, 0, 93, 91, 1, 0, 0, 0, 94, 97, 1, 0, 0, 0, 95, 93, 1, 0,
		0, 0, 95, 96, 1, 0, 0, 0, 96, 99, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 98, 100,
		5, 17, 0, 0, 99, 98, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 101, 1, 0, 0,
		0, 101, 102, 5, 13, 0, 0, 102, 21, 1, 0, 0, 0, 103, 105, 5, 14, 0, 0, 104,
		106, 3, 24, 12, 0, 105, 104, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 111,
		1, 0, 0, 0, 107, 108, 5, 17, 0, 0, 108, 110, 3, 24, 12, 0, 109, 107, 1,
		0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 111, 112, 1, 0, 0,
		0, 112, 115, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 114, 116, 5, 17, 0, 0, 115,
		114, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 118,
		5, 15, 0, 0, 118, 23, 1, 0, 0, 0, 119, 120, 3, 10, 5, 0, 120, 121, 5, 18,
		0, 0, 121, 122, 3, 10, 5, 0, 122, 25, 1, 0, 0, 0, 123, 124, 7, 0, 0, 0,
		124, 27, 1, 0, 0, 0, 125, 126, 5, 44, 0, 0, 126, 29, 1, 0, 0, 0, 127, 128,
		7, 1, 0, 0, 128, 31, 1, 0, 0, 0, 12, 35, 40, 53, 60, 72, 85, 89, 95, 99,
		105, 111, 115,
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

// GolflangParserInit initializes any static state used to implement GolflangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGolflangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GolflangParserInit() {
	staticData := &GolflangParserStaticData
	staticData.once.Do(golflangParserInit)
}

// NewGolflangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGolflangParser(input antlr.TokenStream) *GolflangParser {
	GolflangParserInit()
	this := new(GolflangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GolflangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Golflang.g4"

	return this
}

// GolflangParser tokens.
const (
	GolflangParserEOF        = antlr.TokenEOF
	GolflangParserWHITESPACE = 1
	GolflangParserNEWLINE    = 2
	GolflangParserCOMMENT    = 3
	GolflangParserINT        = 4
	GolflangParserDEC        = 5
	GolflangParserSTR        = 6
	GolflangParserTRUE       = 7
	GolflangParserFALSE      = 8
	GolflangParserQUOTE      = 9
	GolflangParserLPAREN     = 10
	GolflangParserRPAREN     = 11
	GolflangParserLBRACKET   = 12
	GolflangParserRBRACKET   = 13
	GolflangParserLBRACE     = 14
	GolflangParserRBRACE     = 15
	GolflangParserDOT        = 16
	GolflangParserCOMMA      = 17
	GolflangParserCOLON      = 18
	GolflangParserSEMICOLON  = 19
	GolflangParserASSIGN     = 20
	GolflangParserBACKSLASH  = 21
	GolflangParserARROW      = 22
	GolflangParserPLUS       = 23
	GolflangParserMINUS      = 24
	GolflangParserSTAR       = 25
	GolflangParserSLASH      = 26
	GolflangParserDOUBLESTAR = 27
	GolflangParserUNDERSCORE = 28
	GolflangParserGT         = 29
	GolflangParserLT         = 30
	GolflangParserEQ         = 31
	GolflangParserGE         = 32
	GolflangParserLE         = 33
	GolflangParserNE         = 34
	GolflangParserDOUBLELT   = 35
	GolflangParserDOUBLEGTL  = 36
	GolflangParserDOUBLEGTA  = 37
	GolflangParserAMPERSAND  = 38
	GolflangParserPIPE       = 39
	GolflangParserCARET      = 40
	GolflangParserIF         = 41
	GolflangParserTHEN       = 42
	GolflangParserELSE       = 43
	GolflangParserIDENT      = 44
)

// GolflangParser rules.
const (
	GolflangParserRULE_prog             = 0
	GolflangParserRULE_stmt             = 1
	GolflangParserRULE_alias            = 2
	GolflangParserRULE_exprStmt         = 3
	GolflangParserRULE_exprList         = 4
	GolflangParserRULE_expr             = 5
	GolflangParserRULE_lambda           = 6
	GolflangParserRULE_identList        = 7
	GolflangParserRULE_ifThenElse       = 8
	GolflangParserRULE_literal          = 9
	GolflangParserRULE_literalList      = 10
	GolflangParserRULE_literalMap       = 11
	GolflangParserRULE_literalMapEntry  = 12
	GolflangParserRULE_literalPrimitive = 13
	GolflangParserRULE_ident            = 14
	GolflangParserRULE_operator         = 15
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GolflangParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GolflangVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GolflangParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GolflangParserRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&21990226285040) != 0 {
		{
			p.SetState(32)
			p.Stmt()
		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEMICOLON() antlr.TerminalNode
	Alias() IAliasContext
	ExprStmt() IExprStmtContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GolflangParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GolflangParserSEMICOLON, 0)
}

func (s *StmtContext) Alias() IAliasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *StmtContext) ExprStmt() IExprStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprStmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GolflangVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GolflangParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GolflangParserRULE_stmt)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(38)
			p.Alias()
		}

	case 2:
		{
			p.SetState(39)
			p.ExprStmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(42)
		p.Match(GolflangParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAliasContext is an interface to support dynamic dispatch.
type IAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name rule contexts.
	GetName() IIdentContext

	// GetValue returns the value rule contexts.
	GetValue() IExprListContext

	// SetName sets the name rule contexts.
	SetName(IIdentContext)

	// SetValue sets the value rule contexts.
	SetValue(IExprListContext)

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	Ident() IIdentContext
	ExprList() IExprListContext

	// IsAliasContext differentiates from other interfaces.
	IsAliasContext()
}

type AliasContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   IIdentContext
	value  IExprListContext
}

func NewEmptyAliasContext() *AliasContext {
	var p = new(AliasContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_alias
	return p
}

func InitEmptyAliasContext(p *AliasContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_alias
}

func (*AliasContext) IsAliasContext() {}

func NewAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasContext {
	var p = new(AliasContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GolflangParserRULE_alias

	return p
}

func (s *AliasContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasContext) GetName() IIdentContext { return s.name }

func (s *AliasContext) GetValue() IExprListContext { return s.value }

func (s *AliasContext) SetName(v IIdentContext) { s.name = v }

func (s *AliasContext) SetValue(v IExprListContext) { s.value = v }

func (s *AliasContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GolflangParserASSIGN, 0)
}

func (s *AliasContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *AliasContext) ExprList() IExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *AliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.EnterAlias(s)
	}
}

func (s *AliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.ExitAlias(s)
	}
}

func (s *AliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GolflangVisitor:
		return t.VisitAlias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GolflangParser) Alias() (localctx IAliasContext) {
	localctx = NewAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GolflangParserRULE_alias)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)

		var _x = p.Ident()

		localctx.(*AliasContext).name = _x
	}
	{
		p.SetState(45)
		p.Match(GolflangParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)

		var _x = p.ExprList()

		localctx.(*AliasContext).value = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprStmtContext is an interface to support dynamic dispatch.
type IExprStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExprList() IExprListContext

	// IsExprStmtContext differentiates from other interfaces.
	IsExprStmtContext()
}

type ExprStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprStmtContext() *ExprStmtContext {
	var p = new(ExprStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_exprStmt
	return p
}

func InitEmptyExprStmtContext(p *ExprStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_exprStmt
}

func (*ExprStmtContext) IsExprStmtContext() {}

func NewExprStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprStmtContext {
	var p = new(ExprStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GolflangParserRULE_exprStmt

	return p
}

func (s *ExprStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprStmtContext) ExprList() IExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *ExprStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.EnterExprStmt(s)
	}
}

func (s *ExprStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.ExitExprStmt(s)
	}
}

func (s *ExprStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GolflangVisitor:
		return t.VisitExprStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GolflangParser) ExprStmt() (localctx IExprStmtContext) {
	localctx = NewExprStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GolflangParserRULE_exprStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.ExprList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprListContext is an interface to support dynamic dispatch.
type IExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsExprListContext differentiates from other interfaces.
	IsExprListContext()
}

type ExprListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprListContext() *ExprListContext {
	var p = new(ExprListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_exprList
	return p
}

func InitEmptyExprListContext(p *ExprListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_exprList
}

func (*ExprListContext) IsExprListContext() {}

func NewExprListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprListContext {
	var p = new(ExprListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GolflangParserRULE_exprList

	return p
}

func (s *ExprListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprListContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprListContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.EnterExprList(s)
	}
}

func (s *ExprListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.ExitExprList(s)
	}
}

func (s *ExprListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GolflangVisitor:
		return t.VisitExprList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GolflangParser) ExprList() (localctx IExprListContext) {
	localctx = NewExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GolflangParserRULE_exprList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(50)
				p.Expr()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IfThenElse() IIfThenElseContext
	Ident() IIdentContext
	Operator() IOperatorContext
	Lambda() ILambdaContext
	Literal() ILiteralContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GolflangParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) IfThenElse() IIfThenElseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfThenElseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfThenElseContext)
}

func (s *ExprContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ExprContext) Operator() IOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *ExprContext) Lambda() ILambdaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILambdaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILambdaContext)
}

func (s *ExprContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GolflangVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GolflangParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GolflangParserRULE_expr)
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GolflangParserIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.IfThenElse()
		}

	case GolflangParserIDENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.Ident()
		}

	case GolflangParserPLUS, GolflangParserMINUS, GolflangParserSTAR, GolflangParserSLASH, GolflangParserDOUBLESTAR, GolflangParserUNDERSCORE, GolflangParserGT, GolflangParserLT, GolflangParserEQ, GolflangParserGE, GolflangParserLE, GolflangParserNE, GolflangParserDOUBLELT, GolflangParserDOUBLEGTL, GolflangParserDOUBLEGTA, GolflangParserAMPERSAND, GolflangParserPIPE, GolflangParserCARET:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(57)
			p.Operator()
		}

	case GolflangParserBACKSLASH:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(58)
			p.Lambda()
		}

	case GolflangParserINT, GolflangParserDEC, GolflangParserSTR, GolflangParserTRUE, GolflangParserFALSE, GolflangParserLBRACKET, GolflangParserLBRACE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(59)
			p.Literal()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILambdaContext is an interface to support dynamic dispatch.
type ILambdaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetArgs returns the args rule contexts.
	GetArgs() IIdentListContext

	// GetBody returns the body rule contexts.
	GetBody() IExprListContext

	// SetArgs sets the args rule contexts.
	SetArgs(IIdentListContext)

	// SetBody sets the body rule contexts.
	SetBody(IExprListContext)

	// Getter signatures
	BACKSLASH() antlr.TerminalNode
	ARROW() antlr.TerminalNode
	IdentList() IIdentListContext
	ExprList() IExprListContext

	// IsLambdaContext differentiates from other interfaces.
	IsLambdaContext()
}

type LambdaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	args   IIdentListContext
	body   IExprListContext
}

func NewEmptyLambdaContext() *LambdaContext {
	var p = new(LambdaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_lambda
	return p
}

func InitEmptyLambdaContext(p *LambdaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_lambda
}

func (*LambdaContext) IsLambdaContext() {}

func NewLambdaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaContext {
	var p = new(LambdaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GolflangParserRULE_lambda

	return p
}

func (s *LambdaContext) GetParser() antlr.Parser { return s.parser }

func (s *LambdaContext) GetArgs() IIdentListContext { return s.args }

func (s *LambdaContext) GetBody() IExprListContext { return s.body }

func (s *LambdaContext) SetArgs(v IIdentListContext) { s.args = v }

func (s *LambdaContext) SetBody(v IExprListContext) { s.body = v }

func (s *LambdaContext) BACKSLASH() antlr.TerminalNode {
	return s.GetToken(GolflangParserBACKSLASH, 0)
}

func (s *LambdaContext) ARROW() antlr.TerminalNode {
	return s.GetToken(GolflangParserARROW, 0)
}

func (s *LambdaContext) IdentList() IIdentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentListContext)
}

func (s *LambdaContext) ExprList() IExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *LambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LambdaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.EnterLambda(s)
	}
}

func (s *LambdaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.ExitLambda(s)
	}
}

func (s *LambdaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GolflangVisitor:
		return t.VisitLambda(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GolflangParser) Lambda() (localctx ILambdaContext) {
	localctx = NewLambdaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GolflangParserRULE_lambda)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(GolflangParserBACKSLASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(63)

		var _x = p.IdentList()

		localctx.(*LambdaContext).args = _x
	}
	{
		p.SetState(64)
		p.Match(GolflangParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(65)

		var _x = p.ExprList()

		localctx.(*LambdaContext).body = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentListContext is an interface to support dynamic dispatch.
type IIdentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdent() []IIdentContext
	Ident(i int) IIdentContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsIdentListContext differentiates from other interfaces.
	IsIdentListContext()
}

type IdentListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentListContext() *IdentListContext {
	var p = new(IdentListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_identList
	return p
}

func InitEmptyIdentListContext(p *IdentListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_identList
}

func (*IdentListContext) IsIdentListContext() {}

func NewIdentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentListContext {
	var p = new(IdentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GolflangParserRULE_identList

	return p
}

func (s *IdentListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentListContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *IdentListContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
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

	return t.(IIdentContext)
}

func (s *IdentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GolflangParserCOMMA)
}

func (s *IdentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GolflangParserCOMMA, i)
}

func (s *IdentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.EnterIdentList(s)
	}
}

func (s *IdentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.ExitIdentList(s)
	}
}

func (s *IdentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GolflangVisitor:
		return t.VisitIdentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GolflangParser) IdentList() (localctx IIdentListContext) {
	localctx = NewIdentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GolflangParserRULE_identList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Ident()
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GolflangParserCOMMA {
		{
			p.SetState(68)
			p.Match(GolflangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.Ident()
		}

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfThenElseContext is an interface to support dynamic dispatch.
type IIfThenElseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCond returns the cond rule contexts.
	GetCond() IExprListContext

	// GetThen returns the then rule contexts.
	GetThen() IExprListContext

	// GetElse_ returns the else_ rule contexts.
	GetElse_() IExprListContext

	// SetCond sets the cond rule contexts.
	SetCond(IExprListContext)

	// SetThen sets the then rule contexts.
	SetThen(IExprListContext)

	// SetElse_ sets the else_ rule contexts.
	SetElse_(IExprListContext)

	// Getter signatures
	IF() antlr.TerminalNode
	THEN() antlr.TerminalNode
	ELSE() antlr.TerminalNode
	AllExprList() []IExprListContext
	ExprList(i int) IExprListContext

	// IsIfThenElseContext differentiates from other interfaces.
	IsIfThenElseContext()
}

type IfThenElseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	cond   IExprListContext
	then   IExprListContext
	else_  IExprListContext
}

func NewEmptyIfThenElseContext() *IfThenElseContext {
	var p = new(IfThenElseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_ifThenElse
	return p
}

func InitEmptyIfThenElseContext(p *IfThenElseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_ifThenElse
}

func (*IfThenElseContext) IsIfThenElseContext() {}

func NewIfThenElseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfThenElseContext {
	var p = new(IfThenElseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GolflangParserRULE_ifThenElse

	return p
}

func (s *IfThenElseContext) GetParser() antlr.Parser { return s.parser }

func (s *IfThenElseContext) GetCond() IExprListContext { return s.cond }

func (s *IfThenElseContext) GetThen() IExprListContext { return s.then }

func (s *IfThenElseContext) GetElse_() IExprListContext { return s.else_ }

func (s *IfThenElseContext) SetCond(v IExprListContext) { s.cond = v }

func (s *IfThenElseContext) SetThen(v IExprListContext) { s.then = v }

func (s *IfThenElseContext) SetElse_(v IExprListContext) { s.else_ = v }

func (s *IfThenElseContext) IF() antlr.TerminalNode {
	return s.GetToken(GolflangParserIF, 0)
}

func (s *IfThenElseContext) THEN() antlr.TerminalNode {
	return s.GetToken(GolflangParserTHEN, 0)
}

func (s *IfThenElseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GolflangParserELSE, 0)
}

func (s *IfThenElseContext) AllExprList() []IExprListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprListContext); ok {
			len++
		}
	}

	tst := make([]IExprListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprListContext); ok {
			tst[i] = t.(IExprListContext)
			i++
		}
	}

	return tst
}

func (s *IfThenElseContext) ExprList(i int) IExprListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
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

	return t.(IExprListContext)
}

func (s *IfThenElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfThenElseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfThenElseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.EnterIfThenElse(s)
	}
}

func (s *IfThenElseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.ExitIfThenElse(s)
	}
}

func (s *IfThenElseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GolflangVisitor:
		return t.VisitIfThenElse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GolflangParser) IfThenElse() (localctx IIfThenElseContext) {
	localctx = NewIfThenElseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GolflangParserRULE_ifThenElse)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(GolflangParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)

		var _x = p.ExprList()

		localctx.(*IfThenElseContext).cond = _x
	}
	{
		p.SetState(77)
		p.Match(GolflangParserTHEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)

		var _x = p.ExprList()

		localctx.(*IfThenElseContext).then = _x
	}
	{
		p.SetState(79)
		p.Match(GolflangParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)

		var _x = p.ExprList()

		localctx.(*IfThenElseContext).else_ = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LiteralList() ILiteralListContext
	LiteralMap() ILiteralMapContext
	LiteralPrimitive() ILiteralPrimitiveContext

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GolflangParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) LiteralList() ILiteralListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralListContext)
}

func (s *LiteralContext) LiteralMap() ILiteralMapContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralMapContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralMapContext)
}

func (s *LiteralContext) LiteralPrimitive() ILiteralPrimitiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralPrimitiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralPrimitiveContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GolflangVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GolflangParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GolflangParserRULE_literal)
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GolflangParserLBRACKET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.LiteralList()
		}

	case GolflangParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.LiteralMap()
		}

	case GolflangParserINT, GolflangParserDEC, GolflangParserSTR, GolflangParserTRUE, GolflangParserFALSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(84)
			p.LiteralPrimitive()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralListContext is an interface to support dynamic dispatch.
type ILiteralListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsLiteralListContext differentiates from other interfaces.
	IsLiteralListContext()
}

type LiteralListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralListContext() *LiteralListContext {
	var p = new(LiteralListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_literalList
	return p
}

func InitEmptyLiteralListContext(p *LiteralListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_literalList
}

func (*LiteralListContext) IsLiteralListContext() {}

func NewLiteralListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralListContext {
	var p = new(LiteralListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GolflangParserRULE_literalList

	return p
}

func (s *LiteralListContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralListContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(GolflangParserLBRACKET, 0)
}

func (s *LiteralListContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(GolflangParserRBRACKET, 0)
}

func (s *LiteralListContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *LiteralListContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *LiteralListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GolflangParserCOMMA)
}

func (s *LiteralListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GolflangParserCOMMA, i)
}

func (s *LiteralListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.EnterLiteralList(s)
	}
}

func (s *LiteralListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.ExitLiteralList(s)
	}
}

func (s *LiteralListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GolflangVisitor:
		return t.VisitLiteralList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GolflangParser) LiteralList() (localctx ILiteralListContext) {
	localctx = NewLiteralListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GolflangParserRULE_literalList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		p.Match(GolflangParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&21990226285040) != 0 {
		{
			p.SetState(88)
			p.Expr()
		}

	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(91)
				p.Match(GolflangParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(92)
				p.Expr()
			}

		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GolflangParserCOMMA {
		{
			p.SetState(98)
			p.Match(GolflangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(101)
		p.Match(GolflangParserRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralMapContext is an interface to support dynamic dispatch.
type ILiteralMapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllLiteralMapEntry() []ILiteralMapEntryContext
	LiteralMapEntry(i int) ILiteralMapEntryContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsLiteralMapContext differentiates from other interfaces.
	IsLiteralMapContext()
}

type LiteralMapContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralMapContext() *LiteralMapContext {
	var p = new(LiteralMapContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_literalMap
	return p
}

func InitEmptyLiteralMapContext(p *LiteralMapContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_literalMap
}

func (*LiteralMapContext) IsLiteralMapContext() {}

func NewLiteralMapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralMapContext {
	var p = new(LiteralMapContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GolflangParserRULE_literalMap

	return p
}

func (s *LiteralMapContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralMapContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GolflangParserLBRACE, 0)
}

func (s *LiteralMapContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GolflangParserRBRACE, 0)
}

func (s *LiteralMapContext) AllLiteralMapEntry() []ILiteralMapEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILiteralMapEntryContext); ok {
			len++
		}
	}

	tst := make([]ILiteralMapEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILiteralMapEntryContext); ok {
			tst[i] = t.(ILiteralMapEntryContext)
			i++
		}
	}

	return tst
}

func (s *LiteralMapContext) LiteralMapEntry(i int) ILiteralMapEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralMapEntryContext); ok {
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

	return t.(ILiteralMapEntryContext)
}

func (s *LiteralMapContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GolflangParserCOMMA)
}

func (s *LiteralMapContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GolflangParserCOMMA, i)
}

func (s *LiteralMapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralMapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralMapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.EnterLiteralMap(s)
	}
}

func (s *LiteralMapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.ExitLiteralMap(s)
	}
}

func (s *LiteralMapContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GolflangVisitor:
		return t.VisitLiteralMap(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GolflangParser) LiteralMap() (localctx ILiteralMapContext) {
	localctx = NewLiteralMapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GolflangParserRULE_literalMap)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(GolflangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&21990226285040) != 0 {
		{
			p.SetState(104)
			p.LiteralMapEntry()
		}

	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(107)
				p.Match(GolflangParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(108)
				p.LiteralMapEntry()
			}

		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GolflangParserCOMMA {
		{
			p.SetState(114)
			p.Match(GolflangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(117)
		p.Match(GolflangParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralMapEntryContext is an interface to support dynamic dispatch.
type ILiteralMapEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key rule contexts.
	GetKey() IExprContext

	// GetValue returns the value rule contexts.
	GetValue() IExprContext

	// SetKey sets the key rule contexts.
	SetKey(IExprContext)

	// SetValue sets the value rule contexts.
	SetValue(IExprContext)

	// Getter signatures
	COLON() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsLiteralMapEntryContext differentiates from other interfaces.
	IsLiteralMapEntryContext()
}

type LiteralMapEntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	key    IExprContext
	value  IExprContext
}

func NewEmptyLiteralMapEntryContext() *LiteralMapEntryContext {
	var p = new(LiteralMapEntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_literalMapEntry
	return p
}

func InitEmptyLiteralMapEntryContext(p *LiteralMapEntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_literalMapEntry
}

func (*LiteralMapEntryContext) IsLiteralMapEntryContext() {}

func NewLiteralMapEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralMapEntryContext {
	var p = new(LiteralMapEntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GolflangParserRULE_literalMapEntry

	return p
}

func (s *LiteralMapEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralMapEntryContext) GetKey() IExprContext { return s.key }

func (s *LiteralMapEntryContext) GetValue() IExprContext { return s.value }

func (s *LiteralMapEntryContext) SetKey(v IExprContext) { s.key = v }

func (s *LiteralMapEntryContext) SetValue(v IExprContext) { s.value = v }

func (s *LiteralMapEntryContext) COLON() antlr.TerminalNode {
	return s.GetToken(GolflangParserCOLON, 0)
}

func (s *LiteralMapEntryContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *LiteralMapEntryContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *LiteralMapEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralMapEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralMapEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.EnterLiteralMapEntry(s)
	}
}

func (s *LiteralMapEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.ExitLiteralMapEntry(s)
	}
}

func (s *LiteralMapEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GolflangVisitor:
		return t.VisitLiteralMapEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GolflangParser) LiteralMapEntry() (localctx ILiteralMapEntryContext) {
	localctx = NewLiteralMapEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GolflangParserRULE_literalMapEntry)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)

		var _x = p.Expr()

		localctx.(*LiteralMapEntryContext).key = _x
	}
	{
		p.SetState(120)
		p.Match(GolflangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)

		var _x = p.Expr()

		localctx.(*LiteralMapEntryContext).value = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralPrimitiveContext is an interface to support dynamic dispatch.
type ILiteralPrimitiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	DEC() antlr.TerminalNode
	STR() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode

	// IsLiteralPrimitiveContext differentiates from other interfaces.
	IsLiteralPrimitiveContext()
}

type LiteralPrimitiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralPrimitiveContext() *LiteralPrimitiveContext {
	var p = new(LiteralPrimitiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_literalPrimitive
	return p
}

func InitEmptyLiteralPrimitiveContext(p *LiteralPrimitiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_literalPrimitive
}

func (*LiteralPrimitiveContext) IsLiteralPrimitiveContext() {}

func NewLiteralPrimitiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralPrimitiveContext {
	var p = new(LiteralPrimitiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GolflangParserRULE_literalPrimitive

	return p
}

func (s *LiteralPrimitiveContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralPrimitiveContext) INT() antlr.TerminalNode {
	return s.GetToken(GolflangParserINT, 0)
}

func (s *LiteralPrimitiveContext) DEC() antlr.TerminalNode {
	return s.GetToken(GolflangParserDEC, 0)
}

func (s *LiteralPrimitiveContext) STR() antlr.TerminalNode {
	return s.GetToken(GolflangParserSTR, 0)
}

func (s *LiteralPrimitiveContext) TRUE() antlr.TerminalNode {
	return s.GetToken(GolflangParserTRUE, 0)
}

func (s *LiteralPrimitiveContext) FALSE() antlr.TerminalNode {
	return s.GetToken(GolflangParserFALSE, 0)
}

func (s *LiteralPrimitiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralPrimitiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralPrimitiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.EnterLiteralPrimitive(s)
	}
}

func (s *LiteralPrimitiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.ExitLiteralPrimitive(s)
	}
}

func (s *LiteralPrimitiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GolflangVisitor:
		return t.VisitLiteralPrimitive(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GolflangParser) LiteralPrimitive() (localctx ILiteralPrimitiveContext) {
	localctx = NewLiteralPrimitiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GolflangParserRULE_literalPrimitive)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&496) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentContext is an interface to support dynamic dispatch.
type IIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode

	// IsIdentContext differentiates from other interfaces.
	IsIdentContext()
}

type IdentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentContext() *IdentContext {
	var p = new(IdentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_ident
	return p
}

func InitEmptyIdentContext(p *IdentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_ident
}

func (*IdentContext) IsIdentContext() {}

func NewIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentContext {
	var p = new(IdentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GolflangParserRULE_ident

	return p
}

func (s *IdentContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(GolflangParserIDENT, 0)
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.ExitIdent(s)
	}
}

func (s *IdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GolflangVisitor:
		return t.VisitIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GolflangParser) Ident() (localctx IIdentContext) {
	localctx = NewIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GolflangParserRULE_ident)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(GolflangParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperatorContext is an interface to support dynamic dispatch.
type IOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	STAR() antlr.TerminalNode
	SLASH() antlr.TerminalNode
	DOUBLESTAR() antlr.TerminalNode
	UNDERSCORE() antlr.TerminalNode
	GT() antlr.TerminalNode
	LT() antlr.TerminalNode
	EQ() antlr.TerminalNode
	GE() antlr.TerminalNode
	LE() antlr.TerminalNode
	NE() antlr.TerminalNode
	DOUBLELT() antlr.TerminalNode
	DOUBLEGTL() antlr.TerminalNode
	DOUBLEGTA() antlr.TerminalNode
	AMPERSAND() antlr.TerminalNode
	PIPE() antlr.TerminalNode
	CARET() antlr.TerminalNode

	// IsOperatorContext differentiates from other interfaces.
	IsOperatorContext()
}

type OperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorContext() *OperatorContext {
	var p = new(OperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_operator
	return p
}

func InitEmptyOperatorContext(p *OperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_operator
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GolflangParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorContext) PLUS() antlr.TerminalNode {
	return s.GetToken(GolflangParserPLUS, 0)
}

func (s *OperatorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GolflangParserMINUS, 0)
}

func (s *OperatorContext) STAR() antlr.TerminalNode {
	return s.GetToken(GolflangParserSTAR, 0)
}

func (s *OperatorContext) SLASH() antlr.TerminalNode {
	return s.GetToken(GolflangParserSLASH, 0)
}

func (s *OperatorContext) DOUBLESTAR() antlr.TerminalNode {
	return s.GetToken(GolflangParserDOUBLESTAR, 0)
}

func (s *OperatorContext) UNDERSCORE() antlr.TerminalNode {
	return s.GetToken(GolflangParserUNDERSCORE, 0)
}

func (s *OperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(GolflangParserGT, 0)
}

func (s *OperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(GolflangParserLT, 0)
}

func (s *OperatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(GolflangParserEQ, 0)
}

func (s *OperatorContext) GE() antlr.TerminalNode {
	return s.GetToken(GolflangParserGE, 0)
}

func (s *OperatorContext) LE() antlr.TerminalNode {
	return s.GetToken(GolflangParserLE, 0)
}

func (s *OperatorContext) NE() antlr.TerminalNode {
	return s.GetToken(GolflangParserNE, 0)
}

func (s *OperatorContext) DOUBLELT() antlr.TerminalNode {
	return s.GetToken(GolflangParserDOUBLELT, 0)
}

func (s *OperatorContext) DOUBLEGTL() antlr.TerminalNode {
	return s.GetToken(GolflangParserDOUBLEGTL, 0)
}

func (s *OperatorContext) DOUBLEGTA() antlr.TerminalNode {
	return s.GetToken(GolflangParserDOUBLEGTA, 0)
}

func (s *OperatorContext) AMPERSAND() antlr.TerminalNode {
	return s.GetToken(GolflangParserAMPERSAND, 0)
}

func (s *OperatorContext) PIPE() antlr.TerminalNode {
	return s.GetToken(GolflangParserPIPE, 0)
}

func (s *OperatorContext) CARET() antlr.TerminalNode {
	return s.GetToken(GolflangParserCARET, 0)
}

func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.ExitOperator(s)
	}
}

func (s *OperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GolflangVisitor:
		return t.VisitOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GolflangParser) Operator() (localctx IOperatorContext) {
	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GolflangParserRULE_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2199014866944) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
