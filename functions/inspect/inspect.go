package inspect

import (
	"errors"
	"fmt"
	"mini/comment"
	"mini/functions/filter"
	"mini/functions/info"
	"mini/sql"
	"os"
	"time"
)

const (
	tasks  = 5 //需要初始化的任务数量
	tstime = 4 //初始化超时(秒)
)

func Inspect() {
	start := time.Now()
	c := make(chan error, tasks)
	//计时器
	t := time.NewTimer(tasks * time.Second)
	defer t.Stop()
	//新建一个进度条 当前已经完成0个
	cur := uint(0)
	bar := newbar(cur, tasks)

	//必须初始化的任务，其他任务依赖这个初始化
	go InfoYaml(c)
	select {
	case err := <-c:
		if err == nil {
			cur++
			bar.displaybar(cur)
		} else {
			bar.finsh()
			panic(err)
		}
	case <-t.C:
		//初始化超时，必要配置无法加载，直接关闭程序
		panic("初始化超时")
	}
	//------------------
	//异步执行初始化任务
	go FilePath(c)
	go SqlConnect(c)
	go FilterWordsYaml(c)
	go LogFile(c)
	//------------------
	for cur < tasks {
		select {
		case err := <-c:
			if err == nil {
				//初始化成功
				cur++
				bar.displaybar(cur)
			} else {
				bar.finsh()
				panic(err)
			}
		case <-t.C:
			bar.finsh()
			panic("初始化超时")
		}
	}
	close(c)
	bar.finsh()

	cost := time.Since(start)
	fmt.Println("初始化使用", cost)

}

//初始化必要配置信息
func InfoYaml(c chan<- error) {
	all, err := info.InfoStruct()
	info.Allinfo = all
	c <- err
}

//初始化过滤词汇
func FilterWordsYaml(c chan<- error) {
	words, err := info.FilterWords()
	filter.Words = words
	c <- err
}

//在此之前必须加载了配置文件
func SqlConnect(c chan<- error) {

	if err := sql.DB_init(); err == nil {
		//-------------数据库的表-----------
		tables := []interface{}{
			&sql.Msg{},
			&sql.User{},
		}
		//--------------------------------

		if comment.IsClearTables() {
			for _, table := range tables {
				if err = sql.DB_deltab(table); err != nil {
					c <- err
					return
				}
			}
		}
		for _, table := range tables {
			if err = sql.DB_checktab(table); err != nil {
				c <- err
				return
			}
		}
		c <- err
	}

	c <- errors.New("初始化连接失败")
}

//初始化上传文件的文件夹
func FilePath(c chan<- error) {
	err := comment.CreatePath("./files")
	c <- err
}

//初始化日志记录文件
func LogFile(c chan<- error) {
	file := "./" + "log" + ".txt"

	logPtr, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	comment.LogPtr = logPtr
	c <- err
}

//--------- 华丽的分割线 ----------------
// 以前的旧方案
// //初始化配置检查，并且初始化相关的配置
// func Inspect() bool {
// 	start := time.Now()
// 	//用协程来加载进度条
// 	c := make(chan bool)
// 	//当前已经完成任务，总的任务
// 	go PrinBar(0, 4, c)

// 	//检查yaml的配置文件
// 	if all, err := info.InfoStruct(); err != nil {
// 		fmt.Println(err)
// 		return false
// 	} else {
// 		info.Allinfo = all
// 		c <- true
// 	}

// 	//检查yaml的过滤词汇文件
// 	if words, err := info.FilterWords(); err != nil {
// 		fmt.Println(err)
// 		return false
// 	} else {
// 		filter.Words = words
// 		c <- true
// 	}

// 	//检查数据库的连接
// 	if sql.DB_init() {
// 		// fmt.Println("数据库对接成功")
// 		//检查所有的表的存在
// 		//与表相关的结构放在sql包里面 tables
// 		tables := []interface{}{
// 			&sql.Msg{},
// 		}
// 		for _, table := range tables {
// 			sql.DB_checktab(table)
// 		}
// 		c <- true
// 	} else {
// 		// fmt.Println("数据库初始化失败")
// 		return false
// 	}

// 	//检查上传文件的文件夹files是否存在
// 	if err := comment.CreatePath("./files"); err != nil {
// 		fmt.Println(err)
// 		return false
// 	} else {
// 		c <- true
// 	}
// 	cost := time.Since(start)
// 	fmt.Println("初始化使用", cost, "秒")
// 	return true
// }

// //我们把传统的进度条进一步封装到一个协程里面
// func PrinBar(cur uint, total uint, ch chan bool) {
// 	//实体化一个进度条
// 	bar := NewBar(cur, total)
// 	for cur < total {
// 		<-ch
// 		cur++
// 		bar.DisplayBar(cur)
// 	}
// 	bar.Finsh()
// 	//进度条加载完毕，关闭通道写入，再写入就报错
// 	close(ch)
// 	//协程的返回都会被舍弃掉，这里返回是想关闭协程
// }
