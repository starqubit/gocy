package common

import (
	"testing"
)

func TestSm3(t *testing.T) {
	buf := []byte(`123`)
	t.Log(Sm3Hex(buf))
}

// 生成随机字符串
func TestRandomString(t *testing.T) {
	r := RandomString(16)
	t.Log(r)
}
