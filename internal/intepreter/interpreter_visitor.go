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

func (g *Interpreter) VisitAlias(alias ast.Alias) error {
	name := g.VisitIdent(alias.Name)
	value, err := g.VisitExpr(alias.Value)
	if err != nil {
		return err
	}
	g.currentScope.SetAlias(name, value)
	return nil
}

func (g *Interpreter) VisitExprStmt(expr ast.Expr) error {
	value, err := g.VisitExpr(expr)
	if err != nil {
		return err
	}
	v, err := g.EvalObj(value)
	if err != nil {
		return err
	}
	g.stack.Push(v)
	return nil
}

func (g *Interpreter) VisitExpr(expr ast.Expr) (obj.Obj, error) {
	switch expr := expr.(type) {
	case ast.Lit:
		return g.VisitLit(expr)
	case ast.Ident:
		return g.VisitIdent(expr), nil
	case ast.Lambda:
		return g.VisitLambda(expr), nil
	case ast.Call:
		return g.VisitCall(expr)
	default:
		panic(fmt.Errorf("%T: unimplemented Expr %T: %s", g, expr, expr.String()))
	}
}

func (g *Interpreter) VisitCall(expr ast.Call) (obj.Obj, error) {
	name := g.VisitIdent(expr.Name)
	fn, err := g.currentScope.GetAlias(name)
	if err != nil {
		return nil, fmt.Errorf("calling function: %w", err)
	}
	// set up by pushing arguments to the stack
	for arg := range util.Reversed(expr.Args) {
		g.Visit(arg)
	}
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

func (g *Interpreter) VisitLit(lit ast.Lit) (obj.Obj, error) {
	switch lit := lit.(type) {
	case ast.LiteralPrimitive:
		return lit.Value, nil
	case ast.LiteralList:
		values := []obj.Obj{}
		for _, e := range lit.Value {
			v, err := g.VisitExpr(e)
			if err != nil {
				return nil, fmt.Errorf("building list: %w", err)
			}
			values = append(values, v)
		}
		return obj.NewList(values...), nil
	case ast.LiteralMap:
		values := []obj.MapEntry{}
		for _, entry := range lit.Value {
			k, err := g.VisitExpr(entry.K)
			if err != nil {
				return nil, fmt.Errorf("building map key: %w", err)
			}
			v, err := g.VisitExpr(entry.V)
			if err != nil {
				return nil, fmt.Errorf("building map value: %w", err)
			}
			values = append(values, obj.MapEntry{K: k, V: v})
		}
		return obj.MapFromEntries(values...), nil
	default:
		panic(fmt.Errorf("%T: unknown Lit %T: %s", g, lit, lit.String()))
	}
}

func (g *Interpreter) VisitIdent(ident ast.Ident) obj.Ident {
	return obj.Ident(ident)
}
