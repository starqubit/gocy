package common

import (
	"fmt"
	"time"
)

// 格式化事件为标准的字符串
func FormatTimeStr(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// 将时长格式化为字幕时长格式 39.77 >> 00:00:39,770
func FormatSubtitleTime(duration float64) string {
	var h int
	var m int
	if duration >= 3600 {
		h = int(duration / 3600)
		duration -= float64(h * 3600)
	}
	if duration >= 60 {
		m = int(duration / 60)
		duration -= float64(m * 60)
	}
	timeVar := fmt.Sprintf("%02d:%02d:%.2f", h, m, duration)
	return timeVar
}
