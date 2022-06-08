package inspect

import (
	"fmt"
)

//这个方法是为导出的方法服务的，只用于导出函数，小写也并不影响要导出方法的导出性
//这里的都是为 Insepect 函数服务的
//并且都在同一个包中，所以结构体也不用导出，故全小写了

//为检查的时候添加一个进度条

type bar struct {
	cur   uint
	total uint
	per   uint //进度百分比，这里我们用百分数前面的整数

	graph string
	rate  string
}

//展示进度条
func (b *bar) displaybar(cur uint) {
	b.cur = cur
	b.per = b.getPer()
	b.rate = ""
	for i := uint(0); i < b.per; i += 2 {
		b.rate += b.graph
	}
	fmt.Printf("\r%-50s %d%%", b.rate, b.per)
}

//初始化一个进度条
func newbar(cur, total uint) bar {
	//防止cur 和 total 配置不当
	if cur > total {
		fmt.Println("tips:cur和total大小关系错误")
		cur, total = total, cur
	}

	var b bar
	b.cur = cur
	b.total = total
	b.graph = "#"
	return b
}

func (b *bar) getPer() uint {
	return uint(float32(b.cur) / float32(b.total) * 100)
}

func (b *bar) finsh() {
	fmt.Println("")
}
