package msg

import (
	"fmt"
	"mini/sql"
	"net/http"

	echo "github.com/labstack/echo"
)

//当前页面的数据
type Page struct {
	Msgs []sql.Msg
}

// msg := sql.Msg{
// 		Name: "go",
// 		Msg: "welcome to go the world",
// 	}
// 	if err := sql.DB.Create(&msg).Error; err != nil {
// 		fmt.Println("添加数据失败")
// 	}

func Msg(c echo.Context) error {

	//提取上传上来的有效数据
	updata := sql.Msg{}
	c.Bind(&updata)

	fmt.Println()
	fmt.Println(updata)
	fmt.Println("获取到ip", c.RealIP())
	fmt.Println()

	if updata.Msg != "" && updata.Name != "" {
		//过滤单词
		// need fixing
		// updata.Msg = filter.Filter(updata.Msg)
		// updata.Name = filter.Filter(updata.Name)

		//有效的提交 放进数据库
		if err := sql.DB.Create(&updata).Error; err != nil {
			fmt.Println("添加失败")
		}
	}

	//从数据库取出所有的msg数据
	var page Page
	var msgs []sql.Msg
	if err := sql.DB.Find(&msgs).Error; err != nil {
		return c.String(http.StatusOK, "数据库加载错误")
	}
	page.Msgs = msgs
	return c.Render(http.StatusOK, "msg", page)
}
