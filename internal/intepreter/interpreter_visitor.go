package interpreter

import (
	"bigyihsuan/golflang/internal/ast"
	"bigyihsuan/golflang/internal/obj"
	"bigyihsuan/golflang/internal/scope"
	"bigyihsuan/golflang/internal/util"
	"fmt"
)

func (g *Interpreter) Visit(node ast.Node) error {
	switch node := node.(type) {
	case ast.Prog:
		return g.VisitProg(node)
	case ast.Stmt:
		return g.VisitStmt(node)
	default:
		panic(fmt.Errorf("unknown Node %T: %s", node, node.String()))
	}
}

func (g *Interpreter) VisitProg(prog ast.Prog) error {
	for _, stmt := range prog.Stmts {
		err := g.VisitStmt(stmt)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Interpreter) VisitStmt(stmt ast.Stmt) error {
	switch stmt := stmt.(type) {
	case ast.Alias:
		return g.VisitAlias(stmt)
	case ast.Expr:
		return g.VisitExprStmt(stmt)
	default:
		panic(fmt.Errorf("%T: unimplemented Stmt %T: %s", g, stmt, stmt.String()))
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
	if v.Kind() != obj.ObjKindNone {
		g.stack.Push(v)
	}
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
	// set up by pushing arguments to the stack
	for arg := range util.Reversed(expr.Args) {
		g.Visit(arg)
	}

	name := g.VisitIdent(expr.Name)
	// check for program-defined funcs first
	if fn, err := g.currentScope.GetAlias(name); err == nil {
		return g.EvalObj(fn)
	}
	// check for builtin
	f, ok := g.builtins.Get(name)
	if !ok {
		return nil, fmt.Errorf("calling function: %w", scope.ErrUnknownAlias{Name: name.String()})
	}
	return f(g)
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
		panic(fmt.Errorf("%T: unimplemented Lit %T: %s", g, lit, lit.String()))
	}
}

func (g *Interpreter) VisitIdent(ident ast.Ident) obj.Ident {
	return obj.Ident(ident)
}
