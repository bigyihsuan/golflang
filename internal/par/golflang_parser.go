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
		"", "", "", "", "", "", "'true'", "'false'", "", "'\"'", "'('", "')'",
		"'['", "']'", "'{'", "'}'", "'.'", "','", "':'", "';'", "':='",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "NEWLINE", "INT", "DEC", "STR", "TRUE", "FALSE", "IDENT",
		"QUOTE", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "LBRACE", "RBRACE",
		"DOT", "COMMA", "COLON", "SEMICOLON", "ASSIGN",
	}
	staticData.RuleNames = []string{
		"prog", "stmt", "alias", "expr", "literal", "literalList", "literalMap",
		"literalMapEntry", "literalPrimitive",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 20, 82, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 5, 0, 20, 8, 0,
		10, 0, 12, 0, 23, 9, 0, 1, 1, 1, 1, 3, 1, 27, 8, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 37, 8, 3, 1, 4, 1, 4, 1, 4, 3, 4, 42,
		8, 4, 1, 5, 1, 5, 3, 5, 46, 8, 5, 1, 5, 1, 5, 5, 5, 50, 8, 5, 10, 5, 12,
		5, 53, 9, 5, 1, 5, 3, 5, 56, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 3, 6, 62, 8,
		6, 1, 6, 1, 6, 5, 6, 66, 8, 6, 10, 6, 12, 6, 69, 9, 6, 1, 6, 3, 6, 72,
		8, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 0, 0, 9, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 0, 1, 1, 0, 3, 7, 83, 0, 21, 1, 0, 0, 0, 2,
		26, 1, 0, 0, 0, 4, 30, 1, 0, 0, 0, 6, 36, 1, 0, 0, 0, 8, 41, 1, 0, 0, 0,
		10, 43, 1, 0, 0, 0, 12, 59, 1, 0, 0, 0, 14, 75, 1, 0, 0, 0, 16, 79, 1,
		0, 0, 0, 18, 20, 3, 2, 1, 0, 19, 18, 1, 0, 0, 0, 20, 23, 1, 0, 0, 0, 21,
		19, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 1, 1, 0, 0, 0, 23, 21, 1, 0, 0,
		0, 24, 27, 3, 4, 2, 0, 25, 27, 3, 6, 3, 0, 26, 24, 1, 0, 0, 0, 26, 25,
		1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 29, 5, 19, 0, 0, 29, 3, 1, 0, 0, 0,
		30, 31, 5, 8, 0, 0, 31, 32, 5, 20, 0, 0, 32, 33, 3, 6, 3, 0, 33, 5, 1,
		0, 0, 0, 34, 37, 5, 8, 0, 0, 35, 37, 3, 8, 4, 0, 36, 34, 1, 0, 0, 0, 36,
		35, 1, 0, 0, 0, 37, 7, 1, 0, 0, 0, 38, 42, 3, 10, 5, 0, 39, 42, 3, 12,
		6, 0, 40, 42, 3, 16, 8, 0, 41, 38, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 41,
		40, 1, 0, 0, 0, 42, 9, 1, 0, 0, 0, 43, 45, 5, 12, 0, 0, 44, 46, 3, 6, 3,
		0, 45, 44, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 51, 1, 0, 0, 0, 47, 48,
		5, 17, 0, 0, 48, 50, 3, 6, 3, 0, 49, 47, 1, 0, 0, 0, 50, 53, 1, 0, 0, 0,
		51, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 51, 1,
		0, 0, 0, 54, 56, 5, 17, 0, 0, 55, 54, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56,
		57, 1, 0, 0, 0, 57, 58, 5, 13, 0, 0, 58, 11, 1, 0, 0, 0, 59, 61, 5, 14,
		0, 0, 60, 62, 3, 14, 7, 0, 61, 60, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62,
		67, 1, 0, 0, 0, 63, 64, 5, 17, 0, 0, 64, 66, 3, 14, 7, 0, 65, 63, 1, 0,
		0, 0, 66, 69, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 71,
		1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 70, 72, 5, 17, 0, 0, 71, 70, 1, 0, 0, 0,
		71, 72, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 74, 5, 15, 0, 0, 74, 13, 1,
		0, 0, 0, 75, 76, 3, 6, 3, 0, 76, 77, 5, 18, 0, 0, 77, 78, 3, 6, 3, 0, 78,
		15, 1, 0, 0, 0, 79, 80, 7, 0, 0, 0, 80, 17, 1, 0, 0, 0, 10, 21, 26, 36,
		41, 45, 51, 55, 61, 67, 71,
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
	GolflangParserINT        = 3
	GolflangParserDEC        = 4
	GolflangParserSTR        = 5
	GolflangParserTRUE       = 6
	GolflangParserFALSE      = 7
	GolflangParserIDENT      = 8
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
)

// GolflangParser rules.
const (
	GolflangParserRULE_prog             = 0
	GolflangParserRULE_stmt             = 1
	GolflangParserRULE_alias            = 2
	GolflangParserRULE_expr             = 3
	GolflangParserRULE_literal          = 4
	GolflangParserRULE_literalList      = 5
	GolflangParserRULE_literalMap       = 6
	GolflangParserRULE_literalMapEntry  = 7
	GolflangParserRULE_literalPrimitive = 8
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
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&20984) != 0 {
		{
			p.SetState(18)
			p.Stmt()
		}

		p.SetState(23)
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
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(24)
			p.Alias()
		}

	case 2:
		{
			p.SetState(25)
			p.Expr()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(28)
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

	// Getter signatures
	IDENT() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expr() IExprContext

	// IsAliasContext differentiates from other interfaces.
	IsAliasContext()
}

type AliasContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *AliasContext) IDENT() antlr.TerminalNode {
	return s.GetToken(GolflangParserIDENT, 0)
}

func (s *AliasContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GolflangParserASSIGN, 0)
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
		p.SetState(30)
		p.Match(GolflangParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(31)
		p.Match(GolflangParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(32)
		p.Expr()
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
	IDENT() antlr.TerminalNode
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

func (s *ExprContext) IDENT() antlr.TerminalNode {
	return s.GetToken(GolflangParserIDENT, 0)
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
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GolflangParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Match(GolflangParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GolflangParserINT, GolflangParserDEC, GolflangParserSTR, GolflangParserTRUE, GolflangParserFALSE, GolflangParserLBRACKET, GolflangParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
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
	p.EnterRule(localctx, 8, GolflangParserRULE_literal)
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GolflangParserLBRACKET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.LiteralList()
		}

	case GolflangParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.LiteralMap()
		}

	case GolflangParserINT, GolflangParserDEC, GolflangParserSTR, GolflangParserTRUE, GolflangParserFALSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(40)
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
	p.EnterRule(localctx, 10, GolflangParserRULE_literalList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(GolflangParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&20984) != 0 {
		{
			p.SetState(44)
			p.Expr()
		}

	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(47)
				p.Match(GolflangParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(48)
				p.Expr()
			}

		}
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GolflangParserCOMMA {
		{
			p.SetState(54)
			p.Match(GolflangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(57)
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
	p.EnterRule(localctx, 12, GolflangParserRULE_literalMap)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(GolflangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&20984) != 0 {
		{
			p.SetState(60)
			p.LiteralMapEntry()
		}

	}
	p.SetState(67)
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
				p.SetState(63)
				p.Match(GolflangParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(64)
				p.LiteralMapEntry()
			}

		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GolflangParserCOMMA {
		{
			p.SetState(70)
			p.Match(GolflangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(73)
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
	p.EnterRule(localctx, 14, GolflangParserRULE_literalMapEntry)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)

		var _x = p.Expr()

		localctx.(*LiteralMapEntryContext).key = _x
	}
	{
		p.SetState(76)
		p.Match(GolflangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(77)

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
	p.EnterRule(localctx, 16, GolflangParserRULE_literalPrimitive)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&248) != 0) {
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
