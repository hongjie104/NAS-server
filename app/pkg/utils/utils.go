package utils

import "strings"

// SubCn 将字符串 s 从 sub 子串开始截取 len 个字节，支持中文
// 1. 获取子串位置
// 2. 用 []byte 去掉所有子串前面的的字符
// 3. 将剩余部分转化成 rune ，截取 len 个字节
// 4. 最后转化成 string
func SubCn(s, sub string, len int) (string, bool) {
	subIndex := strings.Index(s, sub)
	if subIndex > 0 {
		sByte := []byte(s)[subIndex:]
		sRune := []rune(string(sByte))[0:len]
		return string(sRune), true
	}
	return "", false
}
