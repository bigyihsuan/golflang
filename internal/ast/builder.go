package ast

import (
	"bigyihsuan/golflang/internal/obj"
	"bigyihsuan/golflang/internal/par"
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type Builder struct {
	par.BaseGolflangVisitor
	parser *par.GolflangParser
}

func NewBuilder(parser *par.GolflangParser) *Builder {
	return &Builder{parser: parser}
}

func (g *Builder) Visit(tree antlr.ParseTree) any {
	switch t := tree.(type) {
	case *par.ProgContext:
		return g.VisitProg(t).(Prog)
	}
	return nil
}

func (g *Builder) VisitProg(ctx *par.ProgContext) any {
	stmts := []Stmt{}
	for _, stmt := range ctx.AllStmt() {
		stmts = append(stmts, g.VisitStmt(stmt.(*par.StmtContext)).(Stmt))
	}
	return Prog{stmts}
}

func (g *Builder) VisitStmt(stmt *par.StmtContext) any {
	switch child := stmt.GetChild(0).(type) {
	case *par.AliasContext:
		return g.VisitAlias(child)
	case *par.ExprContext:
		return g.VisitExpr(child)
	}
	return nil
}

func (g *Builder) VisitAlias(alias *par.AliasContext) any {
	ident := g.VisitIdent(alias.GetName().(*par.IdentContext)).(Ident)
	expr := g.VisitExpr(alias.Expr().(*par.ExprContext)).(Expr)
	return Alias{
		Name:  ident,
		Value: expr,
	}
}

func (g *Builder) VisitExpr(expr *par.ExprContext) any {
	switch child := expr.GetChild(0).(type) {
	case *par.LiteralContext:
		return g.VisitLiteral(child)
	case *par.IdentContext:
		return g.VisitIdent(child)
	}
	return nil
}

func (g *Builder) VisitIdent(ident *par.IdentContext) any {
	return Ident(obj.Ident(ident.GetText()))
}

func (g *Builder) VisitLiteral(ctx *par.LiteralContext) any {
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

func (g *Builder) VisitLiteralList(ctx *par.LiteralListContext) any {
	l := []Expr{}
	for _, expr := range ctx.AllExpr() {
		l = append(l, g.VisitExpr(expr.(*par.ExprContext)).(Expr))
	}
	return LiteralList{Value: l}
}

func (g *Builder) VisitLiteralMap(ctx *par.LiteralMapContext) any {
	m := []LiteralMapEntry{}
	for _, entry := range ctx.AllLiteralMapEntry() {
		m = append(m, g.VisitLiteralMapEntry(entry.(*par.LiteralMapEntryContext)).(LiteralMapEntry))
	}
	return LiteralMap{Value: m}
}

// VisitLiteralMapEntry implements par.GolflangVisitor.
func (g *Builder) VisitLiteralMapEntry(ctx *par.LiteralMapEntryContext) any {
	k := g.VisitExpr(ctx.GetKey().(*par.ExprContext)).(Expr)
	v := g.VisitExpr(ctx.GetValue().(*par.ExprContext)).(Expr)
	return LiteralMapEntry{K: k, V: v}
}

// VisitLiteralPrimitive implements par.GolflangVisitor.
func (g *Builder) VisitLiteralPrimitive(ctx *par.LiteralPrimitiveContext) any {
	return LiteralPrimitive{Value: g.primitiveLiteral(ctx)}
}

func (g *Builder) primitiveLiteral(c par.ILiteralPrimitiveContext) obj.Obj {
	switch {
	case c.INT() != nil:
		return g.int(c)
	case c.DEC() != nil:
		return g.dec(c)
	case c.STR() != nil:
		return g.str(c)
	case c.TRUE() != nil:
		return g.bool(c)
	case c.FALSE() != nil:
		return g.bool(c)
	default:
		panic(fmt.Errorf("invalid LiteralPrimitiveContext: %s", c.ToStringTree(g.parser.RuleNames, g.parser)))
	}
}

func (g *Builder) int(c par.ILiteralPrimitiveContext) obj.Int {
	i, err := strconv.ParseInt(c.GetText(), 10, 64)
	if err != nil {
		panic(err)
	}
	return obj.Int(i)
}

func (g *Builder) dec(c par.ILiteralPrimitiveContext) obj.Dec {
	f, err := strconv.ParseFloat(c.GetText(), 64)
	if err != nil {
		panic(err)
	}
	return obj.Dec(f)
}

func (g *Builder) str(c par.ILiteralPrimitiveContext) obj.Str {
	v := c.GetText()
	return obj.NewStr(v)
}

func (g *Builder) bool(c par.ILiteralPrimitiveContext) obj.Bool {
	v, err := strconv.ParseBool(c.GetText())
	if err != nil {
		panic(err)
	}
	return obj.Bool(v)
}
