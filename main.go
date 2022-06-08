package main

import (
	"mini/functions/info"
	"mini/functions/inspect"
	"mini/route"

	echo "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//初始化所有配置文件
func init() {
	inspect.Inspect()
}
func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//删除尾部的 /
	e.Pre(middleware.RemoveTrailingSlash())

	e.Debug = true
	e.HideBanner = true

	route.Route(e)

	port := info.Allinfo.Host.Port
	e.Logger.Fatal(e.Start(":" + port))
}
