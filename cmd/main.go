package main

import (
	"bigyihsuan/golflang/internal/par"
	"flag"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

var (
	filename = flag.String("f", "", "input file name")
)

func main() {
	flag.Parse()
	fmt.Println(*filename)

	fs, err := antlr.NewFileStream(*filename)
	if err != nil {
		panic(fmt.Errorf("making file stream: %w", err))
	}
	lexer := par.NewGolflangLexer(fs)
	if lexer == nil {
		panic("nil lexer")
	}

	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	parser := par.NewGolflangParser(tokenStream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	tree := parser.Prog()
	fmt.Println(tree.ToStringTree(parser.RuleNames, parser))

}
