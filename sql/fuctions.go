package sql

//查询id是否存在于用户数据库中
func DB_existid(id uint) bool {
	err := DB.Where("id = ?", id).First(&User{}).Error
	// fmt.Println(err, "检查id")
	return err == nil
}

//创建表，如果不存在
func DB_checktab(table interface{}) error {
	created := DB.Migrator().HasTable(table)
	if !created {
		err := DB.AutoMigrate(table)
		return err
	}
	return nil
}

//删除一个表
func DB_deltab(table interface{}) error {
	err := DB.Migrator().DropTable(table)
	return err
}
