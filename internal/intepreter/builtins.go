package interpreter

import (
	"bigyihsuan/golflang/internal/obj"
	"bigyihsuan/golflang/internal/stack"
	"fmt"
)

type builtins map[string]BuiltinFunc

func initBuiltins() builtins {
	return builtins{
		"pop":     Pop,
		"print":   Print,
		"println": Println,
		"+":       Plus,
	}
}

func (b builtins) Get(name obj.Ident) (BuiltinFunc, bool) {
	f, ok := b[string(name)]
	return f, ok
}

func Pop(i *Interpreter) (obj.Obj, error) {
	_, ok := i.stack.Pop()
	if !ok {
		return nil, stack.ErrPoppedEmptyStack{}
	}
	return obj.ZeroNone(), nil
}

func Print(i *Interpreter) (obj.Obj, error) {
	v, ok := i.stack.Pop()
	if !ok {
		return nil, stack.ErrPoppedEmptyStack{}
	}
	fmt.Print(v)
	return obj.ZeroNone(), nil
}

func Println(i *Interpreter) (obj.Obj, error) {
	v, ok := i.stack.Pop()
	if !ok {
		return nil, stack.ErrPoppedEmptyStack{}
	}
	fmt.Println(v)
	return obj.ZeroNone(), nil
}
