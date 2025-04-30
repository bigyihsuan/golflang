package interpreter

import (
	"bigyihsuan/golflang/internal/ast"
	"bigyihsuan/golflang/internal/obj"
	"bigyihsuan/golflang/internal/util"
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
		panic(fmt.Errorf("%T: unknown Stmt %T: %s", g, stmt, stmt.String()))
	}
}

func (g *Interpreter) VisitAlias(alias ast.Alias) {
	name := g.VisitIdent(alias.Name)
	value := g.VisitExpr(alias.Value)
	g.currentScope.SetAlias(name, value)
}

func (g *Interpreter) VisitExprStmt(expr ast.Expr) {
	value := g.VisitExpr(expr)
	g.queue.Push(g.EvalObj(value))
}

func (g *Interpreter) VisitExpr(expr ast.Expr) obj.Obj {
	switch expr := expr.(type) {
	case ast.Lit:
		return g.VisitLit(expr)
	case ast.Ident:
		return g.VisitIdent(expr)
	case ast.Lambda:
		return g.VisitLambda(expr)
	case ast.Call:
		return g.VisitCall(expr)
	default:
		panic(fmt.Errorf("%T: unknown Expr %T: %s", g, expr, expr.String()))
	}
}

func (g *Interpreter) VisitCall(expr ast.Call) obj.Obj {
	name := g.VisitIdent(expr.Name)
	fn := g.GetAlias(name)
	// set up by pushing arguments to the queue
	beforeLen := g.queue.Len()
	for _, arg := range expr.Args {
		g.Visit(arg)
	}
	// move the arguments to the front
	g.queue.Rotate(beforeLen)
	return g.EvalObj(fn)
}

func (g *Interpreter) VisitLambda(expr ast.Lambda) Lambda {
	args := util.SliceMap(expr.Args, func(i ast.Ident) obj.Ident { return obj.Ident(i) })
	body := expr.Body
	return Lambda{
		Args: args,
		Body: body,
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
		panic(fmt.Errorf("%T: unknown Lit %T: %s", g, lit, lit.String()))
	}
}

func (g *Interpreter) VisitIdent(ident ast.Ident) obj.Ident {
	return obj.Ident(ident)
}
