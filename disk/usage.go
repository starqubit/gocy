package disk

func DiskUsage(usagePath string) (int64, int64, int64, error) {
	return getUsage(usagePath)
}
