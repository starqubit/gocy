package R

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
定义数据模型
*/

type Message struct {
	Timestamp int64                  `json:"timestamp"` //时间戳，精确到纳秒 时间戳*10^9
	Datetime  string                 `json:"datetime"`  //日期
	Levelname string                 `json:"levelname"` //等级
	Filename  string                 `json:"filename"`  //文件名
	Lineno    int                    `json:"lineno"`    //行号
	Thread    int                    `json:"thread"`    //线程号
	Text      string                 `json:"text"`      //内容
	Type      int                    `json:"type"`      //类型
	Weights   int                    `json:"weights"`   //权重
	FlagId    string                 `json:"flagid"`    //标记id
	Name      string                 `json:"name"`      //程序名
	HostName  string                 `json:"hostname"`  //host name
	RemoteIP  string                 `json:"remoteip"`  //remote ip
	Options   map[string]interface{} `json:"options"`   //额外参数
}

// 格式化为打印的字符串
func (m *Message) format() string {
	datetime := m.formatTime()
	return fmt.Sprintf("%s|%s|%s:%d|%d|%s", datetime, m.Levelname, m.Filename, m.Lineno, m.Thread, m.Text)
}

// 格式化时间
func (m *Message) formatTime() string {
	datetime := time.Unix(m.Timestamp/1000000000, 0).Format("2006-01-02 03:04:05")
	return datetime
}

// 格式化为Json字符串
func (m *Message) formatJson() (string, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
