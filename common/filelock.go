// +build !windows

package common

import (
	"fmt"
	"os"
	"syscall"
)

//文件锁
type fileLock struct {
	dir string
	f   *os.File
}

func NewFileLock(dir string) *fileLock {
	return &fileLock{
		dir: dir,
	}
}

//加锁
func (l *fileLock) Lock() error {
	f, err := os.Open(l.dir)
	if err != nil {
		return err
	}
	l.f = f
	err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		return fmt.Errorf("cannot flock directory %s - %s", l.dir, err)
	}
	return nil
}

//释放锁
func (l *fileLock) Unlock() error {
	defer l.f.Close()
	return syscall.Flock(int(l.f.Fd()), syscall.LOCK_UN)
}
