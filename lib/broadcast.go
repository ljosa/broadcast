package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/mail"

	"gopkg.in/gomail.v2"
)

type Recipient struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
}

type Spec struct {
	FromName   string `json:"from_name"`
	FromAddr   string `json:"from_addr"`
	Subject    string `json:"subject"`
	Html       string `json:"html"`
	Text       string `json:"text"`
	Recipients []Recipient
}

func LoadSpec(specFilename string) (Spec, error) {
	bytes, err := ioutil.ReadFile(specFilename)
	if err != nil {
		return Spec{}, err
	}
	var spec Spec
	err = json.Unmarshal(bytes, &spec)
	if err != nil {
		return Spec{}, err
	}
	return spec, nil
}

func Send(spec Spec, smtpServer string, smtpPort int) error {
	for i := 0; i < len(spec.Recipients); i++ {
		r := spec.Recipients[i]
		fmt.Println(r.Addr)

		m := gomail.NewMessage()
		if spec.FromName != "" {
			from := mail.Address{spec.FromName, spec.FromAddr}
			m.SetHeader("From", from.String())
		} else {
			m.SetHeader("From", spec.FromAddr)
		}
		m.SetHeader("List-Unsubscribe", "<mailto:"+spec.FromAddr+">")
		m.SetHeader("Precedence", "bulk")
		if r.Name != "" {
			to := mail.Address{r.Name, r.Addr}
			m.SetHeader("To", to.String())
		} else {
			m.SetHeader("To", r.Addr)
		}
		m.SetHeader("Subject", spec.Subject)
		if spec.Html != "" && spec.Text != "" {
			m.SetBody("text/plain", spec.Text)
			m.AddAlternative("text/html", spec.Html)
		} else if spec.Html != "" {
			m.SetBody("text/html", spec.Html)
		} else {
			m.SetBody("text/plain", spec.Text)
		}

		d := gomail.NewDialer(smtpServer, smtpPort, "", "")
		err := d.DialAndSend(m)
		if err != nil {
			return err
		}
	}
	return nil
}
