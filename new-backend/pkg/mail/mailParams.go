package mail

import (
	"XinAPI/pkg/path"
	"bytes"
	"fmt"
	"html/template"
	"os"

	strip "github.com/grokify/html-strip-tags-go"
	"github.com/spf13/viper"
)

type mailParam struct {
	apiKey      string
	sender      string
	from        string
	subject     string
	receiver    string
	to          string
	textContent string
	htmlContent string
}

const (
	constMailTemplatePath string = "/resources/emails/"
)

func (mp *mailParam) getParams(p InputParams) {
	mp.apiKey = viper.GetString("mail.api_key")
	mp.sender = viper.GetString("mail.from_name")
	mp.from = viper.GetString("mail.from_address")
	mp.subject = p.Subject
	mp.receiver = p.Name
	mp.to = p.Email

	body, err := parseTemplate(p.Template, p.ContentData)
	mp.htmlContent = body

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}

	mp.textContent = strip.StripTags(mp.htmlContent)
}

func parseTemplate(templateName string, data interface{}) (string, error) {
	_file := path.BasePath(fmt.Sprint(constMailTemplatePath, templateName))
	t, err := template.ParseFiles(_file)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	body := buf.String()
	return body, nil
}
