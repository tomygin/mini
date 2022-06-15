package route

import (
	"mini/functions/user"
	"net/http"

	"github.com/labstack/echo"
)

func Route(e *echo.Echo) {

	usr := e.Group("/user")
	{
		usr.POST("/signup", user.Logup)
		usr.POST("/signin", user.Login)
		usr.GET("/altername", user.AlterName)
		usr.POST("/alterpasswd", user.AlterPassWd)

		usr.GET("/*", LostPage)
		usr.POST("/*", LostPage)
	}

	e.GET("/*", LostPage)
	e.POST("/*", LostPage)

	//用户上传的文件
	e.Static("/files", "../files")

}

func LostPage(c echo.Context) error {
	return c.JSONBlob(http.StatusNotFound, []byte(`{"code":0,"msg":"亲爱的我们暂时不提供此服务"`))
}
