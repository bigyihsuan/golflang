package interpreter

import (
	"bigyihsuan/golflang/internal/ast"
	"bigyihsuan/golflang/internal/obj"
	"bigyihsuan/golflang/internal/scope"
	"bigyihsuan/golflang/internal/util"
	"fmt"
)

func (i *Interpreter) Visit(node ast.Node) error {
	switch node := node.(type) {
	case ast.Prog:
		return i.VisitProg(node)
	case ast.Stmt:
		return i.VisitStmt(node)
	default:
		panic(fmt.Errorf("unknown Node %T: %s", node, node.String()))
	}
}

func (i *Interpreter) VisitProg(prog ast.Prog) error {
	for _, stmt := range prog.Stmts {
		err := i.VisitStmt(stmt)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *Interpreter) VisitStmt(stmt ast.Stmt) error {
	switch stmt := stmt.(type) {
	case ast.Alias:
		return i.VisitAlias(stmt)
	case ast.Expr:
		return i.VisitExprStmt(stmt)
	default:
		panic(fmt.Errorf("%T: unimplemented Stmt %T: %s", i, stmt, stmt.String()))
	}
}

func (i *Interpreter) VisitAlias(alias ast.Alias) error {
	name := i.VisitIdent(alias.Name)
	value, err := i.VisitExpr(alias.Value)
	if err != nil {
		return err
	}
	i.currentScope.SetAlias(name, value)
	return nil
}

func (i *Interpreter) VisitExprStmt(expr ast.Expr) error {
	value, err := i.VisitExpr(expr)
	if err != nil {
		return err
	}
	v, err := i.EvalObj(value)
	if err != nil {
		return err
	}
	if v.Kind() != obj.ObjKindNone {
		i.stack.Push(v)
	}
	return nil
}

func (i *Interpreter) VisitExpr(expr ast.Expr) (obj.Obj, error) {
	switch expr := expr.(type) {
	case ast.Lit:
		return i.VisitLit(expr)
	case ast.Ident:
		return i.VisitIdent(expr), nil
	case ast.Lambda:
		return i.VisitLambda(expr), nil
	case ast.Call:
		return i.VisitCall(expr)
	default:
		panic(fmt.Errorf("%T: unimplemented Expr %T: %s", i, expr, expr.String()))
	}
}

func (i *Interpreter) VisitCall(expr ast.Call) (obj.Obj, error) {
	// set up by pushing arguments to the stack
	for arg := range util.Reversed(expr.Args) {
		i.Visit(arg)
	}

	name := i.VisitIdent(expr.Name)
	// check for program-defined funcs first
	if fn, err := i.currentScope.GetAlias(name); err == nil {
		return i.EvalObj(fn)
	}
	// check for builtin
	f, ok := i.builtins.Get(name)
	if !ok {
		return nil, fmt.Errorf("calling function: %w", scope.ErrUnknownAlias{Name: name.String()})
	}
	return f(i)
}

func (i *Interpreter) VisitLambda(expr ast.Lambda) Lambda {
	args := util.SliceMap(expr.Args, func(i ast.Ident) obj.Ident { return obj.Ident(i) })
	body := expr.Body
	return Lambda{
		Args: args,
		Body: body,
	}
}

func (i *Interpreter) VisitLit(lit ast.Lit) (obj.Obj, error) {
	switch lit := lit.(type) {
	case ast.LiteralPrimitive:
		return lit.Value, nil
	case ast.LiteralList:
		values := []obj.Obj{}
		for _, e := range lit.Value {
			v, err := i.VisitExpr(e)
			if err != nil {
				return nil, fmt.Errorf("building list: %w", err)
			}
			values = append(values, v)
		}
		return obj.NewList(values...), nil
	case ast.LiteralMap:
		values := []obj.MapEntry{}
		for _, entry := range lit.Value {
			k, err := i.VisitExpr(entry.K)
			if err != nil {
				return nil, fmt.Errorf("building map key: %w", err)
			}
			v, err := i.VisitExpr(entry.V)
			if err != nil {
				return nil, fmt.Errorf("building map value: %w", err)
			}
			values = append(values, obj.MapEntry{K: k, V: v})
		}
		return obj.MapFromEntries(values...), nil
	default:
		panic(fmt.Errorf("%T: unimplemented Lit %T: %s", i, lit, lit.String()))
	}
}

func (i *Interpreter) VisitIdent(ident ast.Ident) obj.Ident {
	return obj.Ident(ident)
}
