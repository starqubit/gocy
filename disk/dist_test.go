package disk

import (
	"log"
	"testing"
	"time"
)

// 测试当前计算机C盘磁盘使用状况
func TestDisk(t *testing.T) {
	sTime := time.Now().UnixNano()
	available, free, total, err := DiskUsage("C:")
	if err != nil {
		t.Error(err)
		return
	}
	eTime := time.Now().UnixNano()

	log.Printf("Available  %dmb", available/1024/1024.0)
	log.Printf("Free       %dmb", free/1024/1024.0)
	log.Printf("Total      %dmb", total/1024/1024.0)
	log.Println(sTime, eTime-sTime, (eTime-sTime)/10e6)
}
