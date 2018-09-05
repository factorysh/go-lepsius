package tick

import (
	"fmt"
	"testing"
)

func TestFrom(t *testing.T) {
	i := NewInput()
	s := i.FromStdin()
	fmt.Println(s)
}
