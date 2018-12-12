package gocandy

import (
	"sync"
)

type GoMap struct {
	sync.RWMutex
	Map map[interface{}]interface{}
}

type GoMResult struct {
	value interface{}
}

// new a GoMap struct
func NewMap() *GoMap {
	val := new(GoMap)
	val.Map = make(map[interface{}]interface{})
	return val
}

// set a value for gomap
// 设置一个map值
func (gm *GoMap) Set(keyname interface{}, val interface{}) {
	gm.Lock()
	gm.Map[keyname] = val
	gm.Unlock()
}

// get a value from gomap
// 取得一条value值
func (gm *GoMap) Get(keyname interface{}, def interface{}) *GoMResult {
	gmr := new(GoMResult)
	gm.RLock()
	val, e := gm.Map[keyname]
	gm.RUnlock()
	if !e {
		gmr.value = def
	} else {
		gmr.value = val
	}
	return gmr
}

// delete a value from gomap
// 删除一条value值
func (gm *GoMap) Delete(key interface{}) {
	gm.Lock()
	delete(gm.Map, key)
	gm.Unlock()
}

// get the keys in list
// 获得所有的key 数组
func (gm *GoMap) KeyArray() []GoMResult {
	gmrlist := make([]GoMResult, 0)
	gm.RLock()
	for k := range gm.Map {
		gmrlist = append(gmrlist, GoMResult{value: k})
	}
	gm.RUnlock()
	return gmrlist
}

// clear entire map and set to nil
// 清空整个map
func (gm *GoMap) Clear() {
	gm.Map = make(map[interface{}]interface{})
	gm.Map = nil
}

// format GoMResult to bool
func (gmr *GoMResult) Bool(def bool) bool {
	return MustBool(gmr.value, def)
}

// format GoMResult to string
func (gmr *GoMResult) String(def string) string {
	return MustString(gmr.value, def)
}

// format GoMResult to int8
func (gmr *GoMResult) Int8(def int8) int8 {
	return MustInt8(gmr.value, def)
}

// fotmat GoMResult to int
func (gmr *GoMResult) Int(def int) int {
	return MustInt(gmr.value, def)
}

// format GoMResult to int64
func (gmr *GoMResult) Int64(defautValue int64) int64 {
	return MustInt64(gmr.value, defautValue)
}

// format GoMResult to uint8
func (gmr *GoMResult) Uint8(defautValue uint8) uint8 {
	return MustUint8(gmr.value, defautValue)
}

// format GoMResult to uint32
func (gmr *GoMResult) Uint32(defautValue uint32) uint32 {
	return MustUint32(gmr.value, defautValue)
}

// format GoMResult to uint64
func (gmr *GoMResult) Uint64(defautValue uint64) uint64 {
	return MustUint64(gmr.value, defautValue)
}

// format GoMResult to float32
func (gmr *GoMResult) Float32(defautValue float32) float32 {
	return MustFloat32(gmr.value, defautValue)
}

// format GoMResult to float64
func (gmr *GoMResult) Float64(defautValue float64) float64 {
	return MustFloat64(gmr.value, defautValue)
}

// format GoMResult to []bytes
func (gmr *GoMResult) Bytes(defautValue []byte) []byte {
	return MustBytes(gmr.value, defautValue)
}

// get GoMResult as interface{}
func (gmr *GoMResult) Interface() interface{} {
	return gmr.value
}
