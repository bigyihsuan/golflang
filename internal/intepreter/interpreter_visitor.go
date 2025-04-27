package interpreter

import (
	"bigyihsuan/golflang/internal/obj"
	"bigyihsuan/golflang/internal/par"

	"github.com/antlr4-go/antlr/v4"
)

// Visit implements par.GolflangVisitor.
// Subtle: this method shadows the method (BaseParseTreeVisitor).Visit of Interpreter.BaseParseTreeVisitor.
func (g *Interpreter) Visit(tree antlr.ParseTree) any {
	switch t := tree.(type) {
	case *par.ProgContext:
		return g.VisitProg(t)
	}
	return nil
}

// VisitProg implements par.GolflangVisitor.
func (g *Interpreter) VisitProg(ctx *par.ProgContext) any {
	for _, stmt := range ctx.AllStmt() {
		g.VisitStmt(stmt.(*par.StmtContext))
	}
	return nil
}

// VisitStmt implements par.GolflangVisitor.
func (g *Interpreter) VisitStmt(stmt *par.StmtContext) any {
	switch child := stmt.GetChild(0).(type) {
	case *par.AliasContext:
		return g.VisitAlias(child)
	case *par.ExprContext:
		return g.VisitExpr(child)
	}
	return nil
}

// VisitAlias implements par.GolflangVisitor.
func (g *Interpreter) VisitAlias(alias *par.AliasContext) any {
	ident := AliasName(alias.IDENT().GetText())
	expr := g.VisitExpr(alias.Expr().(*par.ExprContext)).(obj.Obj)
	g.aliases[ident] = expr
	return nil
}

func (g *Interpreter) VisitExpr(expr *par.ExprContext) any {
	switch child := expr.GetChild(0).(type) {
	case *par.LiteralContext:
		v := g.VisitLiteral(child).(obj.Obj)
		g.queue.Enqueue(v)
		return v
	case antlr.TerminalNode:
		return g.aliases[AliasName(child.GetSymbol().GetText())]
	}
	return nil
}

// VisitLiteral implements par.GolflangVisitor.
func (g *Interpreter) VisitLiteral(ctx *par.LiteralContext) any {
	switch child := ctx.GetChild(0).(type) {
	case *par.LiteralPrimitiveContext:
		return g.VisitLiteralPrimitive(child)
	case *par.LiteralListContext:
		return g.VisitLiteralList(child)
	case *par.LiteralMapContext:
		return g.VisitLiteralMap(child)
	}
	return nil
}

// VisitLiteralList implements par.GolflangVisitor.
func (g *Interpreter) VisitLiteralList(ctx *par.LiteralListContext) any {
	l := obj.ZeroList()
	for _, expr := range ctx.AllExpr() {
		l = append(l, g.VisitExpr(expr.(*par.ExprContext)).(obj.Obj))
	}
	return l
}

// VisitLiteralMap implements par.GolflangVisitor.
func (g *Interpreter) VisitLiteralMap(ctx *par.LiteralMapContext) any {
	m := obj.ZeroMap()
	for _, entry := range ctx.AllLiteralMapEntry() {
		entry := g.VisitLiteralMapEntry(entry.(*par.LiteralMapEntryContext)).(obj.Entry)
		m.SetFromEntry(entry)
	}
	return m
}

// VisitLiteralMapEntry implements par.GolflangVisitor.
func (g *Interpreter) VisitLiteralMapEntry(ctx *par.LiteralMapEntryContext) any {
	k := g.VisitExpr(ctx.GetKey().(*par.ExprContext)).(obj.Obj)
	v := g.VisitExpr(ctx.GetValue().(*par.ExprContext)).(obj.Obj)
	return obj.Entry{K: k, V: v}
}

// VisitLiteralPrimitive implements par.GolflangVisitor.
func (g *Interpreter) VisitLiteralPrimitive(ctx *par.LiteralPrimitiveContext) any {
	return g.primitiveLiteral(ctx)
}

// VisitChildren implements par.GolflangVisitor.
// Subtle: this method shadows the method (BaseParseTreeVisitor).VisitChildren of Interpreter.BaseParseTreeVisitor.
func (g *Interpreter) VisitChildren(node antlr.RuleNode) any {
	return g.BaseParseTreeVisitor.VisitChildren(node)
}

// VisitErrorNode implements par.GolflangVisitor.
// Subtle: this method shadows the method (BaseParseTreeVisitor).VisitErrorNode of Interpreter.BaseParseTreeVisitor.
func (g *Interpreter) VisitErrorNode(node antlr.ErrorNode) any {
	return g.BaseParseTreeVisitor.VisitErrorNode(node)
}

// VisitTerminal implements par.GolflangVisitor.
// Subtle: this method shadows the method (BaseParseTreeVisitor).VisitTerminal of Interpreter.BaseParseTreeVisitor.
func (g *Interpreter) VisitTerminal(node antlr.TerminalNode) any {
	return g.BaseParseTreeVisitor.VisitTerminal(node)
}
