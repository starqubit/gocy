package common

import (
	"log"
	"os/exec"

	"github.com/tidwall/gjson"
)

// 解析音视频文件参数 依赖 ffprobe
func GetFileParam(path string, codec_type string) gjson.Result {
	var j gjson.Result
	cmd := exec.Command("ffprobe", `-v`, "quiet", "-print_format", "json", "-show_format", "-show_streams", path)
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("解析音视频文件参数异常:", err)
	}
	if len(b) > 0 {
		j = gjson.Parse(string(b))
		if codec_type != "" {
			streams := j.Get("streams").Array()
			for _, v := range streams {
				if v.Get("codec_type").String() == codec_type {
					j = v
				}
			}
		}
		return j
	}
	return j
}
