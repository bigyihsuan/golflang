package main

import (
	"flag"
	"fmt"
	"os"

	interpreter "bigyihsuan/golflang/internal/intepreter"

	"github.com/fatih/color"
)

var (
	filename  = flag.String("f", "", "input file name")
	debug     = flag.Bool("v", false, "turn on debug printing")
	dumpStack = flag.Bool("d", false, "turn on dumping the stack on program exit")
)

func main() {
	flag.Parse()

	interpreter, err := interpreter.Builder().Filename(*filename).Debug(*debug).DumpStack(*dumpStack).Build()
	if err != nil {
		fmt.Println(color.HiRedString(err.Error()))
		os.Exit(-1)
	}
	err = interpreter.Run()
	if err != nil {
		fmt.Println(color.HiRedString(err.Error()))
		os.Exit(-1)
	}
}
