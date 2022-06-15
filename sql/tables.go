package sql

//数据库所有的表

type User struct {
	ID       uint   `json:"id" form:"id" query:"id"`
	Name     string `json:"name" form:"name" query:"name"`
	Password string `json:"password" form:"password" query:"password"`
	Admin    bool
}

type Msg struct {
	M      string `json:"msg" form:"msg" query:"msg"`          //信息
	Accept uint   `json:"accept" form:"accept" query:"accept"` //给谁
	Isget  bool   //接收状态	true已经表示接受到
	Sip    string //获取到发送者的Ip
	Rip    string // 获取到接收者的Ip
	Sdate  string //发送时间
	Rdate  string //接受时间
	File   string //上传上来的文件的地址
}
