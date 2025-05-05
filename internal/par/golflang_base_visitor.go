// Code generated from Golflang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package par // Golflang
import "github.com/antlr4-go/antlr/v4"

type BaseGolflangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGolflangVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolflangVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolflangVisitor) VisitAlias(ctx *AliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolflangVisitor) VisitExprStmt(ctx *ExprStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolflangVisitor) VisitExprList(ctx *ExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolflangVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolflangVisitor) VisitLambda(ctx *LambdaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolflangVisitor) VisitIdentList(ctx *IdentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolflangVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolflangVisitor) VisitLiteralList(ctx *LiteralListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolflangVisitor) VisitLiteralMap(ctx *LiteralMapContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolflangVisitor) VisitLiteralMapEntry(ctx *LiteralMapEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolflangVisitor) VisitLiteralPrimitive(ctx *LiteralPrimitiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolflangVisitor) VisitIdent(ctx *IdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolflangVisitor) VisitOperator(ctx *OperatorContext) interface{} {
	return v.VisitChildren(ctx)
}
