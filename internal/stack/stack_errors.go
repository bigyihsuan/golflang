package stack

import "fmt"

type ErrNotEnoughStackValues struct {
	Want, Need int
}

func (e ErrNotEnoughStackValues) Error() string {
	return fmt.Sprintf("popped empty stack: want %d, need %d more", e.Want, e.Need)
}

type ErrPoppedEmptyStack struct{}

func (e ErrPoppedEmptyStack) Error() string {
	return "popped empty stack"
}
