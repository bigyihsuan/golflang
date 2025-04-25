// Code generated from Golflang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package par // Golflang
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GolflangParser.
type GolflangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GolflangParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by GolflangParser#expr.
	VisitExpr(ctx *ExprContext) interface{}
}
