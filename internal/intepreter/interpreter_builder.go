package interpreter

type builder struct {
	filename         string
	debug, dumpStack bool
}

func Builder() *builder {
	return new(builder)
}

func (b *builder) Filename(filename string) *builder {
	b.filename = filename
	return b
}

func (b *builder) Debug(debug bool) *builder {
	b.debug = debug
	return b
}

func (b *builder) DumpStack(dumpStack bool) *builder {
	b.dumpStack = dumpStack
	return b
}

func (b builder) Build() (*Interpreter, error) {
	return newInterpreter(b)
}
