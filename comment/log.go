package comment

import (
	"log"
	"os"
)

//初始化后把日记文件记录控制权放这里
var LogPtr *os.File

var isinit bool

//运行情况记录
//第一个为日志前前缀
//第二个及以后参数为日志信息
func Log(s string, i ...any) {
	if !isinit {
		log.SetOutput(LogPtr)
		log.SetFlags(log.Ldate | log.Ltime)
		isinit = true
	}

	if len(i) == 0 {
		return
	}

	if len(s) != 0 {
		//这里是有前缀的情况
		log.SetPrefix("[" + s + "] ")
	} else {
		log.SetPrefix("")
	}
	log.Println(i...)
}
