package gocandy

import (
	"log"
	"testing"
)

//  implode test
func Test_Implode(t *testing.T) {
	t.Log("第一个测试通过了")
	var implodearr = []string{"daddy","mum","pad","node"}
	log.Println(Implode(implodearr, ","))

	var implodearr2 = [4]string{"daddy","mum","pad","node"}
	log.Println(Implode(implodearr2, ","))

	var implodearr3 = []int{1,2,3,4}
	log.Println(Implode(implodearr3, ","))

}

func Test_ContainsValue(t *testing.T) {
	var implodearr = []string{"daddy","mum","pad","node"}
	log.Println(ContainsValue(implodearr, "node"))

	var implodearr2 = []string{"123","456","789","000"}
	log.Println(ContainsValue(implodearr2, 456))
}