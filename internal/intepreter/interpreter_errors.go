package interpreter

import "fmt"

type ErrNotEnoughStackValues struct {
	Want, Need int
}

func (e ErrNotEnoughStackValues) Error() string {
	return fmt.Sprintf("not enough items on the stack: want %d, need %d more", e.Want, e.Need)
}
