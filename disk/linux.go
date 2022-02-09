// +build !windows

package disk

import "syscall"

func getUsage(usagePath string) (int64, int64, int64, error) {

	fs := syscall.Statfs_t{}
	err := syscall.Statfs(usagePath, &fs)
	if err != nil {
		return 0, 0, 0, err
	}
	total := fs.Blocks * uint64(fs.Bsize)
	free := fs.Bfree * uint64(fs.Bsize)
	return int64(free), int64(free), int64(total), nil
}
