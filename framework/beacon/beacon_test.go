package beacon

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	src := "N0CALL"
	cmt := "Test"
	want := fmt.Sprintf("%s>BEACON: %s", src, cmt)

	b := Beacon{
		Src:     src,
		Comment: cmt,
	}

	packet := b.String()
	if want != packet {
		t.Fatalf(`Expected %s, got %s`, want, packet)
	}
}
