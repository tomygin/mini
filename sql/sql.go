package sql

//用于初始化的包
import (
	"fmt"
	"mini/functions/info"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func DB_init() error {
	//获取数据库信息并拼接
	sqlinfo := info.Allinfo.Sql
	dsn := sqlinfo.Name + ":" + sqlinfo.Password + "@tcp(" + sqlinfo.Ip + ":" + sqlinfo.Port + ")/" + sqlinfo.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	// 连接
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		SkipDefaultTransaction: false,

		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	//初始化失败就为err，统一初始化标准而修改
	DB = db
	//如果连接成功
	temp, _ := db.DB()
	temp.SetMaxIdleConns(10)  //空闲
	temp.SetMaxOpenConns(100) //连接个数
	return err
}

//重新连接
func DB_re() {
	if DB == nil || DB.Error != nil {
		if err := DB_init(); err != nil {
			fmt.Println("重新连接数据库失败")
		}
	}
}
