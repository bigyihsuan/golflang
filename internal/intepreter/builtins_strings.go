package interpreter

import (
	"bigyihsuan/golflang/internal/obj"
	"bigyihsuan/golflang/internal/util"
	"strings"
)

func Join(i *Interpreter) (obj.Obj, error) {
	args, err := i.getArgs(2, "join")
	if err != nil {
		return nil, err
	}
	l, r := args[1], args[0]
	if l.Kind() != obj.ObjKindList || r.Kind() != obj.ObjKindStr {
		return nil, ErrIncorrectArgumentTypes{"join", ArgAcceptedTypes{l.Hash(): {obj.ObjKindList}, r.Hash(): {obj.ObjKindStr}}}
	}

	collection := l.(obj.List)
	joiner := r.(obj.Str)

	return obj.NewStr(strings.Join(util.SliceMap(collection, func(o obj.Obj) string { return o.String() }), string(joiner))), nil
}
