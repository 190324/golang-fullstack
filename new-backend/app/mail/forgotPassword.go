package mail

import (
	"XinAPI/pkg/mail"
	"fmt"
)

func ForgotPassword(to string) {
	fmt.Printf("mail, %s", to)

	p := mail.InputParams{}
	p.Name = "FM KUO"
	p.Email = to
	p.Subject = "Test mail"

	p.Template = "members/forgotPassword.html"
	// p.Template = "M_01"
	p.ContentData = mail.Data{
		Title:   "Hello",
		Content: "Welcome to Xinmedia",
	}
	res, info := mail.Send(p)
	fmt.Println(res)
	fmt.Println(info)
	fmt.Println(p)
}
