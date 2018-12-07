package gocandy

import (
	"github.com/issue9/conv"
	"reflect"
)
//  Join array/map/slice elements with a string
// 将array/map/slice用给定的字符串连接
func Implode(listValue interface{}, glue string) (string, error) {
	result, err := Slice(listValue)
	if err != nil {
		return "", err
	}
	return arrayJoin(result, glue),nil
	}

// convert any value to slice
func Slice(input interface{}) ([]interface{}, error) {
	// is an array?
	v := reflect.ValueOf(input)
	switch v.Kind() {
	case reflect.Array:
		s := reflect.ValueOf(input)
		val := make([]interface{}, v.Len())
		for i := 0; i < s.Len(); i++ {
			val[i] = s.Index(i)
		}
		return val,nil
	default:
		val, err := conv.Slice(input)
		if err != nil {
			return nil,err
		}
		return val, nil
	}
}
// only for Implode implementation
func arrayJoin(input []interface{}, glue string) string {
	var result string
	for _, v := range input {
		value, _ := conv.String(v)
		if result == "" {
			result = value
		} else {
			result = result + glue + value
		}
	}
	return result
}
// check if a value exists in array/slice/map does not support sub array
// 在array/slice/map中寻找值 不支持子数组
func ContainsValue(input interface{}, subitem interface{}) (bool, error) {
	result, err := Slice(input)
	if err != nil {
		return false, err
	}
	needle, err := conv.String(subitem)
	if err != nil {
		return false, err
	}
	for _, v := range result {
		s, _ := conv.String(v)
		if s == needle {
			return true, nil
		}
	}
	return false,nil
}