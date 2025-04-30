package interpreter

import (
	"bigyihsuan/golflang/internal/ast"
	"bigyihsuan/golflang/internal/obj"
	"bigyihsuan/golflang/internal/par"
	"bigyihsuan/golflang/internal/queue"
	"errors"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

var _ par.GolflangVisitor = &Interpreter{}

type Interpreter struct {
	antlr.BaseParseTreeVisitor
	lexer      *par.GolflangLexer
	parser     *par.GolflangParser
	astBuilder *ast.Builder
	queue      queue.Queue[obj.Obj]
	aliases    map[AliasName]obj.Obj
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
		queue:      queue.New[obj.Obj](),
		lexer:      lexer,
		parser:     parser,
		astBuilder: ast.NewBuilder(parser),
		aliases:    make(map[AliasName]obj.Obj),
	}, nil
}

func (g *Interpreter) Run() error {
	prog := g.parser.Prog()
	fmt.Println(prog.ToStringTree(g.parser.RuleNames, g.parser))

	progAst := g.astBuilder.Visit(prog)

	fmt.Printf("progAst: %v\n", progAst)

	// g.Visit(prog)
	// fmt.Printf("queue: %s\n", g.queue)
	// fmt.Printf("aliases: %s\n", g.aliases)
	return nil
}
