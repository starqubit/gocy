// +build !windows

package common

import (
	"os"
	"syscall"
	"time"
)

// 文件状态信息
type FileStat struct {
	filePath string
	FileInfo os.FileInfo
}

// 更新文件访问时间为当前时间
func RefreshAccessTime(filePath string) error {
	fileInfo, err := Stat(filePath)
	if err != nil {
		return err
	}
	err = fileInfo.ChangeAccessTime(time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}

// 查询文件状态信息
func Stat(filePath string) (*FileStat, error) {
	s := new(FileStat)
	s.filePath = filePath
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	s.FileInfo = fileInfo
	return s, nil
}

// 获取文件的访问时间, 返回时间戳
func (fs *FileStat) GetAccessTime() int64 {
	stat_t := fs.FileInfo.Sys().(*syscall.Stat_t)
	return int64(stat_t.Atim.Sec)
}

// 获取文件的创建时间, 返回时间戳
func (fs *FileStat) GetCreateTime() int64 {
	stat_t := fs.FileInfo.Sys().(*syscall.Stat_t)
	return int64(stat_t.Ctim.Sec)
}

// 获取文件的修改时间, 返回时间戳
func (fs *FileStat) GetWriteTime() int64 {
	stat_t := fs.FileInfo.Sys().(*syscall.Stat_t)
	return int64(stat_t.Mtim.Sec)
}

// 修改文件访问时间
func (fs *FileStat) ChangeAccessTime(timestamp int64) error {
	err := os.Chtimes(fs.filePath, time.Unix(timestamp, 0), time.Unix(fs.GetWriteTime(), 0))
	if err != nil {
		return err
	}
	return nil
}
