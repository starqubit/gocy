package gocandy

import (
	"log"
	"testing"
)

//  implode test
func Test_Implode(t *testing.T) {
	var implodearr = []string{"daddy", "mum", "pad", "node"}
	log.Println(Implode(implodearr, ","))

	var implodearr2 = [4]string{"daddy", "mum", "pad", "node"}
	log.Println(Implode(implodearr2, ","))

	var implodearr3 = []int{1, 2, 3, 4}
	log.Println(Implode(implodearr3, ","))

}

func Test_ContainsValue(t *testing.T) {
	var implodearr = []string{"daddy", "mum", "pad", "node"}
	log.Print("Test_ContainsValue1: ")
	log.Println(ContainsValue(implodearr, "node"))

	var implodearr2 = []string{"123", "456", "789", "000"}
	log.Print("Test_ContainsValue2: ")
	log.Println(ContainsValue(implodearr2, 456))

	var implodearr3 = map[string]interface{}{"hash": "daddy", "map": "mum", "pad": 123, "node": 6.00}
	log.Print("ContainsValue3: ")
	log.Println(ContainsValue(implodearr3, "123"))
}

func Test_ContainsKey(t *testing.T) {
	var implodearr = []string{"daddy", "mum", "pad", "node"}
	log.Print("Test_ContainsKey1: ")
	log.Println(ContainsKey(implodearr, "2"))

	var implodearr2 = map[string]interface{}{"name": "jason", "age": 24, "height": "1.78"}
	log.Print("Test_ContainsKey2: ")
	log.Println(ContainsKey(implodearr2, "age"))
}
