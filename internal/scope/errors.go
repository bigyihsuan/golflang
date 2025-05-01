package scope

import "fmt"

type ErrUnknownAlias struct {
	Name string
}

func (e ErrUnknownAlias) Error() string {
	return fmt.Sprintf("unknown alias: %s", e.Name)
}
