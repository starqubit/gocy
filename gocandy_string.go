package gocandy

func Substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)
	if end == 0 || end >= length {
		end = length
	}
	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	var substring = ""
	for i := start; i < length; i++ {
		substring += string(r[i])
	}

	return substring
}
