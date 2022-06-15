package user

import (
	"mini/comment"
	"mini/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//用户注册
func Logup(c echo.Context) error {
	u := new(sql.User)
	c.Bind(u)
	comment.DeBugPrint("用户注册信息", u)

	if u.ID == 0 || u.Name == "" || u.Password == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"参数错误"`))
	}

	if exist := sql.DB_existid(u.ID); exist {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"用户已经存在"`))
	}

	u.Password = comment.Lock(u.Password)
	//添加一个用户
	if err := sql.DB.Create(u).Error; err == nil {
		return c.JSONBlob(http.StatusOK,
			[]byte(`{"code":1,"data":[{"id":`+strconv.Itoa(int(u.ID))+`}],"msg":"用户注册成功"`))
	}
	return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"用户注册失败"`))
}

//登录
func Login(c echo.Context) error {
	u := new(sql.User)
	c.Bind(u)
	comment.DeBugPrint("登录参数信息", u)
	//放查寻到的数据
	cmp := new(sql.User)
	if err := sql.DB.Where("id = ?", u.ID).First(cmp).Error; err == nil {
		if comment.Lock(u.Password) == cmp.Password {
			//密码正确 设置cookie
			cookie := &http.Cookie{
				Name:  "token",
				Value: comment.JwtMarsh(u.ID, false),
				Path:  "/",
				//cookie有效期为3600秒
				MaxAge: 60 * 60 * 24 * 7,
			}
			comment.DeBugPrint(cookie.Value)
			c.SetCookie(cookie)
			return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":[],"msg":"用户登录成功"`))
		}
	}
	return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"用户登录失败"`))
}

//改名
func AlterName(c echo.Context) error {
	u := new(sql.User)
	name := c.FormValue("name")
	uid, admin := comment.TokenId(c)
	comment.DeBugPrint("从cookie里面获取到uid ", uid, admin)
	if uid != uint(0) {
		if err := sql.DB.Where("id = ?", uid).Find(u).Error; err == nil {
			comment.DeBugPrint("上传上来的名字", name)
			u.Name = name
			//更新
			if err := sql.DB.Model(sql.User{}).Where("id = ?", uid).Updates(u).Error; err == nil {
				return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":[],"msg":"用户名更改成功"`))
			}
			goto fail
		}
		goto fail
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"用户未登录"`))
	}

fail:
	return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"用户名更改失败"`))

}

//改密码
func AlterPassWd(c echo.Context) error {
	u := new(sql.User)
	old := c.FormValue("old")
	new := c.FormValue("new")
	comment.DeBugPrint("上传上来的老密码和新密码", old, new)
	token, err := c.Cookie("token")
	if err != nil {
		return err
	}
	tokenmap := comment.JwtUnMarsh(token.Value)
	uid := tokenmap["uid"]
	comment.DeBugPrint("从cookie里面获取到uid ", uid)
	if uid != uint(0) {

		if err := sql.DB.Where("id = ?", uid).Find(u).Error; err == nil {
			if comment.Lock(old) == u.Password {
				//密码验证成功
				u.Password = comment.Lock(new)
				//更新
				if err := sql.DB.Model(sql.User{}).Where("id = ?", uid).Updates(u).Error; err == nil {
					return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":[],"msg":"密码更改成功"`))
				}
				goto fail
			}

			goto fail
		}
		goto fail
	}

fail:
	return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"密码更改失败"`))
}

//注销用户
func Del(c echo.Context) error {
	u := new(sql.User)
	c.Bind(u)
	if uid, admin := comment.TokenId(c); uid != uint(0) {

		//如果是不是管理员就真正注销
		if !admin {
			cmp := new(sql.User)
			if err := sql.DB.Where("id = ?", u.ID).First(cmp).Error; err == nil {
				if comment.Lock(u.Password) == cmp.Password {
					//密码正确
					if err = sql.DB.Where("id = ?", uid).Delete(sql.User{}).Error; err == nil {
						return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":[],"msg":"注销成功"`))
					}
					return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"密码错误"`))

				}
				return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"用户不存在"`))

			}
		}

	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"管理员目前不能注销"`))
	}
	return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"用户不存在"`))

}
