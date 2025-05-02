package interpreter

import (
	"bigyihsuan/golflang/internal/obj"
	"fmt"
)

type BuiltinFunc func(*Interpreter) (obj.Obj, error)

type builtins map[string]BuiltinFunc

func initBuiltins() builtins {
	return builtins{
		"pop":     Pop,
		"print":   Print,
		"println": Println,
	}
}

func (b builtins) Get(name obj.Ident) (BuiltinFunc, bool) {
	f, ok := b[string(name)]
	return f, ok
}

func Pop(i *Interpreter) (obj.Obj, error) {
	_, ok := i.stack.Pop()
	if !ok {
		return nil, ErrNotEnoughStackValues{Want: 1, Need: 1}
	}
	return obj.ZeroNone(), nil
}

func Print(i *Interpreter) (obj.Obj, error) {
	v, ok := i.stack.Pop()
	if !ok {
		return nil, ErrNotEnoughStackValues{Want: 1, Need: 1}
	}
	fmt.Print(v)
	return obj.ZeroNone(), nil
}

func Println(i *Interpreter) (obj.Obj, error) {
	v, ok := i.stack.Pop()
	if !ok {
		return nil, ErrNotEnoughStackValues{Want: 1, Need: 1}
	}
	fmt.Println(v)
	return obj.ZeroNone(), nil
}
