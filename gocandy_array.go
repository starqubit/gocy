package gocandy

import (
	"reflect"
)

//  Join array/map/slice elements with a string
// 将array/map/slice用给定的字符串连接
func Implode(listValue interface{}, glue string) (string, error) {
	result, err := Slice(listValue)
	if err != nil {
		return "", err
	}
	return arrayJoin(result, glue), nil
}

// only for Implode implementation
func arrayJoin(input []interface{}, glue string) string {
	var result string
	for _, v := range input {
		value, _ := String(v)
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
func ContainsValue(val interface{}, subitem interface{}) (bool, error) {
	result, err := Slice(val)
	if err != nil {
		return false, err
	}
	needle, err := String(subitem)
	if err != nil {
		return false, err
	}
	for _, v := range result {
		s, _ := String(v)
		if s == needle {
			return true, nil
		}
	}
	return false, nil
}

// check if a value exists in array/slice/map, return def value in case error
func MustContainsValue(val interface{}, subitem interface{}, def bool) bool {
	r, err := ContainsValue(val, subitem)
	if err != nil {
		return def
	}
	return r
}

// check if a key exists in array/slice/map does not support sub array
// 在array/slice/map中寻找键 不支持子数组
func ContainsKey(val interface{}, subitem interface{}) (bool, error) {
	needle, err := String(subitem)
	if err != nil {
		return false, err
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Map:
		s := reflect.ValueOf(val)
		mapkeys := s.MapKeys()
		for _, k := range mapkeys {
			kstring, _ := String(k)
			if kstring == needle {
				return true, nil
			}
		}
	default:
		result, err := Slice(val)
		if err != nil {
			return false, err
		}
		for k := range result {
			s, _ := String(k)
			if s == needle {
				return true, nil
			}
		}
	}
	return false, nil
}

// check if a key exists in array/slice/map, return def value in case error
func MustContainsKey(val interface{}, subitem interface{}, def bool) bool {
	r, err := ContainsKey(val, subitem)
	if err != nil {
		return def
	}
	return r
}
