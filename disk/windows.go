// +build windows

package disk

import (
	"syscall"
	"unsafe"
)

func getUsage(usagePath string) (int64, int64, int64, error) {
	kernel32, err := syscall.LoadLibrary("Kernel32.dll")
	if err != nil {
		return 0, 0, 0, err
	}
	defer syscall.FreeLibrary(kernel32)
	GetDiskFreeSpaceEx, err := syscall.GetProcAddress(syscall.Handle(kernel32), "GetDiskFreeSpaceExW")

	if err != nil {
		return 0, 0, 0, err
	}

	available := int64(0)
	free := int64(0)
	total := int64(0)
	res, err := syscall.UTF16PtrFromString(usagePath)
	if err != nil {
		return 0, 0, 0, err
	}
	syscall.Syscall6(uintptr(GetDiskFreeSpaceEx), 4,
		uintptr(unsafe.Pointer(res)),
		uintptr(unsafe.Pointer(&available)),
		uintptr(unsafe.Pointer(&total)),
		uintptr(unsafe.Pointer(&free)), 0, 0)
	return available, free, total, nil
}
