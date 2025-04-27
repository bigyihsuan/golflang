package interpreter

import (
	"bigyihsuan/golflang/internal/obj"
	"bigyihsuan/golflang/internal/par"

	"github.com/antlr4-go/antlr/v4"
)

// Visit implements par.GolflangVisitor.
// Subtle: this method shadows the method (BaseParseTreeVisitor).Visit of Interpreter.BaseParseTreeVisitor.
func (g *Interpreter) Visit(tree antlr.ParseTree) interface{} {
	switch t := tree.(type) {
	case *par.ProgContext:
		return g.VisitProg(t)
	}
	return nil
}

// VisitProg implements par.GolflangVisitor.
func (g *Interpreter) VisitProg(ctx *par.ProgContext) interface{} {
	for _, lit := range ctx.AllLiteral() {
		g.queue.Enqueue(g.VisitLiteral(lit.(*par.LiteralContext)).(obj.Obj))
	}
	return nil
}

// VisitLiteral implements par.GolflangVisitor.
func (g *Interpreter) VisitLiteral(ctx *par.LiteralContext) interface{} {
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
func (g *Interpreter) VisitLiteralList(ctx *par.LiteralListContext) interface{} {
	l := obj.ZeroList()
	for _, lit := range ctx.AllLiteral() {
		l = append(l, g.VisitLiteral(lit.(*par.LiteralContext)).(obj.Obj))
	}
	return l
}

// VisitLiteralMap implements par.GolflangVisitor.
func (g *Interpreter) VisitLiteralMap(ctx *par.LiteralMapContext) interface{} {
	m := obj.ZeroMap()
	for _, entry := range ctx.AllLiteralMapEntry() {
		entry := g.VisitLiteralMapEntry(entry.(*par.LiteralMapEntryContext)).(obj.Entry)
		m.SetFromEntry(entry)
	}
	return m
}

// VisitLiteralMapEntry implements par.GolflangVisitor.
func (g *Interpreter) VisitLiteralMapEntry(ctx *par.LiteralMapEntryContext) interface{} {
	k := g.VisitLiteral(ctx.GetKey().(*par.LiteralContext)).(obj.Obj)
	v := g.VisitLiteral(ctx.GetValue().(*par.LiteralContext)).(obj.Obj)
	return obj.Entry{K: k, V: v}
}

// VisitLiteralPrimitive implements par.GolflangVisitor.
func (g *Interpreter) VisitLiteralPrimitive(ctx *par.LiteralPrimitiveContext) interface{} {
	return g.primitiveLiteral(ctx)
}

// VisitChildren implements par.GolflangVisitor.
// Subtle: this method shadows the method (BaseParseTreeVisitor).VisitChildren of Interpreter.BaseParseTreeVisitor.
func (g *Interpreter) VisitChildren(node antlr.RuleNode) interface{} {
	return g.BaseParseTreeVisitor.VisitChildren(node)
}

// VisitErrorNode implements par.GolflangVisitor.
// Subtle: this method shadows the method (BaseParseTreeVisitor).VisitErrorNode of Interpreter.BaseParseTreeVisitor.
func (g *Interpreter) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return g.BaseParseTreeVisitor.VisitErrorNode(node)
}

// VisitTerminal implements par.GolflangVisitor.
// Subtle: this method shadows the method (BaseParseTreeVisitor).VisitTerminal of Interpreter.BaseParseTreeVisitor.
func (g *Interpreter) VisitTerminal(node antlr.TerminalNode) interface{} {
	return g.BaseParseTreeVisitor.VisitTerminal(node)
}
