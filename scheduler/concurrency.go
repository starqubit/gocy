package effect

import (
	"fmt"
	"log"
	"sync"
	"time"
)

/*
并发调度
*/

type mytask struct {
	function func(name string, id int) error
	name     string
	id       int
}

type concurrency struct {
	threadChan chan bool
	doneChan   chan int
	tasks      []*mytask
	beginTime  int64
	timeout    int64
	sync.RWMutex
}

// 创建并发任务 threads并发数 timeout超时时间
func NewConcurrency(threads, timeout int64) *concurrency {
	s := new(concurrency)
	s.threadChan = make(chan bool, threads)
	s.doneChan = make(chan int, 999)
	s.tasks = make([]*mytask, 0)
	s.beginTime = time.Now().UnixNano()
	s.timeout = timeout
	return s
}

// 添加任务
func (s *concurrency) AddTask(task func(name string, id int) error, name string, id int) {
	s.Lock()
	defer s.Unlock()
	s.tasks = append(s.tasks, &mytask{
		function: task,
		name:     name,
		id:       id,
	})
}

// 开始任务
func (s *concurrency) Start() {
	s.beginTime = time.Now().UnixNano()
	for _, task := range s.tasks {
		go func(t *mytask) {
			s.threadChan <- true
			defer func() {
				<-s.threadChan
			}()
			err := t.function(t.name, t.id)
			if err != nil {
				log.Println(err)
			}
			s.doneChan <- t.id
		}(task)
	}
}

// 等待所有任务结束
func (s *concurrency) WaitforAllDone() error {
	size := 0
	totalNum := len(s.tasks)
	for {
		select {
		case <-s.doneChan:
			size += 1
			if size >= totalNum {
				return nil
			}
		case <-time.After(100 * time.Millisecond):
			if time.Now().UnixNano()-s.beginTime >= s.timeout {
				return fmt.Errorf("执行任务超时")
			}
		}
	}
}
