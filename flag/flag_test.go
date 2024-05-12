package flag_test

import (
	"testing"

	"github.com/chasecarlson1/go-bitflags/flag"
)

func TestMain(m *testing.M) {
	m.Run()
}
func TestFuncs(t *testing.T) {
	t.Run("Set()", func(t *testing.T) {
		var f = flag.New()
		f.Set(42)
		if f != 0b101010 {
			t.Fatal("var f = flag.New(), f.Set(42), f != 0b101010")
		}
	})
}
