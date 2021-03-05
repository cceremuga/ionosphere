package beacon

import (
	"fmt"
)

type Beacon struct {
	Src     string
	Comment string
}

func (b *Beacon) String() string {
	return fmt.Sprintf("%s>BEACON: %s", b.Src, b.Comment)
}
