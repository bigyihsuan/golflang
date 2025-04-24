package obj

type String string

func ZeroString() String {
	return String("")
}

// AsBool implements Obj.
func (s String) AsBool() bool {
	return s != ""
}

// Equal implements Obj.
func (s String) Equal(o Obj) bool {
	switch o.Kind() {
	case KindString:
		return s == o.(String)
	default:
		return false
	}
}

// Kind implements Obj.
func (s String) Kind() ObjKind {
	return KindString
}

// String implements Obj.
func (s String) String() string {
	return string(s)
}
