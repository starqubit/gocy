package gocandy

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func typeError(val interface{}, t string) error {
	return fmt.Errorf("[%T:%v]无法转换成[%v]类型", val, val, t)
}

// convert any value to bool
// Bool 将 val 转换成 bool 类型或是在无法转换的情况下返回 error。
//  123(true), 0(false),"-123"(true), "on"(true), "off"(false), "true"(true), "false"(false)
func Bool(val interface{}) (bool, error) {
	switch ret := val.(type) {
	case bool:
		return ret, nil
	case int:
		return ret != 0, nil
	case int8:
		return ret != 0, nil
	case int32:
		return ret != 0, nil
	case int64:
		return ret != 0, nil
	case float32:
		return ret != 0, nil
	case float64:
		return ret != 0, nil
	case uint:
		return ret != 0, nil
	case uint8:
		return ret != 0, nil
	case uint32:
		return ret != 0, nil
	case uint64:
		return ret != 0, nil
	case []byte:
		return str2Bool(string(ret))
	case string:
		return str2Bool(ret)
	default:
		return false, typeError(val, "bool")
	}
}

// 字符串转 bool 值，供 Bool() 函数调用。
// 添加了一些 strconv.ParseFloat 不支持但又比较常用的字符串转换
func str2Bool(str string) (bool, error) {
	if val, err := strconv.ParseBool(str); err == nil {
		return val, nil
	} else if val, err := strconv.ParseFloat(str, 32); err == nil {
		return val != 0, nil
	}

	switch strings.ToLower(strings.TrimSpace(str)) {
	case "on":
		return true, nil
	case "off":
		return false, nil
	default:
		return false, typeError(str, "bool")
	}
}

// convert any value to uint64
// Uint64 将 val 转换成 uint64 类型或是在无法转换的情况下返回 error。
// 将一个有符号整数转换成无符号整数，负数将返回错误，正数和零正常转换
func Uint64(val interface{}) (uint64, error) {
	switch ret := val.(type) {
	case uint64:
		return ret, nil
	case int:
		if ret < 0 {
			return 0, typeError(ret, "uint64")
		}
		return uint64(ret), nil
	case int8:
		if ret < 0 {
			return 0, typeError(ret, "uint64")
		}
		return uint64(ret), nil
	case int32:
		if ret < 0 {
			return 0, typeError(ret, "uint64")
		}
		return uint64(ret), nil
	case int64:
		if ret < 0 {
			return 0, typeError(ret, "uint64")
		}
		return uint64(ret), nil
	case uint:
		return uint64(ret), nil
	case uint8:
		return uint64(ret), nil
	case uint32:
		return uint64(ret), nil
	case float32:
		if ret < 0 {
			return 0, typeError(ret, "uint64")
		}
		return uint64(ret), nil
	case float64:
		if ret < 0 {
			return 0, typeError(ret, "uint64")
		}
		return uint64(ret), nil
	case bool:
		if ret {
			return 1, nil
		}
		return 0, nil
	case []byte:
		if val, err := strconv.ParseFloat(string(ret), 32); err == nil {
			return uint64(val), nil
		}
		return 0, typeError(val, "uint64")
	case string:
		if val, err := strconv.ParseFloat(ret, 32); err == nil {
			return uint64(val), nil
		}
		return 0, typeError(val, "uint64")
	default:
		return 0, typeError(ret, "uint64")
	}
}

// convert any value to uint
// Uint 将 val 转换成 uint 类型或是在无法转换的情况下返回 error。
func Uint(val interface{}) (uint, error) {
	ret, err := Uint64(val)
	if err != nil {
		return 0, err
	}
	return uint(ret), nil
}

// convert any value to uint8
// Uint8 将 val 转换成 uint8 类型或是在无法转换的情况下返回 error。
func Uint8(val interface{}) (uint8, error) {
	ret, err := Uint64(val)
	if err != nil {
		return 0, err
	}
	return uint8(ret), nil
}

// convert any value to uint32
// Uint32 将 val 转换成 uint32 类型或是在无法转换的情况下返回 error。
func Uint32(val interface{}) (uint32, error) {
	ret, err := Uint64(val)
	if err != nil {
		return 0, err
	}
	return uint32(ret), nil
}

// convert any value to int64
// Int64 将 val 转换成 int64 类型或是在无法转换的情况下返回 error。
func Int64(val interface{}) (int64, error) {
	switch ret := val.(type) {
	case int64:
		return ret, nil
	case int:
		return int64(ret), nil
	case int8:
		return int64(ret), nil
	case int32:
		return int64(ret), nil
	case uint:
		return int64(ret), nil
	case uint8:
		return int64(ret), nil
	case uint32:
		return int64(ret), nil
	case uint64:
		return int64(ret), nil
	case float32:
		return int64(ret), nil
	case float64:
		return int64(ret), nil
	case bool:
		if ret {
			return 1, nil
		}
		return 0, nil
	case []byte:
		val, err := strconv.ParseFloat(string(ret), 32)
		if err == nil {
			return int64(val), nil
		}
		return -1, typeError(val, "int64")
	case string:
		val, err := strconv.ParseFloat(ret, 0)
		if err == nil {
			return int64(val), nil
		}

		return -1, typeError(val, "int64")
	default:
		return -1, typeError(ret, "int64")
	}
}

func MustInt64(val interface{}, def int64) int64 {
	r, err := Int64(val)
	if err != nil {
		return def
	}
	return r
}

// convert any value to int
// Int 将 val 转换成 int 类型或是在无法转换的情况下返回 error。
func Int(val interface{}) (int, error) {
	ret, err := Int64(val)
	if err != nil {
		return -1, err
	}
	return int(ret), err
}

func MustInt(val interface{}, def int) int {
	r, err := Int64(val)
	if err != nil {
		return def
	}
	return int(r)
}

// convert any value to int8
// Int8 将 val 转换成 int8 类型或是在无法转换的情况下返回 error。
func Int8(val interface{}) (int8, error) {
	ret, err := Int64(val)
	if err != nil {
		return -1, err
	}
	return int8(ret), err
}

// convert any value to int32
// Int32 将 val 转换成 int32 类型或是在无法转换的情况下返回 error。
func Int32(val interface{}) (int32, error) {
	ret, err := Int64(val)
	if err != nil {
		return -1, err
	}
	return int32(ret), err
}

// convert any value to float64
// Float64 将 val 转换成 float64 类型或是在无法转换的情况下返回 error。
func Float64(val interface{}) (float64, error) {
	switch ret := val.(type) {
	case float64:
		return ret, nil
	case int:
		return float64(ret), nil
	case int8:
		return float64(ret), nil
	case int32:
		return float64(ret), nil
	case int64:
		return float64(ret), nil
	case uint:
		return float64(ret), nil
	case uint8:
		return float64(ret), nil
	case uint32:
		return float64(ret), nil
	case uint64:
		return float64(ret), nil
	case float32:
		return float64(ret), nil
	case bool:
		if ret {
			return 1.0, nil
		}
		return 0.0, nil
	case []byte:
		val, err := strconv.ParseFloat(string(ret), 64)
		if err == nil {
			return float64(val), nil
		}
		return -1, typeError(val, "float64")
	case string:
		val, err := strconv.ParseFloat(ret, 64)
		if err == nil {
			return float64(val), nil
		}
		return -1, typeError(val, "float64")
	default:
		return -1, typeError(ret, "float64")
	}
}

// convert any value to float32
// Float32 将 val 转换成 float32 类型或是在无法转换的情况下返回 error。
func Float32(val interface{}) (float32, error) {
	ret, err := Float64(val)
	if err != nil {
		return -1.0, err
	}
	return float32(ret), nil
}

// convert any value to a string
// String 将 val 转换成 string 类型或是在无法转换的情况下返回 error。
func String(val interface{}) (string, error) {
	switch ret := val.(type) {
	case string:
		return ret, nil
	case []byte:
		return string(ret), nil
	case int64:
		return strconv.FormatInt(ret, 10), nil
	case int:
		return strconv.FormatInt(int64(ret), 10), nil
	case int8:
		return strconv.FormatInt(int64(ret), 10), nil
	case int32:
		return strconv.FormatInt(int64(ret), 10), nil
	case uint:
		return strconv.FormatInt(int64(ret), 10), nil
	case uint8:
		return strconv.FormatInt(int64(ret), 10), nil
	case uint32:
		return strconv.FormatInt(int64(ret), 10), nil
	case uint64:
		return strconv.FormatInt(int64(ret), 10), nil
	case float32:
		return strconv.FormatFloat(float64(ret), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(ret, 'f', -1, 64), nil
	case bool:
		return strconv.FormatBool(ret), nil
	case fmt.Stringer:
		return ret.String(), nil
	case error:
		return ret.Error(), nil
	default:
		return "", typeError(ret, "string")
	}
}

// convert any value to string or return del value in case of error
func MustString(val interface{}, def string) string {
	r, err := String(val)
	if err != nil {
		return def
	}
	return r
}

// convert any value to []byte
// Bytes 将 val 转换成 []byte 类型或是在无法转换的情况下返回 error。
func Bytes(val interface{}) ([]byte, error) {
	switch ret := val.(type) {
	case []byte:
		return ret, nil
	case string:
		return []byte(ret), nil
	case int64:
		return []byte(strconv.FormatInt(ret, 10)), nil
	case int:
		return []byte(strconv.FormatInt(int64(ret), 10)), nil
	case int8:
		return []byte(strconv.FormatInt(int64(ret), 10)), nil
	case int32:
		return []byte(strconv.FormatInt(int64(ret), 10)), nil
	case uint:
		return []byte(strconv.FormatInt(int64(ret), 10)), nil
	case uint8:
		return []byte(strconv.FormatInt(int64(ret), 10)), nil
	case uint32:
		return []byte(strconv.FormatInt(int64(ret), 10)), nil
	case uint64:
		return []byte(strconv.FormatInt(int64(ret), 10)), nil
	case float32:
		return []byte(strconv.FormatFloat(float64(ret), 'f', 5, 32)), nil
	case float64:
		return []byte(strconv.FormatFloat(ret, 'f', 5, 64)), nil
	case bool:
		return []byte(strconv.FormatBool(ret)), nil
	default:
		return nil, typeError(ret, "[]byte")
	}
}

func Slice(val interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Array:
		s := reflect.ValueOf(val)
		result := make([]interface{}, v.Len())
		for i := 0; i < s.Len(); i++ {
			result[i] = s.Index(i)
		}
		return result, nil
	case reflect.Map:
		result := make([]interface{}, v.Len())
		s := reflect.ValueOf(val)
		mapkeys := s.MapKeys()
		var key int
		for _, k := range mapkeys {
			result[key] = v.MapIndex(k).Interface()
			key++
		}
		return result, nil
	default:
		switch data := val.(type) {
		case []interface{}:
			return data, nil
		case []int:
			ret := make([]interface{}, len(data))
			for k, v := range data {
				ret[k] = v
			}
			return ret, nil
		case []int8:
			ret := make([]interface{}, len(data))
			for k, v := range data {
				ret[k] = v
			}
			return ret, nil
		case []int32:
			ret := make([]interface{}, len(data))
			for k, v := range data {
				ret[k] = v
			}
			return ret, nil
		case []int64:
			ret := make([]interface{}, len(data))
			for k, v := range data {
				ret[k] = v
			}
			return ret, nil
		case []uint:
			ret := make([]interface{}, len(data))
			for k, v := range data {
				ret[k] = v
			}
			return ret, nil
		case []uint8:
			ret := make([]interface{}, len(data))
			for k, v := range data {
				ret[k] = v
			}
			return ret, nil
		case []uint32:
			ret := make([]interface{}, len(data))
			for k, v := range data {
				ret[k] = v
			}
			return ret, nil
		case []uint64:
			ret := make([]interface{}, len(data))
			for k, v := range data {
				ret[k] = v
			}
			return ret, nil
		case []float32:
			ret := make([]interface{}, len(data))
			for k, v := range data {
				ret[k] = v
			}
			return ret, nil
		case []string:
			ret := make([]interface{}, len(data))
			for k, v := range data {
				ret[k] = v
			}
			return ret, nil
		case string:
			ret := make([]interface{}, len(data))
			for k, v := range data {
				ret[k] = v
			}
			return ret, nil
		default:
			return nil, typeError(data, "slice")
		}
	}
	return nil, nil
}
