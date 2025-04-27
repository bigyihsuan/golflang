// Code generated from Golflang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package par // Golflang
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GolflangParser.
type GolflangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GolflangParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by GolflangParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by GolflangParser#literalList.
	VisitLiteralList(ctx *LiteralListContext) interface{}

	// Visit a parse tree produced by GolflangParser#literalMap.
	VisitLiteralMap(ctx *LiteralMapContext) interface{}

	// Visit a parse tree produced by GolflangParser#literalMapEntry.
	VisitLiteralMapEntry(ctx *LiteralMapEntryContext) interface{}

	// Visit a parse tree produced by GolflangParser#literalPrimitive.
	VisitLiteralPrimitive(ctx *LiteralPrimitiveContext) interface{}
}
