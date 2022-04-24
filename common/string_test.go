package common

import (
	"testing"
)

func TestSm3(t *testing.T) {
	buf := []byte(`123`)
	t.Log(Sm3Hex(buf))
}
