package disk

/*
// 获取路径usagePath的磁盘使用状态
return: 可用空间，剩余空间，总空间
*/
func DiskUsage(usagePath string) (int64, int64, int64, error) {
	return getUsage(usagePath)
}
