package main

import (
	"flag"

	interpreter "bigyihsuan/golflang/internal/intepreter"
)

var (
	filename = flag.String("f", "", "input file name")
)

func main() {
	flag.Parse()

	interpreter, err := interpreter.New(*filename)
	if err != nil {
		panic(err)
	}
	err = interpreter.Parse()
	if err != nil {
		panic(err)
	}
}
