package interpreter

import (
	"bigyihsuan/golflang/internal/obj"
	"fmt"
)

func Plus(i *Interpreter) (obj.Obj, error) {
	args, err := i.getArgs(2, "Plus")
	if err != nil {
		return nil, err
	}
	l, r := args[1], args[0]

	if l.Kind() != r.Kind() {
		return nil, ErrMismatchedKinds{l, r}
	}

	switch l.Kind() {
	case obj.ObjKindInt:
		return obj.Int(int64(l.(obj.Int)) + int64(r.(obj.Int))), nil
	case obj.ObjKindDec:
		return obj.Dec(float64(l.(obj.Dec)) + float64(r.(obj.Dec))), nil
	case obj.ObjKindStr:
		return obj.Str(string(l.(obj.Str)) + string(r.(obj.Str))), nil
	case obj.ObjKindList:
		return append(l.(obj.List), r.(obj.List)...), nil
	case obj.ObjKindMap:
		l := l.(obj.Map)
		r := r.(obj.Map)
		for e := range r.Entries() {
			l.SetEntry(e)
		}
		return l, nil
	default:
		return nil, fmt.Errorf("unknown type for +: %s", l.Kind())
	}
}

func Minus(i *Interpreter) (obj.Obj, error) {
	args, err := i.getArgs(2, "Minus")
	if err != nil {
		return nil, err
	}
	l, r := args[1], args[0]

	if l.Kind() != r.Kind() {
		return nil, ErrMismatchedKinds{l, r}
	}

	switch l.Kind() {
	case obj.ObjKindInt:
		return obj.Int(int64(l.(obj.Int)) - int64(r.(obj.Int))), nil
	case obj.ObjKindDec:
		return obj.Dec(float64(l.(obj.Dec)) - float64(r.(obj.Dec))), nil
	// case obj.ObjKindStr, obj.ObjKindList, obj.ObjKindMap:
	default:
		return nil, fmt.Errorf("unknown type for -: %s", l.Kind())
	}
}

func Underscore(i *Interpreter) (obj.Obj, error) {
	args, err := i.getArgs(1, "Underscore")
	if err != nil {
		return nil, err
	}
	v := args[0]

	switch v.Kind() {
	case obj.ObjKindInt:
		return obj.Int(-int64(v.(obj.Int))), nil
	case obj.ObjKindDec:
		return obj.Dec(-float64(v.(obj.Dec))), nil
	// case obj.ObjKindStr, obj.ObjKindList, obj.ObjKindMap:
	// 	return nil, fmt.Errorf("invalid type for _: %s", v.Kind())
	default:
		return nil, fmt.Errorf("unknown type for _: %s", v.Kind())
	}
}

func GreaterThan(i *Interpreter) (obj.Obj, error) {
	args, err := i.getArgs(2, "GreaterThan")
	if err != nil {
		return nil, err
	}
	l, r := args[1], args[0]

	if l.Kind() != r.Kind() {
		return nil, ErrMismatchedKinds{l, r}
	}

	switch l.Kind() {
	case obj.ObjKindInt:
		return obj.Bool(int64(l.(obj.Int)) > int64(r.(obj.Int))), nil
	case obj.ObjKindDec:
		return obj.Bool(float64(l.(obj.Dec)) > float64(r.(obj.Dec))), nil
	case obj.ObjKindStr:
		return obj.Bool(string(l.(obj.Str)) > string(r.(obj.Str))), nil
	default:
		return nil, fmt.Errorf("unknown type for +: %s", l.Kind())
	}
}

func LessThan(i *Interpreter) (obj.Obj, error) {
	args, err := i.getArgs(2, "LessThan")
	if err != nil {
		return nil, err
	}
	l, r := args[1], args[0]

	if l.Kind() != r.Kind() {
		return nil, ErrMismatchedKinds{l, r}
	}

	switch l.Kind() {
	case obj.ObjKindInt:
		return obj.Bool(int64(l.(obj.Int)) < int64(r.(obj.Int))), nil
	case obj.ObjKindDec:
		return obj.Bool(float64(l.(obj.Dec)) < float64(r.(obj.Dec))), nil
	case obj.ObjKindStr:
		return obj.Bool(string(l.(obj.Str)) < string(r.(obj.Str))), nil
	default:
		return nil, fmt.Errorf("unknown type for +: %s", l.Kind())
	}
}
