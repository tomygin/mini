package comment

import (
	"math/rand"
	"time"
)

//生成随机字符串
func RandomStr(n int) string {
	var randomStr = [...]string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
		"k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
		"u", "v", "w", "x", "y", "z", "A", "B", "C", "D",
		"E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X",
		"Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7",
		"8", "9"}
	data := ""
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		data += randomStr[rand.Intn(len(randomStr))]
	}
	return data
}

// 加密由a-z组成的字符串
func Lock(s string) string {
	shifs := make([]int, 26)
	b := []byte(s)
	//获取单词的频率
	for i := range b {
		shifs[b[i]-'a']++
	}
	//力扣848题的加密方法
	//这是我的题解 日期 20220428
	for i, sum := len(b)-1, 0; i >= 0; i-- {
		shifs = shifs[:i+1]
		sum += shifs[i]
		sum %= 26
		b[i] = 'a' + (b[i]-'a'+byte(sum))%26
	}
	return string(b)
}
