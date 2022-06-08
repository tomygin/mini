package route

import (
	"mini/functions/user"

	"github.com/labstack/echo"
)

func Route(e *echo.Echo) {
	e.POST("/signup", user.Add)
	e.POST("/signin", user.Login)
	e.GET("/altername", user.AlterName)
	e.POST("/alterpasswd", user.AlterPassWd)
}
