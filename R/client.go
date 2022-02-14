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
	cache   chan Message
	retry   chan Message //异常需要重试的队列
	url     string
	name    string
	errTime time.Time //异常时间
}

// 初始化服务
func InitServer(name, url string) {
	if r == nil {
		r = &server{
			name:    name,
			url:     url,
			cache:   make(chan Message, 999999),
			retry:   make(chan Message, 999999),
			errTime: time.Now(),
		}
		go r.start()
	}
}

// 启动服务
func (s *server) start() {
	go func() {
		// 发送异常的数据重新加入发送队列
		for {
			m := <-s.retry
			if time.Now().Unix()-s.errTime.Unix() <= 3 {
				// 避免发生异常不断请求造成过大负载
				time.Sleep(time.Second * 3)
			}
			s.cache <- m
			if len(s.retry) > 10000 {
				log.Println("异常日志队列数据超过10000，请检查接口是否异常")
			}
		}
	}()
	// 发送消息
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
			s.errTime = time.Now()
			go func() {
				select {
				case s.retry <- m:
				case <-time.After(time.Millisecond * 100):
					return
				}
			}()
		}
	}()
	client := resty.New()
	resp, err := client.R().SetFormData(map[string]string{
		"data": string(b),
	}).Post(s.url)
	if err != nil {
		return
	}
	if resp.StatusCode() != 200 {
		err = fmt.Errorf(resp.String())
		return
	}
}

// 向缓存添加数据
func (s *server) Output(level, flagId string, calldeep int, text string, options ...map[string]interface{}) {
	if r == nil {
		return
	}
	_, file, line, ok := runtime.Caller(calldeep + 1)
	if !ok {
		file = "???"
		line = 0
	}
	hostname, err := os.Hostname()
	if err != nil || hostname == "" {
		hostname = os.Getenv("HOSTNAME")
	}
	extra := make(map[string]interface{})
	if len(options) > 0 {
		extra = options[0]
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
		HostName:  hostname,
		Options:   extra,
	}
	go func() {
		select {
		case s.cache <- m:
		case <-time.After(time.Millisecond * 100):
			return
		}
	}()

}
