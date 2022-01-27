package R

import (
	"log"
	"testing"
	"time"
)

func TestR(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	InitServer("R_test", "http://127.0.0.1:29721/report?token=123sdffd6t536")
	for i := 0; i < 5; i++ {
		Debug("测试测试", i)
		time.Sleep(time.Second * 2)
	}
	time.Sleep(time.Second * 10)
}
