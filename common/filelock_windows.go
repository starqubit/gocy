// +build windows

package common

//文件锁
type fileLock struct {
}

func NewFileLock(dir string) *fileLock {
	return &fileLock{}
}

//加锁
func (l *fileLock) Lock() error {
	return nil
}

//释放锁
func (l *fileLock) Unlock() error {
	return nil
}
