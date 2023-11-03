package pkg

import (
	"fmt"
	"unicode/utf8"
)

func HandleUTF8(str string) string {
	if !utf8.ValidString(str) {
		// 字符串包含非UTF-8字符
		// 将非UTF-8字符替换为utf8.RuneError
		validStr := make([]rune, 0, len(str))
		for i, r := range str {
			if r == utf8.RuneError {
				// 非UTF-8字符
				fmt.Printf("Invalid UTF-8 character at index %d\n", i)
				validStr = append(validStr, '�') // 使用utf8.RuneError替换
			} else {
				validStr = append(validStr, r)
			}
		}
		str = string(validStr)
	}
	return str
}
