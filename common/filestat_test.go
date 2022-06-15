package common

import (
	"log"
	"testing"
	"time"
)

// 测试文件修改功能
func TestFileStat(t *testing.T) {
	filePath := `C:\Users\Administrator\Videos\video d216fb22a2cc3bd1.mp4`
	fileInfo, err := Stat(filePath)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(fileInfo.GetCreateTime())
	t.Log(fileInfo.GetWriteTime())
	t.Log(fileInfo.GetAccessTime())

	fileInfo.ChangeAccessTime(time.Now().Unix())

	fileInfo, err = Stat(filePath)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(fileInfo.GetCreateTime())
	t.Log(fileInfo.GetWriteTime())
	t.Log(fileInfo.GetAccessTime())
}

func TestHash(t *testing.T) {
	filePath := `E:\video\1.mp4`
	log.Println(Md5FileHex(filePath))
}
