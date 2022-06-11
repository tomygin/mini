package route

import (
	"mini/functions/user"
	"net/http"

	"github.com/labstack/echo"
)

func Route(e *echo.Echo) {
	e.POST("/signup", user.Add)
	e.POST("/signin", user.Login)
	e.GET("/altername", user.AlterName)
	e.POST("/alterpasswd", user.AlterPassWd)

	e.GET("/*", LostPage)
}

func LostPage(c echo.Context) error {
	return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"未发现此服务"`))
}
