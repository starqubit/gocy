package common

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
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
