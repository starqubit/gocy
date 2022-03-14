package R

import (
	"fmt"
	"log"
)

/*
异常收集日志
*/

func Debug(v ...interface{}) {
	text := fmt.Sprint(v...)
	log.Output(2, text)
	r.Output("DEBUG", "default", 1, text)

}

func Info(v ...interface{}) {
	text := fmt.Sprint(v...)
	log.Output(2, text)
	r.Output("INFO", "default", 1, text)

}

func Warning(v ...interface{}) {
	text := fmt.Sprint(v...)
	log.Output(2, text)
	r.Output("WARNING", "default", 1, text)

}

func Error(v ...interface{}) {
	text := fmt.Sprint(v...)
	log.Output(2, text)
	r.Output("ERROR", "default", 1, text)

}

func Fatal(v ...interface{}) {
	text := fmt.Sprint(v...)
	log.Output(2, text)
	r.Output("FATAL", "default", 1, text)

}

/*
// 调用Email将会立即邮件提醒，
// sleep>0 同样的邮件(sleepkey相同)发送过将不会再次发送
// flagId 可能是某个文章id 、句子id、视频id等任何发生异常的数据id方便定位
*/
func FatalSendEmail(sleep int, sleepkey, flagId, text string) {
	log.Output(2, text)
	r.Output("FatalSendEmail", flagId, 1, text, map[string]interface{}{
		"sleep":    sleep,
		"sleepkey": sleepkey,
	})
}

// flagId 长度不超过32
func Notice(flagId string, text string) {
	log.Output(2, text)
	r.Output("NOTICE", flagId, 1, text)
}

/*
// 调用Email将会立即邮件提醒，
// sleep>0 同样的邮件(sleepkey相同)发送过将不会再次发送
// flagId 可能是某个文章id 、句子id、视频id等任何发生异常的数据id方便定位
*/
func Email(sleep int, sleepkey, flagId, text string) {
	log.Output(2, text)
	r.Output("EMAIL", flagId, 1, text, map[string]interface{}{
		"sleep":    sleep,
		"sleepkey": sleepkey,
	})
}
