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
		"prog", "stmt", "alias", "expr", "call", "exprList", "lambda", "identList",
		"ident", "literal", "literalList", "literalMap", "literalMapEntry",
		"literalPrimitive",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 23, 120, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 5, 0, 30, 8, 0, 10,
		0, 12, 0, 33, 9, 0, 1, 1, 1, 1, 3, 1, 37, 8, 1, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 3, 3, 48, 8, 3, 1, 4, 1, 4, 3, 4, 52, 8,
		4, 1, 5, 1, 5, 1, 5, 5, 5, 57, 8, 5, 10, 5, 12, 5, 60, 9, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 5, 7, 70, 8, 7, 10, 7, 12, 7, 73, 9,
		7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 3, 9, 80, 8, 9, 1, 10, 1, 10, 3, 10, 84,
		8, 10, 1, 10, 1, 10, 5, 10, 88, 8, 10, 10, 10, 12, 10, 91, 9, 10, 1, 10,
		3, 10, 94, 8, 10, 1, 10, 1, 10, 1, 11, 1, 11, 3, 11, 100, 8, 11, 1, 11,
		1, 11, 5, 11, 104, 8, 11, 10, 11, 12, 11, 107, 9, 11, 1, 11, 3, 11, 110,
		8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 0,
		0, 14, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 0, 1, 1, 0, 4,
		8, 120, 0, 31, 1, 0, 0, 0, 2, 36, 1, 0, 0, 0, 4, 40, 1, 0, 0, 0, 6, 47,
		1, 0, 0, 0, 8, 49, 1, 0, 0, 0, 10, 53, 1, 0, 0, 0, 12, 61, 1, 0, 0, 0,
		14, 66, 1, 0, 0, 0, 16, 74, 1, 0, 0, 0, 18, 79, 1, 0, 0, 0, 20, 81, 1,
		0, 0, 0, 22, 97, 1, 0, 0, 0, 24, 113, 1, 0, 0, 0, 26, 117, 1, 0, 0, 0,
		28, 30, 3, 2, 1, 0, 29, 28, 1, 0, 0, 0, 30, 33, 1, 0, 0, 0, 31, 29, 1,
		0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 1, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 34,
		37, 3, 4, 2, 0, 35, 37, 3, 6, 3, 0, 36, 34, 1, 0, 0, 0, 36, 35, 1, 0, 0,
		0, 37, 38, 1, 0, 0, 0, 38, 39, 5, 20, 0, 0, 39, 3, 1, 0, 0, 0, 40, 41,
		3, 16, 8, 0, 41, 42, 5, 21, 0, 0, 42, 43, 3, 6, 3, 0, 43, 5, 1, 0, 0, 0,
		44, 48, 3, 8, 4, 0, 45, 48, 3, 12, 6, 0, 46, 48, 3, 18, 9, 0, 47, 44, 1,
		0, 0, 0, 47, 45, 1, 0, 0, 0, 47, 46, 1, 0, 0, 0, 48, 7, 1, 0, 0, 0, 49,
		51, 3, 16, 8, 0, 50, 52, 3, 10, 5, 0, 51, 50, 1, 0, 0, 0, 51, 52, 1, 0,
		0, 0, 52, 9, 1, 0, 0, 0, 53, 58, 3, 6, 3, 0, 54, 55, 5, 1, 0, 0, 55, 57,
		3, 6, 3, 0, 56, 54, 1, 0, 0, 0, 57, 60, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0,
		58, 59, 1, 0, 0, 0, 59, 11, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 61, 62, 5,
		22, 0, 0, 62, 63, 3, 14, 7, 0, 63, 64, 5, 23, 0, 0, 64, 65, 3, 6, 3, 0,
		65, 13, 1, 0, 0, 0, 66, 71, 3, 16, 8, 0, 67, 68, 5, 18, 0, 0, 68, 70, 3,
		16, 8, 0, 69, 67, 1, 0, 0, 0, 70, 73, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71,
		72, 1, 0, 0, 0, 72, 15, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 74, 75, 5, 9, 0,
		0, 75, 17, 1, 0, 0, 0, 76, 80, 3, 20, 10, 0, 77, 80, 3, 22, 11, 0, 78,
		80, 3, 26, 13, 0, 79, 76, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 78, 1, 0,
		0, 0, 80, 19, 1, 0, 0, 0, 81, 83, 5, 13, 0, 0, 82, 84, 3, 6, 3, 0, 83,
		82, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 89, 1, 0, 0, 0, 85, 86, 5, 18,
		0, 0, 86, 88, 3, 6, 3, 0, 87, 85, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87,
		1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 93, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0,
		92, 94, 5, 18, 0, 0, 93, 92, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 95, 1,
		0, 0, 0, 95, 96, 5, 14, 0, 0, 96, 21, 1, 0, 0, 0, 97, 99, 5, 15, 0, 0,
		98, 100, 3, 24, 12, 0, 99, 98, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 105,
		1, 0, 0, 0, 101, 102, 5, 18, 0, 0, 102, 104, 3, 24, 12, 0, 103, 101, 1,
		0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 105, 106, 1, 0, 0,
		0, 106, 109, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 108, 110, 5, 18, 0, 0, 109,
		108, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 112,
		5, 16, 0, 0, 112, 23, 1, 0, 0, 0, 113, 114, 3, 6, 3, 0, 114, 115, 5, 19,
		0, 0, 115, 116, 3, 6, 3, 0, 116, 25, 1, 0, 0, 0, 117, 118, 7, 0, 0, 0,
		118, 27, 1, 0, 0, 0, 13, 31, 36, 47, 51, 58, 71, 79, 83, 89, 93, 99, 105,
		109,
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
	GolflangParserIDENT      = 9
	GolflangParserQUOTE      = 10
	GolflangParserLPAREN     = 11
	GolflangParserRPAREN     = 12
	GolflangParserLBRACKET   = 13
	GolflangParserRBRACKET   = 14
	GolflangParserLBRACE     = 15
	GolflangParserRBRACE     = 16
	GolflangParserDOT        = 17
	GolflangParserCOMMA      = 18
	GolflangParserCOLON      = 19
	GolflangParserSEMICOLON  = 20
	GolflangParserASSIGN     = 21
	GolflangParserBACKSLASH  = 22
	GolflangParserARROW      = 23
)

// GolflangParser rules.
const (
	GolflangParserRULE_prog             = 0
	GolflangParserRULE_stmt             = 1
	GolflangParserRULE_alias            = 2
	GolflangParserRULE_expr             = 3
	GolflangParserRULE_call             = 4
	GolflangParserRULE_exprList         = 5
	GolflangParserRULE_lambda           = 6
	GolflangParserRULE_identList        = 7
	GolflangParserRULE_ident            = 8
	GolflangParserRULE_literal          = 9
	GolflangParserRULE_literalList      = 10
	GolflangParserRULE_literalMap       = 11
	GolflangParserRULE_literalMapEntry  = 12
	GolflangParserRULE_literalPrimitive = 13
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
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4236272) != 0 {
		{
			p.SetState(28)
			p.Stmt()
		}

		p.SetState(33)
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
	Expr() IExprContext

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

func (s *StmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
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
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(34)
			p.Alias()
		}

	case 2:
		{
			p.SetState(35)
			p.Expr()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(38)
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
	GetValue() IExprContext

	// SetName sets the name rule contexts.
	SetName(IIdentContext)

	// SetValue sets the value rule contexts.
	SetValue(IExprContext)

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	Ident() IIdentContext
	Expr() IExprContext

	// IsAliasContext differentiates from other interfaces.
	IsAliasContext()
}

type AliasContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   IIdentContext
	value  IExprContext
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

func (s *AliasContext) GetValue() IExprContext { return s.value }

func (s *AliasContext) SetName(v IIdentContext) { s.name = v }

func (s *AliasContext) SetValue(v IExprContext) { s.value = v }

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

func (s *AliasContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
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
		p.SetState(40)

		var _x = p.Ident()

		localctx.(*AliasContext).name = _x
	}
	{
		p.SetState(41)
		p.Match(GolflangParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)

		var _x = p.Expr()

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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Call() ICallContext
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

func (s *ExprContext) Call() ICallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallContext)
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
	p.EnterRule(localctx, 6, GolflangParserRULE_expr)
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GolflangParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(44)
			p.Call()
		}

	case GolflangParserBACKSLASH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.Lambda()
		}

	case GolflangParserINT, GolflangParserDEC, GolflangParserSTR, GolflangParserTRUE, GolflangParserFALSE, GolflangParserLBRACKET, GolflangParserLBRACE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(46)
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

// ICallContext is an interface to support dynamic dispatch.
type ICallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name rule contexts.
	GetName() IIdentContext

	// GetArgs returns the args rule contexts.
	GetArgs() IExprListContext

	// SetName sets the name rule contexts.
	SetName(IIdentContext)

	// SetArgs sets the args rule contexts.
	SetArgs(IExprListContext)

	// Getter signatures
	Ident() IIdentContext
	ExprList() IExprListContext

	// IsCallContext differentiates from other interfaces.
	IsCallContext()
}

type CallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   IIdentContext
	args   IExprListContext
}

func NewEmptyCallContext() *CallContext {
	var p = new(CallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_call
	return p
}

func InitEmptyCallContext(p *CallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GolflangParserRULE_call
}

func (*CallContext) IsCallContext() {}

func NewCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallContext {
	var p = new(CallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GolflangParserRULE_call

	return p
}

func (s *CallContext) GetParser() antlr.Parser { return s.parser }

func (s *CallContext) GetName() IIdentContext { return s.name }

func (s *CallContext) GetArgs() IExprListContext { return s.args }

func (s *CallContext) SetName(v IIdentContext) { s.name = v }

func (s *CallContext) SetArgs(v IExprListContext) { s.args = v }

func (s *CallContext) Ident() IIdentContext {
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

func (s *CallContext) ExprList() IExprListContext {
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

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.EnterCall(s)
	}
}

func (s *CallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GolflangListener); ok {
		listenerT.ExitCall(s)
	}
}

func (s *CallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GolflangVisitor:
		return t.VisitCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GolflangParser) Call() (localctx ICallContext) {
	localctx = NewCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GolflangParserRULE_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)

		var _x = p.Ident()

		localctx.(*CallContext).name = _x
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4236272) != 0 {
		{
			p.SetState(50)

			var _x = p.ExprList()

			localctx.(*CallContext).args = _x
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

// IExprListContext is an interface to support dynamic dispatch.
type IExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode

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

func (s *ExprListContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(GolflangParserWHITESPACE)
}

func (s *ExprListContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(GolflangParserWHITESPACE, i)
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
	p.EnterRule(localctx, 10, GolflangParserRULE_exprList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Expr()
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(54)
				p.Match(GolflangParserWHITESPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(55)
				p.Expr()
			}

		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
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

// ILambdaContext is an interface to support dynamic dispatch.
type ILambdaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetArgs returns the args rule contexts.
	GetArgs() IIdentListContext

	// GetBody returns the body rule contexts.
	GetBody() IExprContext

	// SetArgs sets the args rule contexts.
	SetArgs(IIdentListContext)

	// SetBody sets the body rule contexts.
	SetBody(IExprContext)

	// Getter signatures
	BACKSLASH() antlr.TerminalNode
	ARROW() antlr.TerminalNode
	IdentList() IIdentListContext
	Expr() IExprContext

	// IsLambdaContext differentiates from other interfaces.
	IsLambdaContext()
}

type LambdaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	args   IIdentListContext
	body   IExprContext
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

func (s *LambdaContext) GetBody() IExprContext { return s.body }

func (s *LambdaContext) SetArgs(v IIdentListContext) { s.args = v }

func (s *LambdaContext) SetBody(v IExprContext) { s.body = v }

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

func (s *LambdaContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
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
		p.SetState(61)
		p.Match(GolflangParserBACKSLASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(62)

		var _x = p.IdentList()

		localctx.(*LambdaContext).args = _x
	}
	{
		p.SetState(63)
		p.Match(GolflangParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)

		var _x = p.Expr()

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
		p.SetState(66)
		p.Ident()
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GolflangParserCOMMA {
		{
			p.SetState(67)
			p.Match(GolflangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.Ident()
		}

		p.SetState(73)
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
	p.EnterRule(localctx, 16, GolflangParserRULE_ident)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
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
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GolflangParserLBRACKET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.LiteralList()
		}

	case GolflangParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.LiteralMap()
		}

	case GolflangParserINT, GolflangParserDEC, GolflangParserSTR, GolflangParserTRUE, GolflangParserFALSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(78)
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
		p.SetState(81)
		p.Match(GolflangParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4236272) != 0 {
		{
			p.SetState(82)
			p.Expr()
		}

	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(85)
				p.Match(GolflangParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(86)
				p.Expr()
			}

		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GolflangParserCOMMA {
		{
			p.SetState(92)
			p.Match(GolflangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(95)
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
		p.SetState(97)
		p.Match(GolflangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4236272) != 0 {
		{
			p.SetState(98)
			p.LiteralMapEntry()
		}

	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(101)
				p.Match(GolflangParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(102)
				p.LiteralMapEntry()
			}

		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GolflangParserCOMMA {
		{
			p.SetState(108)
			p.Match(GolflangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(111)
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
		p.SetState(113)

		var _x = p.Expr()

		localctx.(*LiteralMapEntryContext).key = _x
	}
	{
		p.SetState(114)
		p.Match(GolflangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)

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
		p.SetState(117)
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
