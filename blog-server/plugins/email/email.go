package email

import (
	"blog-server/global"
	"fmt"
	"gopkg.in/gomail.v2"
	"time"
)

type Subject string

const (
	Code  Subject = "平台验证码"
	Note  Subject = "操作通知"
	Alarm Subject = "告警通知"
)

type Api struct {
	Subject Subject
}

func (a Api) Send(name, body string) error {
	return send(name, string(a.Subject), body)
}

func NewCode() Api {
	return Api{
		Subject: Code,
	}
}
func NewNote() Api {
	return Api{
		Subject: Note,
	}
}
func NewAlarm() Api {
	return Api{
		Subject: Alarm,
	}
}

// send 邮件发送  发给谁，主题，正文
func send(name, subject, body string) error {
	e := global.GlobalC.Email
	return sendMail(
		e.User,             // 发送人
		e.Password,         // 发送人的认证码
		e.Host,             // 对应邮箱厂商的host
		e.Port,             // 对应邮箱厂商的port
		name,               // 收件人
		e.DefaultFromEmail, // 默认的发件人名字
		subject,            // 发送标题
		body,               // 发送主题
	)
}

func sendMail(userName, authCode, host string, port int, mailTo, sendName string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(userName, sendName)) // 谁发的,发送标题
	m.SetHeader("To", mailTo)                                // 发送给谁
	m.SetHeader("Subject", subject)                          // 主题
	m.SetBody("text/html", genHtml(body))
	d := gomail.NewDialer(host, port, userName, authCode)
	err := d.DialAndSend(m)
	return err
}

func genHtml(body string) string {
	bodyHtml := fmt.Sprintf(`
<h3>博客验证平台邮件</h3>
<p>验证码: %v</p>
<p>发送时间: %v</p>`, body, time.Now().Format("2006-01-02 15:04:06"))
	return bodyHtml
}
