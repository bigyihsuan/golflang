// Code generated from Golflang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package par // Golflang
import "github.com/antlr4-go/antlr/v4"

// BaseGolflangListener is a complete listener for a parse tree produced by GolflangParser.
type BaseGolflangListener struct{}

var _ GolflangListener = &BaseGolflangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGolflangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGolflangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGolflangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGolflangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseGolflangListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseGolflangListener) ExitProg(ctx *ProgContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseGolflangListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseGolflangListener) ExitStmt(ctx *StmtContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseGolflangListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseGolflangListener) ExitAlias(ctx *AliasContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseGolflangListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseGolflangListener) ExitExpr(ctx *ExprContext) {}

// EnterCall is called when production call is entered.
func (s *BaseGolflangListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BaseGolflangListener) ExitCall(ctx *CallContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BaseGolflangListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BaseGolflangListener) ExitExprList(ctx *ExprListContext) {}

// EnterLambda is called when production lambda is entered.
func (s *BaseGolflangListener) EnterLambda(ctx *LambdaContext) {}

// ExitLambda is called when production lambda is exited.
func (s *BaseGolflangListener) ExitLambda(ctx *LambdaContext) {}

// EnterIdentList is called when production identList is entered.
func (s *BaseGolflangListener) EnterIdentList(ctx *IdentListContext) {}

// ExitIdentList is called when production identList is exited.
func (s *BaseGolflangListener) ExitIdentList(ctx *IdentListContext) {}

// EnterIdent is called when production ident is entered.
func (s *BaseGolflangListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *BaseGolflangListener) ExitIdent(ctx *IdentContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseGolflangListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseGolflangListener) ExitLiteral(ctx *LiteralContext) {}

// EnterLiteralList is called when production literalList is entered.
func (s *BaseGolflangListener) EnterLiteralList(ctx *LiteralListContext) {}

// ExitLiteralList is called when production literalList is exited.
func (s *BaseGolflangListener) ExitLiteralList(ctx *LiteralListContext) {}

// EnterLiteralMap is called when production literalMap is entered.
func (s *BaseGolflangListener) EnterLiteralMap(ctx *LiteralMapContext) {}

// ExitLiteralMap is called when production literalMap is exited.
func (s *BaseGolflangListener) ExitLiteralMap(ctx *LiteralMapContext) {}

// EnterLiteralMapEntry is called when production literalMapEntry is entered.
func (s *BaseGolflangListener) EnterLiteralMapEntry(ctx *LiteralMapEntryContext) {}

// ExitLiteralMapEntry is called when production literalMapEntry is exited.
func (s *BaseGolflangListener) ExitLiteralMapEntry(ctx *LiteralMapEntryContext) {}

// EnterLiteralPrimitive is called when production literalPrimitive is entered.
func (s *BaseGolflangListener) EnterLiteralPrimitive(ctx *LiteralPrimitiveContext) {}

// ExitLiteralPrimitive is called when production literalPrimitive is exited.
func (s *BaseGolflangListener) ExitLiteralPrimitive(ctx *LiteralPrimitiveContext) {}
