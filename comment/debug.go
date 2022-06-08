package comment

import (
	"flag"
	"fmt"
)

//目前只能让一个参数生效
type flags struct {
	dp bool //debug switch
	ct bool //clear the sql tables whenn init
}

func getflags() flags {
	f := new(flags)
	flag.BoolVar(&f.dp, "dp", false, "debug print more information")
	flag.BoolVar(&f.ct, "ct", false, "clear the sql tables before use")
	flag.Parse()
	return *f
}

var debug = getflags()

//只暴露api函数

func DeBugPrint(i ...any) {
	if debug.dp {
		fmt.Println()
		fmt.Println(i...)
		fmt.Println()
	}
}

//导出到inspect
func IsClearTables() bool {
	return debug.ct
}
