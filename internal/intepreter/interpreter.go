package interpreter

import (
	"bigyihsuan/golflang/internal/ast"
	"bigyihsuan/golflang/internal/obj"
	"bigyihsuan/golflang/internal/par"
	"bigyihsuan/golflang/internal/scope"
	"bigyihsuan/golflang/internal/stack"
	"errors"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type Interpreter struct {
	lexer        *par.GolflangLexer
	parser       *par.GolflangParser
	astBuilder   *ast.Builder
	stack        stack.Stack[obj.Obj]
	baseScope    scope.Scope  // the base scope for the whole program
	currentScope *scope.Scope // the current scope
	builtins
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
		stack:      stack.New[obj.Obj](),
		lexer:      lexer,
		parser:     parser,
		astBuilder: astBuilder,
		builtins:   initBuiltins(),
	}
	interpreter.baseScope = scope.NewBase()
	interpreter.currentScope = &interpreter.baseScope
	return interpreter, nil
}

func (g *Interpreter) Run() error {
	defer g.Exit()
	parseTree := g.parser.Prog()
	fmt.Println(parseTree.ToStringTree(g.parser.RuleNames, g.parser))

	ast := g.astBuilder.Visit(parseTree).(ast.Prog)

	fmt.Printf("progAst: %v\n", ast)

	err := g.Visit(ast)
	if err != nil {
		return err
	}

	fmt.Printf("stack: %s\n", g.QueueString())
	fmt.Printf("aliases: %s\n", g.baseScope.String())
	return nil
}

func (g *Interpreter) Exit() {
	for g.stack.Len() > 0 {
		ele, _ := g.stack.Pop()
		fmt.Println(ele)
	}
}

func (g Interpreter) QueueString() string {
	ss := []string{}
	for _, e := range g.stack.Elements() {
		ss = append(ss, e.Repr())
	}
	return fmt.Sprintf("<%s>", strings.Join(ss, ", "))
}

func (g *Interpreter) EvalObj(o obj.Obj) (obj.Obj, error) {
	switch o.Kind() {
	case obj.ObjKindNone:
		return o, nil
	case obj.ObjKindBool:
		return o, nil
	case obj.ObjKindDec:
		return o, nil
	case obj.ObjKindInt:
		return o, nil
	case obj.ObjKindList:
		return o, nil
	case obj.ObjKindMap:
		return o, nil
	case obj.ObjKindStr:
		return o, nil
	case obj.ObjKindIdent:
		return g.currentScope.GetAlias(o.(obj.Ident))
	case obj.ObjKindLambda:
		return g.EvalLambda(o.(Lambda))
	default:
		panic(fmt.Errorf("unimplemented obj.ObjKind %s", o.Kind()))
	}
}

func (g *Interpreter) EvalLambda(lambda Lambda) (obj.Obj, error) {
	// set up a new scope for this lambda
	lambdaScope := scope.New(g.currentScope)
	g.currentScope = &lambdaScope
	defer func() {
		// destroy the lambda scope, and move back to the outer scope
		g.currentScope = g.currentScope.Parent
	}()

	// assign values to arguments
	values, ok := g.stack.PopN(len(lambda.Args))
	if !ok {
		panic(ErrNotEnoughStackValues{
			Want: len(lambda.Args),
			Need: len(lambda.Args) - len(values),
		})
	}

	for i, arg := range lambda.Args {
		v, err := g.EvalObj(values[i])
		if err != nil {
			return nil, err
		}
		g.currentScope.SetAlias(arg, v)
	}

	// execute the lambda body
	err := g.VisitExprStmt(lambda.Body)
	if err != nil {
		return nil, err
	}
	// get the return value
	v, ok := g.stack.Pop()
	if !ok {
		return v, ErrNotEnoughStackValues{Want: 1, Need: 1}
	}
	return v, nil
}
