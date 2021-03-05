// Package beacon represents a single APRS beacon frame.
package beacon

import (
	"fmt"
)

// Beacon represents an APRS beacon frame.
type Beacon struct {
	Src     string
	Comment string
}

// String generates a raw APRS beacon frame in string form.
func (b *Beacon) String() string {
	return fmt.Sprintf("%s>BEACON: %s", b.Src, b.Comment)
}
