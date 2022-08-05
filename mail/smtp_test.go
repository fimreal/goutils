package mail

import (
	"testing"
)

func TestSend(t *testing.T) {
	s := &SMTPConfig{
		Username: "123456@qq.com",
		Password: "",
		Host:     "smtp.qq.com",
		Port:     465,
	}

	agent := NewAgent(s)
	// ezap.Info("agent 初始化完成")

	l := &Letter{
		Address:        []string{s.Username},
		Subject:        "Golang SMTP testing",
		Sender:         s.Username,
		SenderNickName: "GoMail Robot",
		Body:           "hello, This is a simple test.",
		BodyType:       "text/plain",
	}
	agent.Put(l)
	// ezap.Info("成功写入信件信息")

	// err := agent.Send()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// ezap.Info("发送成功")
}
