package R

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

func TestR(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	InitServer("R_test", "http://127.0.0.1:29721/report?token=123sdffd6t536")
	for i := 0; i < 5; i++ {
		Debug("测试测试", i)
		time.Sleep(time.Second * 2)
	}
	time.Sleep(time.Second * 10)
}

func TestROutput(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	InitServer("R_test", "")
	for i := 0; i < 5; i++ {
		Output("DEBUG", "dddd", 1, "hhhhhahahaha")
		time.Sleep(time.Second * 1)
	}
}

func TestReportCustom(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	InitServer("R_test", "http://127.0.0.1:29721/report?token=123sdffd6t536")
	for i := 0; i < 2; i++ {
		ReportCustom("DEBUG222", "default", 1, "ddddd")
		time.Sleep(time.Second * 2)
	}
	time.Sleep(time.Second * 3)
}

func TestMessage(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	m := Message{
		Name: "test",
		Options: map[string]interface{}{
			"dddd":  1,
			"dfdfd": "test",
		},
	}
	b, e := json.Marshal(m)
	if e != nil {
		log.Println(e)
		return
	}
	t.Log(string(b))
}
