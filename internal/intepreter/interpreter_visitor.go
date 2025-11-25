package interpreter

import (
	"bigyihsuan/golflang/internal/ast"
	"bigyihsuan/golflang/internal/obj"
	"bigyihsuan/golflang/internal/stack"
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
			return fmt.Errorf("prog: %w", err)
		}
	}
	return nil
}

func (i *Interpreter) VisitStmt(stmt ast.Stmt) error {
	switch stmt := stmt.(type) {
	case ast.Alias:
		return i.VisitAlias(stmt)
	case ast.ExprStmt:
		return i.VisitExprStmt(stmt)
	default:
		return fmt.Errorf("stmt: unknown Stmt %T: %s", stmt, stmt.String())
	}
}

func (i *Interpreter) VisitAlias(alias ast.Alias) error {
	name, err := i.VisitIdent(alias.Name)
	if err != nil {
		return fmt.Errorf("alias name: %w", err)
	}
	value := NewAliasValue(alias.Value)
	i.currentScope.SetAlias(name, value)
	return nil
}

func (i *Interpreter) VisitExprStmt(exprStmt ast.ExprStmt) error {
	_, err := i.EvalExprList(exprStmt.ExprList)
	if err != nil {
		return fmt.Errorf("exprStmt: %w", err)
	}
	// i.pushValue(value)
	return nil
}

func (i *Interpreter) EvalExprList(exprList ast.ExprList) (obj.Obj, error) {
	var lastValue obj.Obj
	for e := range util.Reversed(exprList) {
		value, err := i.EvalExpr(e)
		if err != nil {
			return value, fmt.Errorf("exprlist: %w", err)
		}
		i.pushValue(value)
		lastValue = value
	}
	return lastValue, nil
}

func (i *Interpreter) EvalExprListNoPush(exprList ast.ExprList) (obj.Obj, error) {
	var lastValue obj.Obj
	for e := range util.Reversed(exprList) {
		value, err := i.EvalExpr(e)
		if err != nil {
			return value, fmt.Errorf("exprlist: %w", err)
		}
		lastValue = value
	}
	return lastValue, nil
}

func (i *Interpreter) EvalExpr(expr ast.Expr) (obj.Obj, error) {
	switch expr := expr.(type) {
	case ast.IfThenElse:
		return i.EvalIfThenElse(expr)
	case ast.Ident:
		return i.EvalIdent(expr)
	case ast.Lambda:
		return i.EvalLambda(expr)
	case ast.Lit:
		return i.EvalLit(expr)
	default:
		return nil, fmt.Errorf("expr: unknown Expr %T %s", expr, expr.String())
	}
}

func (i *Interpreter) EvalIfThenElse(expr ast.IfThenElse) (obj.Obj, error) {
	condObj, err := i.EvalExprList(expr.Cond)
	if err != nil {
		return condObj, fmt.Errorf("ifThenElse: %w", err)
	}
	if condObj.Bool() {
		return i.EvalExprList(expr.Then)
	} else {
		return i.EvalExprList(expr.Else)
	}
}

func (i *Interpreter) VisitLambda(expr ast.Lambda) (obj.Obj, error) {
	return Lambda{
		Args: expr.Args,
		Body: expr.Body,
	}, nil
}

func (i *Interpreter) VisitIdent(name ast.Ident) (obj.Ident, error) {
	return obj.Ident(string(name)), nil
}

func (i *Interpreter) EvalIdent(ident ast.Ident) (obj.Obj, error) {
	value, err := i.getIdent(obj.Ident(ident))
	if err != nil {
		return value, fmt.Errorf("eval ident: %w", err)
	}
	switch value := value.(type) {
	case BuiltinFunc:
		return value(i)
	case AliasValue: // lazy evaluation of aliases
		return i.EvalExprList(value.ExprList)
	default:
		return value, nil
		// return nil, fmt.Errorf("eval ident: unknown value %T %s for Ident %s", value, value.String(), ident)
	}
}

func (i *Interpreter) EvalLambda(expr ast.Lambda) (obj.Obj, error) {
	// set up lambda scope
	lambdaScope := i.currentScope.Child()
	i.currentScope = &lambdaScope
	defer func() {
		i.currentScope = i.currentScope.Parent
	}()

	// set up arg variables
	for _, arg := range expr.Args {
		argObj, err := i.VisitIdent(arg)
		if err != nil {
			return nil, fmt.Errorf("lambda arg: %w", err)
		}
		argVal, ok := i.stack.Pop()
		if !ok {
			return nil, fmt.Errorf("lambda arg: %w", stack.ErrPoppedEmptyStack{})
		}
		i.currentScope.SetAlias(argObj, argVal)
	}

	// run lambda body
	return i.EvalExprList(expr.Body)
}

func (i *Interpreter) EvalLit(lit ast.Lit) (obj.Obj, error) {
	switch lit := lit.(type) {
	case ast.LiteralPrimitive:
		return lit.Value, nil
	case ast.LiteralList:
		return i.EvalLiteralList(lit)
	case ast.LiteralMap:
		return i.EvalLiteralMap(lit)
	default:
		return nil, fmt.Errorf("lit: unknown Lit %T %s", lit, lit.String())
	}
}

func (i *Interpreter) EvalLiteralList(lit ast.LiteralList) (obj.Obj, error) {
	list := obj.ZeroList()
	for _, e := range lit.Value {
		v, err := i.EvalExpr(e)
		if err != nil {
			return list, fmt.Errorf("list: %w", err)
		}
		list = append(list, v)
	}
	return list, nil
}

func (i *Interpreter) EvalLiteralMap(lit ast.LiteralMap) (obj.Obj, error) {
	mapObj := obj.ZeroMap()
	for _, e := range lit.Value {
		k, err := i.EvalExpr(e.K)
		if err != nil {
			return mapObj, fmt.Errorf("map key: %w", err)
		}
		v, err := i.EvalExpr(e.V)
		if err != nil {
			return mapObj, fmt.Errorf("map value: %w", err)
		}
		mapObj.Set(k, v)
	}
	return mapObj, nil
}
