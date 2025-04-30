// Code generated from Golflang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package par // Golflang
import "github.com/antlr4-go/antlr/v4"

// GolflangListener is a complete listener for a parse tree produced by GolflangParser.
type GolflangListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// EnterExprList is called when entering the exprList production.
	EnterExprList(c *ExprListContext)

	// EnterLambda is called when entering the lambda production.
	EnterLambda(c *LambdaContext)

	// EnterIdentList is called when entering the identList production.
	EnterIdentList(c *IdentListContext)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterLiteralList is called when entering the literalList production.
	EnterLiteralList(c *LiteralListContext)

	// EnterLiteralMap is called when entering the literalMap production.
	EnterLiteralMap(c *LiteralMapContext)

	// EnterLiteralMapEntry is called when entering the literalMapEntry production.
	EnterLiteralMapEntry(c *LiteralMapEntryContext)

	// EnterLiteralPrimitive is called when entering the literalPrimitive production.
	EnterLiteralPrimitive(c *LiteralPrimitiveContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)

	// ExitExprList is called when exiting the exprList production.
	ExitExprList(c *ExprListContext)

	// ExitLambda is called when exiting the lambda production.
	ExitLambda(c *LambdaContext)

	// ExitIdentList is called when exiting the identList production.
	ExitIdentList(c *IdentListContext)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitLiteralList is called when exiting the literalList production.
	ExitLiteralList(c *LiteralListContext)

	// ExitLiteralMap is called when exiting the literalMap production.
	ExitLiteralMap(c *LiteralMapContext)

	// ExitLiteralMapEntry is called when exiting the literalMapEntry production.
	ExitLiteralMapEntry(c *LiteralMapEntryContext)

	// ExitLiteralPrimitive is called when exiting the literalPrimitive production.
	ExitLiteralPrimitive(c *LiteralPrimitiveContext)
}
