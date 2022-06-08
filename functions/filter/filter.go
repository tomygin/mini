package filter

import (
	"strings"
)

//先执行全局声明再执行init函数
var Words []string

// func init() {
// 	Words, _ = info.FilterWords()
// }

//过滤一些词汇
func Filter(str string) string {
	// 如果配置文件读取失败，words的值为nil,长度为0
	for i := 0; i < len(Words); i++ {
		str = strings.ReplaceAll(str, Words[i], "*")
	}
	return str
}

//暴力查找字符串位置,返回第一个匹配的索引
func BF(s, p string) int {
	i, j := 0, 0
	for i < len(s) && j < len(p) {
		if s[i] == p[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}
	if j == len(p) {
		return i - j
	}
	return -1
}
