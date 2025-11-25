package interpreter

import (
	"bigyihsuan/golflang/internal/obj"
	"bigyihsuan/golflang/internal/util"
	"fmt"
)

func Zip(i *Interpreter) (obj.Obj, error) {
	args, err := i.getArgs(2, "zip")
	if err != nil {
		return nil, err
	}
	l, r := args[1], args[0]

	if l.Kind() != r.Kind() {
		return nil, ErrMismatchedKinds{l, r}
	}

	switch l.Kind() {
	case obj.ObjKindList:
		return obj.NewList(util.SliceZip(l.(obj.List), r.(obj.List))...), nil
	case obj.ObjKindStr:
		return obj.NewStr(string(util.SliceZip([]rune(l.(obj.Str)), []rune(r.(obj.Str))))), nil
	default:
		return nil, fmt.Errorf("unknown types for zip: %s, %s", l.Kind(), r.Kind())
	}
}
