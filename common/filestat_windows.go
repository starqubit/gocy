// +build windows

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
	winFileAttr := fs.FileInfo.Sys().(*syscall.Win32FileAttributeData)
	return winFileAttr.LastAccessTime.Nanoseconds() / 1e9
}

// 获取文件的创建时间, 返回时间戳
func (fs *FileStat) GetCreateTime() int64 {
	winFileAttr := fs.FileInfo.Sys().(*syscall.Win32FileAttributeData)
	return winFileAttr.CreationTime.Nanoseconds() / 1e9
}

// 获取文件的修改时间, 返回时间戳
func (fs *FileStat) GetWriteTime() int64 {
	winFileAttr := fs.FileInfo.Sys().(*syscall.Win32FileAttributeData)
	return winFileAttr.LastWriteTime.Nanoseconds() / 1e9
}

// 修改文件访问时间
func (fs *FileStat) ChangeAccessTime(timestamp int64) error {
	err := os.Chtimes(fs.filePath, time.Unix(timestamp, 0), time.Unix(fs.GetWriteTime(), 0))
	if err != nil {
		return err
	}
	return nil
}
