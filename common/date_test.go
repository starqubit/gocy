package common

import (
	"log"
	"testing"
)

func TestDate(t *testing.T) {
	t.Log("begin")
	strDate := "5分钟前"
	dstTime := Str2Time(strDate)
	log.Println(dstTime.Format("2006-01-02 15:04:05"))
}
