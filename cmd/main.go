package main

import (
	"flag"
	"fmt"
	"os"

	interpreter "bigyihsuan/golflang/internal/intepreter"

	"github.com/fatih/color"
)

var (
	filename = flag.String("f", "", "input file name")
)

func main() {
	flag.Parse()

	interpreter, err := interpreter.New(*filename)
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
