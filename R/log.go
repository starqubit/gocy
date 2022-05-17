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
// flagId 可能是某个文章id 、句子id、视频id等任何发生异常的数据id方便定位，发送邮件时候为标题后缀
// options 第1个参数为自定义接收邮件通知的地址，格式为逗号分割的邮件地址，如：1@qq.com,2@qq.com
*/
func FatalSendEmail(sleep int, sleepkey, flagId, text string, options ...string) {
	log.Output(2, text)

	extraMap := map[string]interface{}{
		"sleep":    sleep,
		"sleepkey": sleepkey,
	}
	if len(options) > 0 {
		extraMap["noticeEmail"] = options[0]
	}
	r.Output("FatalSendEmail", flagId, 1, text, extraMap)
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

// 自定义报告，服务端自动忽略level白名单外的接口反馈，除非在服务端存在匹配的自定义处理器
func ReportCustom(level string, flagId string, calldeep int, v ...interface{}) {
	if len(v) < 1 {
		return
	}
	text := fmt.Sprint(v...)
	log.Output(calldeep+1, text)
	r.Output(level, flagId, calldeep, text)
}

// 自定义Output 直接调用 calldeep=1
func Output(level string, flagId string, calldeep int, isPrint bool, text string, options ...map[string]interface{}) {
	if isPrint {
		log.Output(calldeep+1, fmt.Sprintf("%s %s", level, text))
	}
	r.Output(level, flagId, calldeep, text, options...)
}
