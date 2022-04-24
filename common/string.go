package common

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"

	"github.com/starqubit/gocy/common/sm"
)

// 获取数据md5的hex编码
func Md5BufHex(buf []byte) string {
	hash := md5.New()
	_, err := hash.Write(buf)
	if err != nil {
		return "00000000000000000000000000000000"
	}
	return hex.EncodeToString(hash.Sum(nil))
}

// 删除字符串中的所有空格符
func TrimAllSpace(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "\t", "")
	return str
}

// 删除换行符
func TrimEnter(str string) string {
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\r", "")

	return str
}

// 是否空字符
func IsNullOrEmpty(str string) bool {
	str = TrimAllSpace(str)
	str = TrimEnter(str)
	return str == ""
}

// sha1 hex
func Sha1Hex(buf []byte) string {
	m := sha1.New()
	m.Write(buf)
	return hex.EncodeToString(m.Sum(nil))
}

// sm3 hex
func Sm3Hex(buf []byte) string {
	s := sm.Sm3Sum(buf)
	return hex.EncodeToString(s)
}

// 生成随机字符串
func RandomString(n int, allowedChars ...[]rune) string {
	var letters []rune
	if len(allowedChars) == 0 {
		letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	} else {
		letters = allowedChars[0]
	}
	b := make([]rune, n)
	for i := range b {
		rand.Seed(time.Now().UTC().UnixNano() + int64(i<<20))
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
