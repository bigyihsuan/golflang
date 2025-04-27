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
		"", "", "", "", "", "", "'true'", "'false'", "'\"'", "'('", "')'", "'['",
		"']'", "'{'", "'}'", "'.'", "','", "':'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "NEWLINE", "INT", "DEC", "STR", "TRUE", "FALSE", "QUOTE",
		"LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "LBRACE", "RBRACE", "DOT",
		"COMMA", "COLON", "SEMICOLON",
	}
	staticData.RuleNames = []string{
		"prog", "literal", "literalList", "literalMap", "literalMapEntry", "literalPrimitive",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 18, 62, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 5, 0, 14, 8, 0, 10, 0, 12, 0, 17, 9, 0, 1, 1, 1, 1,
		1, 1, 3, 1, 22, 8, 1, 1, 2, 1, 2, 3, 2, 26, 8, 2, 1, 2, 1, 2, 5, 2, 30,
		8, 2, 10, 2, 12, 2, 33, 9, 2, 1, 2, 3, 2, 36, 8, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 3, 3, 42, 8, 3, 1, 3, 1, 3, 5, 3, 46, 8, 3, 10, 3, 12, 3, 49, 9, 3,
		1, 3, 3, 3, 52, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 0, 0, 6, 0, 2, 4, 6, 8, 10, 0, 1, 1, 0, 3, 7, 64, 0, 15, 1, 0, 0, 0,
		2, 21, 1, 0, 0, 0, 4, 23, 1, 0, 0, 0, 6, 39, 1, 0, 0, 0, 8, 55, 1, 0, 0,
		0, 10, 59, 1, 0, 0, 0, 12, 14, 3, 2, 1, 0, 13, 12, 1, 0, 0, 0, 14, 17,
		1, 0, 0, 0, 15, 13, 1, 0, 0, 0, 15, 16, 1, 0, 0, 0, 16, 1, 1, 0, 0, 0,
		17, 15, 1, 0, 0, 0, 18, 22, 3, 4, 2, 0, 19, 22, 3, 6, 3, 0, 20, 22, 3,
		10, 5, 0, 21, 18, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 21, 20, 1, 0, 0, 0, 22,
		3, 1, 0, 0, 0, 23, 25, 5, 11, 0, 0, 24, 26, 3, 2, 1, 0, 25, 24, 1, 0, 0,
		0, 25, 26, 1, 0, 0, 0, 26, 31, 1, 0, 0, 0, 27, 28, 5, 16, 0, 0, 28, 30,
		3, 2, 1, 0, 29, 27, 1, 0, 0, 0, 30, 33, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0,
		31, 32, 1, 0, 0, 0, 32, 35, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 34, 36, 5,
		16, 0, 0, 35, 34, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37,
		38, 5, 12, 0, 0, 38, 5, 1, 0, 0, 0, 39, 41, 5, 13, 0, 0, 40, 42, 3, 8,
		4, 0, 41, 40, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 47, 1, 0, 0, 0, 43, 44,
		5, 16, 0, 0, 44, 46, 3, 8, 4, 0, 45, 43, 1, 0, 0, 0, 46, 49, 1, 0, 0, 0,
		47, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 51, 1, 0, 0, 0, 49, 47, 1,
		0, 0, 0, 50, 52, 5, 16, 0, 0, 51, 50, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52,
		53, 1, 0, 0, 0, 53, 54, 5, 14, 0, 0, 54, 7, 1, 0, 0, 0, 55, 56, 3, 2, 1,
		0, 56, 57, 5, 17, 0, 0, 57, 58, 3, 2, 1, 0, 58, 9, 1, 0, 0, 0, 59, 60,
		7, 0, 0, 0, 60, 11, 1, 0, 0, 0, 8, 15, 21, 25, 31, 35, 41, 47, 51,
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
	GolflangParserQUOTE      = 8
	GolflangParserLPAREN     = 9
	GolflangParserRPAREN     = 10
	GolflangParserLBRACKET   = 11
	GolflangParserRBRACKET   = 12
	GolflangParserLBRACE     = 13
	GolflangParserRBRACE     = 14
	GolflangParserDOT        = 15
	GolflangParserCOMMA      = 16
	GolflangParserCOLON      = 17
	GolflangParserSEMICOLON  = 18
)

// GolflangParser rules.
const (
	GolflangParserRULE_prog             = 0
	GolflangParserRULE_literal          = 1
	GolflangParserRULE_literalList      = 2
	GolflangParserRULE_literalMap       = 3
	GolflangParserRULE_literalMapEntry  = 4
	GolflangParserRULE_literalPrimitive = 5
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLiteral() []ILiteralContext
	Literal(i int) ILiteralContext

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

func (s *ProgContext) AllLiteral() []ILiteralContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILiteralContext); ok {
			len++
		}
	}

	tst := make([]ILiteralContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILiteralContext); ok {
			tst[i] = t.(ILiteralContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Literal(i int) ILiteralContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
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

	return t.(ILiteralContext)
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&10488) != 0 {
		{
			p.SetState(12)
			p.Literal()
		}

		p.SetState(17)
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
	p.EnterRule(localctx, 2, GolflangParserRULE_literal)
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GolflangParserLBRACKET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(18)
			p.LiteralList()
		}

	case GolflangParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(19)
			p.LiteralMap()
		}

	case GolflangParserINT, GolflangParserDEC, GolflangParserSTR, GolflangParserTRUE, GolflangParserFALSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(20)
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
	AllLiteral() []ILiteralContext
	Literal(i int) ILiteralContext
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

func (s *LiteralListContext) AllLiteral() []ILiteralContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILiteralContext); ok {
			len++
		}
	}

	tst := make([]ILiteralContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILiteralContext); ok {
			tst[i] = t.(ILiteralContext)
			i++
		}
	}

	return tst
}

func (s *LiteralListContext) Literal(i int) ILiteralContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
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

	return t.(ILiteralContext)
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
	p.EnterRule(localctx, 4, GolflangParserRULE_literalList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(23)
		p.Match(GolflangParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&10488) != 0 {
		{
			p.SetState(24)
			p.Literal()
		}

	}
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(27)
				p.Match(GolflangParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(28)
				p.Literal()
			}

		}
		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GolflangParserCOMMA {
		{
			p.SetState(34)
			p.Match(GolflangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(37)
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
	p.EnterRule(localctx, 6, GolflangParserRULE_literalMap)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.Match(GolflangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&10488) != 0 {
		{
			p.SetState(40)
			p.LiteralMapEntry()
		}

	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(43)
				p.Match(GolflangParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(44)
				p.LiteralMapEntry()
			}

		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GolflangParserCOMMA {
		{
			p.SetState(50)
			p.Match(GolflangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(53)
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
	GetKey() ILiteralContext

	// GetValue returns the value rule contexts.
	GetValue() ILiteralContext

	// SetKey sets the key rule contexts.
	SetKey(ILiteralContext)

	// SetValue sets the value rule contexts.
	SetValue(ILiteralContext)

	// Getter signatures
	COLON() antlr.TerminalNode
	AllLiteral() []ILiteralContext
	Literal(i int) ILiteralContext

	// IsLiteralMapEntryContext differentiates from other interfaces.
	IsLiteralMapEntryContext()
}

type LiteralMapEntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	key    ILiteralContext
	value  ILiteralContext
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

func (s *LiteralMapEntryContext) GetKey() ILiteralContext { return s.key }

func (s *LiteralMapEntryContext) GetValue() ILiteralContext { return s.value }

func (s *LiteralMapEntryContext) SetKey(v ILiteralContext) { s.key = v }

func (s *LiteralMapEntryContext) SetValue(v ILiteralContext) { s.value = v }

func (s *LiteralMapEntryContext) COLON() antlr.TerminalNode {
	return s.GetToken(GolflangParserCOLON, 0)
}

func (s *LiteralMapEntryContext) AllLiteral() []ILiteralContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILiteralContext); ok {
			len++
		}
	}

	tst := make([]ILiteralContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILiteralContext); ok {
			tst[i] = t.(ILiteralContext)
			i++
		}
	}

	return tst
}

func (s *LiteralMapEntryContext) Literal(i int) ILiteralContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
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

	return t.(ILiteralContext)
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
	p.EnterRule(localctx, 8, GolflangParserRULE_literalMapEntry)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)

		var _x = p.Literal()

		localctx.(*LiteralMapEntryContext).key = _x
	}
	{
		p.SetState(56)
		p.Match(GolflangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)

		var _x = p.Literal()

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
	p.EnterRule(localctx, 10, GolflangParserRULE_literalPrimitive)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
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
