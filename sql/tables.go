package sql

//数据库所有的表

type User struct {
	ID       uint   `json:"id" form:"id" query:"id"`
	Name     string `json:"name" form:"name" query:"name"`
	Password string `json:"password" form:"password" query:"password"`
}

type Msg struct {
	Name string `json:"name" form:"name" query:"name"`
	Msg  string `json:"msg" form:"msg" query:"msg"`
}
