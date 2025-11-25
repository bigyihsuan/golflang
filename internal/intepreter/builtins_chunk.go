package interpreter

import (
	"bigyihsuan/golflang/internal/obj"
	"bigyihsuan/golflang/internal/util"
	"fmt"
	"slices"
	"strings"
)

func NChunks(i *Interpreter) (obj.Obj, error) {
	args, err := i.getArgs(2, "nChunks")
	if err != nil {
		return nil, err
	}
	l, r := args[1], args[0]

	if l.Kind() != obj.ObjKindList || r.Kind() != obj.ObjKindInt {
		return nil, ErrIncorrectArgumentTypes{
			"nChunks",
			ArgAcceptedTypes{l.Hash(): {obj.ObjKindList}, r.Hash(): {obj.ObjKindInt}},
		}
	}

	collection := l.(obj.List)
	chunkCount := r.(obj.Int)

	if int64(chunkCount) < 1 {
		return collection, nil
	}

	o := []obj.Obj{}
	for e := range slices.Chunk(collection, int(chunkCount)) {
		o = append(o, e)
	}

	return obj.NewList(o...), nil
}

func ChunkN(i *Interpreter) (obj.Obj, error) {
	args, err := i.getArgs(2, "chunkN")
	if err != nil {
		return nil, err
	}
	l, r := args[1], args[0]

	if l.Kind() != obj.ObjKindList || r.Kind() != obj.ObjKindInt {
		return nil, ErrIncorrectArgumentTypes{"chunkN", ArgAcceptedTypes{l.Hash(): {obj.ObjKindList}, r.Hash(): {obj.ObjKindInt}}}
	}

	collection := l.(obj.List)
	chunkSize := r.(obj.Int)

	if int64(chunkSize) < 1 {
		return collection, nil
	}

	o := []obj.Obj{}
	for i := 0; i < len(collection); i += int(chunkSize) {
		o = append(o, obj.NewList(collection[i:i+int(chunkSize)]))
	}

	return obj.NewList(o...), nil
}

func ChunkSame(i *Interpreter) (obj.Obj, error) {
	args, err := i.getArgs(1, "chunkSame")
	if err != nil {
		return nil, err
	}
	l := args[0]

	if l.Kind() != obj.ObjKindList && l.Kind() != obj.ObjKindStr {
		return nil, ErrIncorrectArgumentTypes{
			"chunkSame",
			ArgAcceptedTypes{l.Hash(): {obj.ObjKindList, obj.ObjKindStr}},
		}
	}

	var o []obj.Obj
	switch l.Kind() {
	case obj.ObjKindStr:
		collection := []rune(string(l.(obj.Str)))
		collection, counts := util.SliceCompactCount(collection)
		repeateds := util.SliceMap2(collection, counts,
			func(r rune, n int) string { return strings.Repeat(string([]rune{r}), n) })
		o = util.SliceMap(repeateds, func(s string) obj.Obj { return obj.Obj(obj.NewStr(s)) })

	case obj.ObjKindList:
		collection := l.(obj.List)
		collection, counts := util.SliceCompactCountFunc(collection, func(a, b obj.Obj) bool { return a.Equal(b) })
		repeateds := util.SliceMap2(collection, counts,
			func(o obj.Obj, n int) []obj.Obj { return slices.Repeat([]obj.Obj{o}, n) })
		o = util.SliceMap(repeateds, func(o []obj.Obj) obj.Obj { return obj.Obj(obj.NewList(o...)) })

	default:
		return nil, fmt.Errorf("chunkSame: invalid type %s", l.Kind())
	}
	return obj.NewList(o...), nil
}
