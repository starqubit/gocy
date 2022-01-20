package R

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/go-resty/resty/v2"
)

/*
异常信息收集客户端
*/

var r *server

type server struct {
	cache chan Message
	url   string
	name  string
}

// 初始化服务
func InitServer(name, url string) {
	if r == nil {
		r = &server{
			name:  name,
			url:   url,
			cache: make(chan Message, 999999),
		}
		go r.start()
	}
}

// 启动服务
func (s *server) start() {
	for {
		m := <-s.cache
		s.post(m)
	}
}

// 上报信息
func (s *server) post(m Message) {
	b, e := json.Marshal(m)
	if e != nil {
		log.Println(e)
		return
	}
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
			// go func() {
			// 	select {
			// 	case s.cache <- m:
			// 	case <-time.After(time.Millisecond * 100):
			// 		return
			// 	}
			// }()
		}
	}()
	client := resty.New()
	resp, err := client.R().SetFormData(map[string]string{
		"data": string(b),
	}).Post(s.url)
	if err != nil {
		log.Println(err)
		return
	}
	if resp.StatusCode() != 200 {
		err = fmt.Errorf(resp.String())
		return
	}
}

// 向缓存添加数据
func (s *server) Output(level, flagId string, calldeep int, text string) {
	if r == nil {
		return
	}
	_, file, line, ok := runtime.Caller(calldeep + 1)
	if !ok {
		file = "???"
		line = 0
	}
	m := Message{
		Timestamp: time.Now().UnixNano(),
		Levelname: level,
		Filename:  file,
		Lineno:    line,
		Text:      text,
		Weights:   1,
		Name:      s.name,
		FlagId:    flagId,
		HostName:  os.Getenv("HOSTNAME"),
	}
	go func() {
		select {
		case s.cache <- m:
		case <-time.After(time.Millisecond * 100):
			return
		}
	}()

}
