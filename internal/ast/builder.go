package ast

import (
	"bigyihsuan/golflang/internal/obj"
	"bigyihsuan/golflang/internal/par"
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

// var _ par.GolflangVisitor = (*Builder)(nil)

type Builder struct {
	// par.BaseGolflangVisitor
	parser *par.GolflangParser
}

func NewBuilder(parser *par.GolflangParser) *Builder {
	return &Builder{parser: parser}
}

func (g Builder) unknown(kind string, tree antlr.ParseTree) error {
	return fmt.Errorf("%T: unknown %s %T: %s", g, kind, tree, tree.ToStringTree(g.parser.RuleNames, g.parser))
}

func (g *Builder) Visit(tree antlr.ParseTree) any {
	switch t := tree.(type) {
	case *par.ProgContext:
		return g.VisitProg(t).(Prog)
	default:
		panic(g.unknown("ParseTree", t))
	}
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
	default:
		panic(g.unknown("StmtContext", child.(antlr.ParseTree)))
	}
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
	switch expr := expr.GetChild(0).(type) {
	case *par.LiteralContext:
		return g.VisitLiteral(expr)
	case *par.IdentContext:
		return g.VisitIdent(expr)
	case *par.LambdaContext:
		return g.VisitLambda(expr)
	case *par.CallContext:
		return g.VisitCall(expr)
	default:
		panic(g.unknown("ExprContext", expr.(antlr.ParseTree)))
	}
}

func (g *Builder) VisitCall(call *par.CallContext) any {
	name := g.VisitIdent(call.GetName().(*par.IdentContext)).(Ident)
	args := []Expr{}
	if call.GetArgs() != nil {
		args = g.VisitExprList(call.GetArgs().(*par.ExprListContext)).([]Expr)
	}

	return Call{
		Name: name,
		Args: args,
	}
}

func (g *Builder) VisitExprList(exprs *par.ExprListContext) any {
	es := []Expr{}
	for _, expr := range exprs.AllExpr() {
		es = append(es, g.VisitExpr(expr.(*par.ExprContext)).(Expr))
	}
	return es
}

func (g *Builder) VisitLambda(lambda *par.LambdaContext) any {
	args := g.VisitIdentList(lambda.GetArgs().(*par.IdentListContext)).(IdentList)
	body := g.VisitExpr(lambda.GetBody().(*par.ExprContext)).(Expr)
	return Lambda{
		Args: args,
		Body: body,
	}
}

func (g *Builder) VisitIdentList(idents *par.IdentListContext) any {
	is := IdentList{}
	for _, ident := range idents.AllIdent() {
		is = append(is, g.VisitIdent(ident.(*par.IdentContext)).(Ident))
	}
	return is
}

func (g *Builder) VisitIdent(ident *par.IdentContext) any {
	return Ident(obj.Ident(ident.GetText()))
}

func (g *Builder) VisitLiteral(literal *par.LiteralContext) any {
	switch child := literal.GetChild(0).(type) {
	case *par.LiteralPrimitiveContext:
		return g.VisitLiteralPrimitive(child)
	case *par.LiteralListContext:
		return g.VisitLiteralList(child)
	case *par.LiteralMapContext:
		return g.VisitLiteralMap(child)
	default:
		panic(g.unknown("LiteralContext", child.(antlr.ParseTree)))
	}
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
		panic(fmt.Errorf("%T: unknown primitiveLiteral: %s", g, c.ToStringTree(g.parser.RuleNames, g.parser)))
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
