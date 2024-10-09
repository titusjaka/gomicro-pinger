package commands

import (
	"fmt"
)

type Flags struct {
	Port int64 `kong:"optional,default='6066'"`
}

func (f Flags) Listen() string {
	return fmt.Sprintf(":%d", f.Port)
}
