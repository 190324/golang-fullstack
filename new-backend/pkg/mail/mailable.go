package mail

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func Send(p InputParams) (bool, string) {
	chk, err := checkParams(p)
	if chk == false {
		return false, err
	}

	return doSend(p)
}

func checkParams(p InputParams) (bool, string) {
	if p.Name == "" {
		return false, "收件人姓名為必填"
	}

	if p.Email == "" {
		return false, "收件人 email 為必填"
	}

	if p.Subject == "" {
		return false, "email 主旨為必填"
	}

	if p.Template == "" {
		return false, "email 版型為必填"
	}

	return true, ""
}

func doSend(p InputParams) (bool, string) {
	mp := mailParam{}
	mp.getParams(p)

	from := mail.NewEmail(mp.sender, mp.from)
	subject := mp.subject
	to := mail.NewEmail(mp.receiver, mp.to)
	message := mail.NewSingleEmail(from, subject, to, mp.textContent, mp.htmlContent)
	client := sendgrid.NewSendClient(mp.apiKey)
	_, err := client.Send(message)
	if err != nil {
		return false, err.Error()
	}

	return true, ""
}
