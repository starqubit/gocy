package common

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"time"
)

/*
日期时间操作
*/

// 字符串转时间 处理一些不规则字符串
func Str2Time(strDate string) time.Time {
	if t, err := BeforeHouse2Time(strDate); err == nil {
		return t
	} else if t, err := BeforeDay2Time(strDate); err == nil {
		return t
	} else if t, err := Yesterday2Time(strDate); err == nil {
		return t
	} else if t, err := DateYear2Time(strDate); err == nil {
		return t
	} else if t, err := DateMonth2Time(strDate); err == nil {
		return t
	}
	log.Println("不支持的时间格式", strDate)
	return time.Date(1970, 1, 1, 0, 0, 0, 0, time.Now().Location())
}

// 日期转时间 6小时前
func BeforeHouse2Time(strDate string) (time.Time, error) {
	tTime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.Now().Location())
	re, err := regexp.Compile(`^(\d{1,2})小时前$`)
	if err != nil {
		return tTime, err
	}
	results := re.FindAllStringSubmatch(strDate, -1)
	if len(results) == 0 {
		return tTime, errors.New("匹配失败")
	}

	strOffset := results[0][1]
	nOffset, err := strconv.Atoi(strOffset)
	if err != nil {
		return tTime, err
	}
	dstTime := time.Now().Add(-time.Hour * time.Duration(nOffset))
	return dstTime, nil
}

// 日期转时间 2天前
func BeforeDay2Time(strDate string) (time.Time, error) {
	tTime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.Now().Location())
	re, err := regexp.Compile(`^(\d{1,4})天前$`)
	if err != nil {
		return tTime, err
	}
	results := re.FindAllStringSubmatch(strDate, -1)
	if len(results) == 0 {
		return tTime, errors.New("匹配失败")
	}

	strOffset := results[0][1]
	nOffset, err := strconv.Atoi(strOffset)
	if err != nil {
		return tTime, err
	}
	dstTime := time.Now().AddDate(0, 0, -nOffset)
	return dstTime, nil
}

// 日期转时间 昨天
func Yesterday2Time(strDate string) (time.Time, error) {
	if strDate == "昨天" {
		return time.Time{}, errors.New("匹配失败")
	}
	dstTime := time.Now().AddDate(0, 0, -1)
	return dstTime, nil
}

// 日期转时间 2021年12月31日
func DateYear2Time(strDate string) (time.Time, error) {
	tTime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.Now().Location())
	re, err := regexp.Compile(`^(\d{4})年(\d{2})月(\d{2})日$`)
	if err != nil {
		return tTime, err
	}
	results := re.FindAllStringSubmatch(strDate, -1)
	if len(results) == 0 {
		return tTime, errors.New("匹配失败")
	}

	strYear := results[0][1]
	nYear, err := strconv.Atoi(strYear)
	if err != nil {
		return tTime, err
	}
	strMonth := results[0][2]
	nMonth, err := strconv.Atoi(strMonth)
	if err != nil {
		return tTime, err
	}
	strDay := results[0][2]
	nDay, err := strconv.Atoi(strDay)
	if err != nil {
		return tTime, err
	}
	now := time.Now()
	dstTime := time.Date(nYear, time.Month(nMonth), nDay, now.Hour(), now.Minute(), now.Second(), 0, now.Location())
	return dstTime, nil
}

// 日期转时间 12月31日
func DateMonth2Time(strDate string) (time.Time, error) {
	tTime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.Now().Location())
	re, err := regexp.Compile(`^(\d{2})月(\d{2})日$`)
	if err != nil {
		return tTime, err
	}
	results := re.FindAllStringSubmatch(strDate, -1)
	if len(results) == 0 {
		return tTime, errors.New("匹配失败")
	}
	log.Println(results)
	strMonth := results[0][2]
	nMonth, err := strconv.Atoi(strMonth)
	if err != nil {
		return tTime, err
	}
	strDay := results[0][2]
	nDay, err := strconv.Atoi(strDay)
	if err != nil {
		return tTime, err
	}
	now := time.Now()
	dstTime := time.Date(now.Year(), time.Month(nMonth), nDay, now.Hour(), now.Minute(), now.Second(), 0, now.Location())
	return dstTime, nil
}
