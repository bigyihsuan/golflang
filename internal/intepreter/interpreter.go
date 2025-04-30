package interpreter

import (
	"bigyihsuan/golflang/internal/ast"
	"bigyihsuan/golflang/internal/obj"
	"bigyihsuan/golflang/internal/par"
	"bigyihsuan/golflang/internal/queue"
	"errors"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type Interpreter struct {
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

	astBuilder := ast.NewBuilder(parser)

	return &Interpreter{
		queue:      queue.New[obj.Obj](),
		lexer:      lexer,
		parser:     parser,
		astBuilder: astBuilder,
		aliases:    make(map[AliasName]obj.Obj),
	}, nil
}

func (g *Interpreter) Run() error {
	parseTree := g.parser.Prog()
	fmt.Println(parseTree.ToStringTree(g.parser.RuleNames, g.parser))

	ast := g.astBuilder.Visit(parseTree).(ast.Prog)

	fmt.Printf("progAst: %v\n", ast)

	g.Visit(ast)

	fmt.Printf("queue: %s\n", g.QueueString())
	fmt.Printf("aliases: %s\n", g.aliases)
	return nil
}

func (g Interpreter) QueueString() string {
	ss := []string{}
	for _, e := range g.queue.Elements() {
		ss = append(ss, e.Repr())
	}
	return fmt.Sprintf("<%s>", strings.Join(ss, ", "))
}
