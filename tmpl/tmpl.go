//注册echo的模板
package tmpl

import (
	"io"
	"text/template"

	"github.com/labstack/echo"
)

type Renderer interface {
	//渲染函数定义
	//第一参数用于保存渲染模版后的结果
	//第二个参数是模版名字
	//第三个参数是传入模版的参数，可以是任意类型
	//第四个参数是echo.Context
	Render(io.Writer, string, interface{}, echo.Context) error
}

//自定义的模版引擎struct
type Template struct {
	Templates *template.Template
}

//实现接口，Render函数
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	//调用模版引擎渲染模版
	return t.Templates.ExecuteTemplate(w, name, data)
}

//初始化模板引擎，加载views目录下所有的模板
func TmplInit() *Template {
	return &Template{Templates: template.Must(template.ParseGlob("views/*.html"))}
}
