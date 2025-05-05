package ast

import (
	"bigyihsuan/golflang/internal/obj"
	"bigyihsuan/golflang/internal/par"
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

var _ par.GolflangVisitor = (*Builder)(nil)

type Builder struct {
	par.BaseGolflangVisitor
	parser *par.GolflangParser
}

func NewBuilder(parser *par.GolflangParser) *Builder {
	return &Builder{parser: parser}
}

func (b Builder) unknown(kind string, tree antlr.ParseTree) error {
	return fmt.Errorf("%T: unknown %s %T: %s", b, kind, tree, tree.ToStringTree(b.parser.RuleNames, b.parser))
}

func (b *Builder) Visit(tree antlr.ParseTree) any {
	switch t := tree.(type) {
	case *par.ProgContext:
		return b.VisitProg(t).(Prog)
	default:
		panic(b.unknown("ParseTree", t))
	}
}

func (b *Builder) VisitProg(ctx *par.ProgContext) any {
	stmts := []Stmt{}
	for _, stmt := range ctx.AllStmt() {
		stmts = append(stmts, b.VisitStmt(stmt.(*par.StmtContext)).(Stmt))
	}
	return Prog{stmts}
}

func (b *Builder) VisitStmt(stmt *par.StmtContext) any {
	switch child := stmt.GetChild(0).(type) {
	case *par.AliasContext:
		return b.VisitAlias(child)
	case *par.ExprStmtContext:
		return b.VisitExprStmt(child)
	default:
		panic(b.unknown("StmtContext", child.(antlr.ParseTree)))
	}
}

func (b *Builder) VisitAlias(alias *par.AliasContext) any {
	ident := b.VisitIdent(alias.GetName().(*par.IdentContext)).(Ident)
	value := b.VisitExprList(alias.GetValue().(*par.ExprListContext)).(ExprList)
	return Alias{
		Name:  ident,
		Value: value,
	}
}

func (b *Builder) VisitExprStmt(exprStmt *par.ExprStmtContext) any {
	exprList := b.VisitExprList(exprStmt.ExprList().(*par.ExprListContext)).(ExprList)
	return ExprStmt{
		ExprList: exprList,
	}
}

func (b *Builder) VisitExprList(exprList *par.ExprListContext) any {
	es := ExprList{}
	for _, e := range exprList.AllExpr() {
		es = append(es, b.VisitExpr(e.(*par.ExprContext)).(Expr))
	}
	return es
}

func (b *Builder) VisitExpr(expr *par.ExprContext) any {
	switch expr := expr.GetChild(0).(type) {
	case *par.LambdaContext:
		return b.VisitLambda(expr)
	case *par.LiteralContext:
		return b.VisitLiteral(expr)
	case *par.IdentContext:
		return b.VisitIdent(expr)
	case *par.OperatorContext:
		return b.VisitOperator(expr)
	default:
		panic(b.unknown("ExprContext", expr.(antlr.ParseTree)))
	}
}

func (b *Builder) VisitOperator(operator *par.OperatorContext) any {
	return Ident(operator.GetText())
}

func (b *Builder) VisitLambda(lambda *par.LambdaContext) any {
	args := b.VisitIdentList(lambda.GetArgs().(*par.IdentListContext)).(IdentList)
	body := b.VisitExprList(lambda.GetBody().(*par.ExprListContext)).(ExprList)
	return Lambda{
		Args: args,
		Body: body,
	}
}

func (b *Builder) VisitIdentList(idents *par.IdentListContext) any {
	is := IdentList{}
	for _, ident := range idents.AllIdent() {
		is = append(is, b.VisitIdent(ident.(*par.IdentContext)).(Ident))
	}
	return is
}

func (b *Builder) VisitIdent(ident *par.IdentContext) any {
	return Ident(ident.GetText())
}

func (b *Builder) VisitLiteral(literal *par.LiteralContext) any {
	switch child := literal.GetChild(0).(type) {
	case *par.LiteralPrimitiveContext:
		return b.VisitLiteralPrimitive(child)
	case *par.LiteralListContext:
		return b.VisitLiteralList(child)
	case *par.LiteralMapContext:
		return b.VisitLiteralMap(child)
	default:
		panic(b.unknown("LiteralContext", child.(antlr.ParseTree)))
	}
}

func (b *Builder) VisitLiteralList(ctx *par.LiteralListContext) any {
	l := []Expr{}
	for _, expr := range ctx.AllExpr() {
		l = append(l, b.VisitExpr(expr.(*par.ExprContext)).(Expr))
	}
	return LiteralList{Value: l}
}

func (b *Builder) VisitLiteralMap(ctx *par.LiteralMapContext) any {
	m := []LiteralMapEntry{}
	for _, entry := range ctx.AllLiteralMapEntry() {
		m = append(m, b.VisitLiteralMapEntry(entry.(*par.LiteralMapEntryContext)).(LiteralMapEntry))
	}
	return LiteralMap{Value: m}
}

// VisitLiteralMapEntry implements par.GolflangVisitor.
func (b *Builder) VisitLiteralMapEntry(ctx *par.LiteralMapEntryContext) any {
	k := b.VisitExpr(ctx.GetKey().(*par.ExprContext)).(Expr)
	v := b.VisitExpr(ctx.GetValue().(*par.ExprContext)).(Expr)
	return LiteralMapEntry{K: k, V: v}
}

// VisitLiteralPrimitive implements par.GolflangVisitor.
func (b *Builder) VisitLiteralPrimitive(ctx *par.LiteralPrimitiveContext) any {
	return LiteralPrimitive{Value: b.primitiveLiteral(ctx)}
}

func (b *Builder) primitiveLiteral(c par.ILiteralPrimitiveContext) obj.Obj {
	switch {
	case c.INT() != nil:
		return b.int(c)
	case c.DEC() != nil:
		return b.dec(c)
	case c.STR() != nil:
		return b.str(c)
	case c.TRUE() != nil:
		return b.bool(c)
	case c.FALSE() != nil:
		return b.bool(c)
	default:
		panic(fmt.Errorf("%T: unknown primitiveLiteral: %s", b, c.ToStringTree(b.parser.RuleNames, b.parser)))
	}
}

func (b *Builder) int(c par.ILiteralPrimitiveContext) obj.Int {
	i, err := strconv.ParseInt(c.GetText(), 10, 64)
	if err != nil {
		panic(err)
	}
	return obj.Int(i)
}

func (b *Builder) dec(c par.ILiteralPrimitiveContext) obj.Dec {
	f, err := strconv.ParseFloat(c.GetText(), 64)
	if err != nil {
		panic(err)
	}
	return obj.Dec(f)
}

func (b *Builder) str(c par.ILiteralPrimitiveContext) obj.Str {
	v := c.GetText()
	return obj.NewStr(v)
}

func (b *Builder) bool(c par.ILiteralPrimitiveContext) obj.Bool {
	v, err := strconv.ParseBool(c.GetText())
	if err != nil {
		panic(err)
	}
	return obj.Bool(v)
}

// VisitChildren implements par.GolflangVisitor.
func (b *Builder) VisitChildren(node antlr.RuleNode) interface{} {
	return b.BaseGolflangVisitor.VisitChildren(node)
}

// VisitErrorNode implements par.GolflangVisitor.
func (b *Builder) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return b.BaseGolflangVisitor.VisitErrorNode(node)
}

// VisitTerminal implements par.GolflangVisitor.
func (b *Builder) VisitTerminal(node antlr.TerminalNode) interface{} {
	return b.BaseGolflangVisitor.VisitTerminal(node)
}
