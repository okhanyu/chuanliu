package pkg

import (
	"log"
	"strings"
	"time"
)

func PubDateTimeConvert(timeStr string) (time.Time, error) {
	var t time.Time
	var err error
	if timeStr == "" {
		return t, nil
	}

	if strings.Contains(timeStr, "CST") {
		// 将字符串转换为 time.Time 类型
		t, err = time.Parse(time.RFC1123, timeStr)
	} else if strings.Contains(timeStr, "GMT") {
		// 将字符串转换为 time.Time 类型
		t, err = time.Parse(time.RFC1123, timeStr)
	} else if strings.Contains(timeStr, "+0000") {
		// 将字符串转换为 time.Time 类型
		t, err = time.Parse(time.RFC1123Z, timeStr)
	} else if strings.Contains(timeStr, "+0800") {
		// 将字符串转换为 time.Time 类型
		t, err = time.Parse(time.RFC1123Z, timeStr)
	} else {
		// 将字符串转换为 time.Time 类型
		//t, err = time.Parse(time.RFC3339, timeStr)
		timeFormats := []string{
			"2006-01-02T15:04:05+08:00",
			"2006-01-02T15:04:05Z",
			"2006-01-02 15:04:05",
			"2006-01-02T15:04:05.000000+00:00",
			"2006-01-02T15:04:05+00:00",
			//"2026-01-02T15:04:05.000000+00:00",
		}
		// 尝试解析时间字符串
		for _, format := range timeFormats {
			t, err = time.Parse(format, timeStr)
			if err == nil {
				break
			}
		}
	}

	if err != nil {
		log.Printf("解析日期时间字符串 %s 失败：%s\n", timeStr, err)
		return t, err
	}

	return t, nil
}
