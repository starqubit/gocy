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

// flagId 长度不超过32
func Notice(flagId string, text string) {
	log.Output(2, text)
	r.Output("NOTICE", flagId, 1, text)
}

// flagId 长度不超过32
func Email(flagId string, text string) {
	log.Output(2, text)
	r.Output("EMAIL", flagId, 1, text)
}
