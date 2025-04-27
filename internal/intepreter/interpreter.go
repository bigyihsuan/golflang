package interpreter

import (
	"bigyihsuan/golflang/internal/obj"
	"bigyihsuan/golflang/internal/par"
	"bigyihsuan/golflang/internal/queue"
	"errors"
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

var _ par.GolflangVisitor = &Interpreter{}

type Interpreter struct {
	antlr.BaseParseTreeVisitor
	lexer   *par.GolflangLexer
	parser  *par.GolflangParser
	queue   queue.Queue[obj.Obj]
	aliases map[AliasName]obj.Obj
}

type AliasName string

func New(filename string) (*Interpreter, error) {
	fs, err := antlr.NewFileStream(filename)
	if err != nil {
		return nil, fmt.Errorf("making file stream: %w", err)
	}
	lexer := par.NewGolflangLexer(fs)
	if lexer == nil {
		return nil, errors.New("nil lexer")
	}

	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	parser := par.NewGolflangParser(tokenStream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(false))

	return &Interpreter{
		queue:   queue.New[obj.Obj](),
		lexer:   lexer,
		parser:  parser,
		aliases: make(map[AliasName]obj.Obj),
	}, nil
}

func (g *Interpreter) Parse() error {
	prog := g.parser.Prog()
	fmt.Println(prog.ToStringTree(g.parser.RuleNames, g.parser))
	g.Visit(prog)
	fmt.Printf("queue: %s\n", g.queue)
	fmt.Printf("aliases: %s\n", g.aliases)
	return nil
}

func (g *Interpreter) primitiveLiteral(c par.ILiteralPrimitiveContext) obj.Obj {
	switch {
	case c.INT() != nil:
		return g.int(c)
	case c.DEC() != nil:
		return g.dec(c)
	case c.STR() != nil:
		return g.str(c)
	case c.TRUE() != nil:
		return g.bool(c)
	case c.FALSE() != nil:
		return g.bool(c)
	default:
		panic(fmt.Errorf("invalid LiteralPrimitiveContext: %s", c.ToStringTree(g.parser.RuleNames, g.parser)))
	}
}

func (g *Interpreter) int(c par.ILiteralPrimitiveContext) obj.Int {
	i, err := strconv.ParseInt(c.GetText(), 10, 64)
	if err != nil {
		panic(err)
	}
	return obj.Int(i)
}

func (g *Interpreter) dec(c par.ILiteralPrimitiveContext) obj.Dec {
	f, err := strconv.ParseFloat(c.GetText(), 64)
	if err != nil {
		panic(err)
	}
	return obj.Dec(f)
}

func (g *Interpreter) str(c par.ILiteralPrimitiveContext) obj.Str {
	v := c.GetText()
	return obj.Str(v)
}

func (g *Interpreter) bool(c par.ILiteralPrimitiveContext) obj.Bool {
	v, err := strconv.ParseBool(c.GetText())
	if err != nil {
		panic(err)
	}
	return obj.Bool(v)
}
