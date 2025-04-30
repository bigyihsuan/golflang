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
	// aliases    map[obj.Ident]obj.Obj
	baseScope    Scope  // the base scope for the whole program
	currentScope *Scope // the current scope
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

	interpreter := &Interpreter{
		queue:      queue.New[obj.Obj](),
		lexer:      lexer,
		parser:     parser,
		astBuilder: astBuilder,
	}
	interpreter.baseScope = NewBaseScope()
	interpreter.currentScope = &interpreter.baseScope
	return interpreter, nil
}

func (g *Interpreter) Run() error {
	defer g.Exit()
	parseTree := g.parser.Prog()
	fmt.Println(parseTree.ToStringTree(g.parser.RuleNames, g.parser))

	ast := g.astBuilder.Visit(parseTree).(ast.Prog)

	fmt.Printf("progAst: %v\n", ast)

	g.Visit(ast)

	fmt.Printf("queue: %s\n", g.QueueString())
	fmt.Printf("aliases: %s\n", g.baseScope.String())
	return nil
}

func (g *Interpreter) Exit() {
	for g.queue.Len() > 0 {
		ele := g.DequeueOne()
		fmt.Println(ele)
	}
}

func (g Interpreter) QueueString() string {
	ss := []string{}
	for _, e := range g.queue.Elements() {
		ss = append(ss, e.Repr())
	}
	return fmt.Sprintf("<%s>", strings.Join(ss, ", "))
}

func (g *Interpreter) EvalObj(o obj.Obj) obj.Obj {
	switch o.Kind() {
	case obj.ObjKindNone:
		return o
	case obj.ObjKindBool:
		return o
	case obj.ObjKindDec:
		return o
	case obj.ObjKindInt:
		return o
	case obj.ObjKindList:
		return o
	case obj.ObjKindMap:
		return o
	case obj.ObjKindStr:
		return o
	case obj.ObjKindIdent:
		return g.GetAlias(o.(obj.Ident))
	case obj.ObjKindLambda:
		return g.EvalLambda(o.(Lambda))
	default:
		panic(fmt.Errorf("unexpected obj.ObjKind %s", o.Kind()))
	}
}

func (g *Interpreter) EvalLambda(lambda Lambda) obj.Obj {
	// set up a new scope for this lambda
	lambdaScope := NewScope(g.currentScope)
	g.currentScope = &lambdaScope
	defer func() {
		// destroy the lambda scope, and move back to the outer scope
		g.currentScope = g.currentScope.parent
	}()

	// assign values to arguments
	values := g.Dequeue(len(lambda.Args))
	for i, arg := range lambda.Args {
		g.SetAlias(arg, g.EvalObj(values[i]))
	}

	// execute the lambda body
	g.VisitExprStmt(lambda.Body)
	// get the return value
	return g.DequeueOne()
}

func (g *Interpreter) Push(value obj.Obj) {
	g.queue.Push(value)
}

func (g *Interpreter) DequeueOne() obj.Obj {
	return g.Dequeue(1)[0]
}
func (g *Interpreter) Dequeue(n int) (o []obj.Obj) {
	for i := range n {
		left := n - i
		v, ok := g.queue.Dequeue()
		if !ok {
			panic(fmt.Errorf(InterpreterErrorNotEnoughValuesOnQueue, left))
		}
		o = append(o, v)
	}
	return
}

func (g Interpreter) GetAlias(name obj.Ident) obj.Obj {
	return g.currentScope.GetAlias(name)
}
func (g *Interpreter) SetAlias(name obj.Ident, value obj.Obj) {
	g.currentScope.SetAlias(name, value)
}
