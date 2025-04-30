package interpreter

import (
	"bigyihsuan/golflang/internal/ast"
	"bigyihsuan/golflang/internal/obj"
	"fmt"
)

func (g *Interpreter) Visit(node ast.Node) {
	switch node := node.(type) {
	case ast.Prog:
		g.VisitProg(node)
	case ast.Stmt:
		g.VisitStmt(node)
	case ast.Expr:
		g.VisitExpr(node)
	default:
		panic(fmt.Errorf("unknown Node %T: %s", node, node.String()))
	}
}

func (g *Interpreter) VisitProg(prog ast.Prog) {
	for _, stmt := range prog.Stmts {
		g.VisitStmt(stmt)
	}
}

func (g *Interpreter) VisitStmt(stmt ast.Stmt) {
	switch stmt := stmt.(type) {
	case ast.Alias:
		g.VisitAlias(stmt)
	case ast.Expr:
		g.VisitExprStmt(stmt)
	default:
		panic(fmt.Errorf("unknown Stmt %T: %s", stmt, stmt.String()))
	}
}

func (g *Interpreter) VisitAlias(alias ast.Alias) {
	name := AliasName(g.VisitIdent(alias.Name).String())
	value := g.VisitExpr(alias.Value)
	g.aliases[name] = value
}

func (g *Interpreter) VisitExprStmt(expr ast.Expr) {
	value := g.VisitExpr(expr)
	g.queue.Push(value)
}

func (g *Interpreter) VisitExpr(expr ast.Expr) obj.Obj {
	switch expr := expr.(type) {
	case ast.Lit:
		return g.VisitLit(expr)
	case ast.Ident:
		return g.VisitIdent(expr)
	default:
		panic(fmt.Errorf("unknown Expr %T: %s", expr, expr.String()))
	}
}

func (g *Interpreter) VisitLit(lit ast.Lit) obj.Obj {
	switch lit := lit.(type) {
	case ast.LiteralPrimitive:
		return lit.Value
	case ast.LiteralList:
		values := []obj.Obj{}
		for _, e := range lit.Value {
			values = append(values, g.VisitExpr(e))
		}
		return obj.NewList(values...)
	case ast.LiteralMap:
		values := []obj.MapEntry{}
		for _, entry := range lit.Value {
			k := g.VisitExpr(entry.K)
			v := g.VisitExpr(entry.V)
			values = append(values, obj.MapEntry{K: k, V: v})
		}
		return obj.MapFromEntries(values...)
	default:
		panic(fmt.Errorf("unknown Lit %T: %s", lit, lit.String()))
	}
}

func (g *Interpreter) VisitIdent(ident ast.Ident) obj.Ident {
	return obj.Ident(ident)
}
