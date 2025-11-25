package interpreter

import (
	"bigyihsuan/golflang/internal/ast"
	"bigyihsuan/golflang/internal/obj"
	"bigyihsuan/golflang/internal/par"
	"bigyihsuan/golflang/internal/scope"
	"bigyihsuan/golflang/internal/stack"
	"bigyihsuan/golflang/internal/util"
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
	verbose, dumpStack bool
}

type AliasName string

func newInterpreter(b builder) (*Interpreter, error) {
	fs, err := antlr.NewFileStream(b.filename)
	if err != nil {
		return nil, fmt.Errorf("making file stream: %w", err)
	}
	lexer := par.NewGolflangLexer(fs)
	if lexer == nil {
		return nil, errors.New("nil lexer")
	}

	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	parser := par.NewGolflangParser(tokenStream)
	// parser.AddErrorListener(antlr.NewDiagnosticErrorListener(false))

	astBuilder := ast.NewBuilder(parser)

	interpreter := &Interpreter{
		stack:      stack.New[obj.Obj](),
		lexer:      lexer,
		parser:     parser,
		astBuilder: astBuilder,
		builtins:   initBuiltins(),
		verbose:    b.debug,
		dumpStack:  b.dumpStack,
	}
	interpreter.baseScope = scope.NewBase()
	interpreter.currentScope = &interpreter.baseScope
	return interpreter, nil
}

func (i *Interpreter) Run() error {
	if i.dumpStack {
		defer i.Exit()
	}
	if i.verbose {
		defer func() {
			fmt.Printf("stack: %s\n", i.StackString())
			fmt.Printf("aliases: %s\n", i.baseScope.String())
		}()
	}

	parseTree := i.parser.Prog()

	if i.verbose {
		// fmt.Println(parseTree.ToStringTree(i.parser.RuleNames, i.parser))
		fmt.Println(util.NewTreePrettifier().ToPrettyTree(parseTree, i.parser.RuleNames, i.parser))
	}

	ast := i.astBuilder.Visit(parseTree).(ast.Prog)
	if i.verbose {
		fmt.Printf("progAst: %v\n", ast)
	}

	err := i.Visit(ast)
	if err != nil {
		return err
	}

	return nil
}

func (i *Interpreter) Exit() {
	for i.stack.Len() != 0 {
		ele, _ := i.stack.Pop()
		fmt.Println(ele)
	}
}

func (i Interpreter) StackString() string {
	ss := []string{}
	for _, e := range i.stack.Elements() {
		ss = append(ss, e.Repr())
	}
	return fmt.Sprintf("<%s>", strings.Join(ss, ", "))
}

func (i *Interpreter) getArgs(n int, caller string) (args []obj.Obj, err error) {
	args, ok := i.stack.PopN(n)
	if !ok {
		return args, fmt.Errorf("%s: %w", caller, stack.ErrNotEnoughStackValues{Want: n, Need: n - len(args)})
	}
	return args, nil
}

func (i *Interpreter) getIdent(ident obj.Ident) (obj.Obj, error) {
	// check aliases first
	if value, err := i.currentScope.GetAlias(ident); err == nil {
		return value, nil
	} else if builtinFunc, ok := i.builtins.Get(ident); ok {
		return builtinFunc, nil
	} else {
		return nil, scope.ErrUnknownAlias{Name: ident.String()}
	}
}

func (i *Interpreter) pushValue(value obj.Obj) {
	if value.Kind() == obj.ObjKindNone {
		return
	}
	i.stack.Push(value)
	if i.dumpStack {
		fmt.Println(i.StackString())
	}
}
