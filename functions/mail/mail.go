package mail

import (
	"mini/functions/info"

	gomail "gopkg.in/gomail.v2"
)

//发邮件模块
func Mail(address, sub, context string) bool {
	mailinfo := info.Allinfo.Email
	//开始调用gomail模块
	m := gomail.NewMessage()
	//填入邮件信息
	m.SetHeader("From", mailinfo.Account)
	m.SetHeader("To", address)
	m.SetHeader("Subject", sub)
	m.SetBody("text/html", context)

	//准备发送
	d := gomail.NewDialer("smtp.qq.com", 465, mailinfo.Account, mailinfo.Kay)
	//开始发送
	err := d.DialAndSend()
	return err == nil
}
