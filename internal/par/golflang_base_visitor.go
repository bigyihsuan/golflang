// Code generated from Golflang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package par // Golflang
import "github.com/antlr4-go/antlr/v4"

type BaseGolflangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGolflangVisitor) VisitProg(ctx *ProgContext) interface{} {
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
