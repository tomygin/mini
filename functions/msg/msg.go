package msg

import (
	"fmt"
	"mini/comment"
	"mini/functions/info"
	"mini/sql"
	"net/http"
	"os"
	"time"

	echo "github.com/labstack/echo"
)

//接受上传的留言信息
func UpMsg(c echo.Context) error {
	uid, _ := comment.TokenId(c)
	if uid == uint(0) {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"用户不存在"`))
	}
	msg := new(sql.Msg)
	c.Bind(msg) //to who and what the msg
	msg.Sip = c.RealIP()
	msg.Sdate = time.Now().Format("20060102150405")

	file, err := c.FormFile("file") //filename要与前端对应上
	if err != nil {
		return err
	}
	// 先打开文件源
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	// 下面创建保存路径文件 file.Filename 即上传文件的名字
	path := "files/" + fmt.Sprintf("%d_", info.FileCount) + file.Filename
	info.FileCount++
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	msg.File = path

	if err := sql.DB.Create(msg).Error; err == nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":[],"msg":"留言成功"`))
	}

	return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"留言失败"`))
}

//用于返回的json
type objmsgs struct {
	code uint
	msgs *[]sql.Msg
	msg  string
}

func GetMsg(c echo.Context) error {
	uid, _ := comment.TokenId(c)
	if uid == uint(0) {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"用户不存在"`))
	}
	var obj objmsgs
	if err := sql.DB.Where("id = ? && isget = ?", uid, false).Find(&obj.msgs).Error; err == nil {
		obj.code = 1
		obj.msg = "获取成功"
		//更新状态
		if err = sql.DB.Model(&sql.Msg{}).Where("id = ? && isget = ?", uid, false).Updates(sql.Msg{
			Isget: true,
			Rip:   c.RealIP(),
			Rdate: time.Now().Format("20060102150405"),
		}).Error; err == nil {
			return c.JSON(http.StatusOK, obj)
		}
		obj.msg = "获取成功，但是数据更新失败"
		return c.JSON(http.StatusOK, obj)

	}
	return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"获取失败"`))
}
