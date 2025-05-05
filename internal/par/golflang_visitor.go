// Code generated from Golflang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package par // Golflang
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GolflangParser.
type GolflangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GolflangParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by GolflangParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by GolflangParser#alias.
	VisitAlias(ctx *AliasContext) interface{}

	// Visit a parse tree produced by GolflangParser#exprStmt.
	VisitExprStmt(ctx *ExprStmtContext) interface{}

	// Visit a parse tree produced by GolflangParser#exprList.
	VisitExprList(ctx *ExprListContext) interface{}

	// Visit a parse tree produced by GolflangParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by GolflangParser#lambda.
	VisitLambda(ctx *LambdaContext) interface{}

	// Visit a parse tree produced by GolflangParser#identList.
	VisitIdentList(ctx *IdentListContext) interface{}

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

	// Visit a parse tree produced by GolflangParser#ident.
	VisitIdent(ctx *IdentContext) interface{}

	// Visit a parse tree produced by GolflangParser#operator.
	VisitOperator(ctx *OperatorContext) interface{}
}
