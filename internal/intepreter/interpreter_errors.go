package interpreter

import "fmt"

type ErrUnknownAlias struct {
	Name string
}

func (e ErrUnknownAlias) Error() string {
	return fmt.Sprintf("unknown alias: %s", e.Name)
}

type ErrNotEnoughStackValues struct {
	Want, Need int
}

func (e ErrNotEnoughStackValues) Error() string {
	return fmt.Sprintf("not enough items on the stack: want %d, need %d more", e.Want, e.Need)
}
