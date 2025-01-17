package mailers

import (
	"log"

	"github.com/gobuffalo/buffalo/mail"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/buffalo/worker"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr/v2"
)

var smtp mail.Sender
var r *render.Engine
var Mails = map[string]func(args worker.Args) error{}

const (
	EmailWelcomeJob = "welcome"
)

func init() {

	// Pulling config from the env.
	port := envy.Get("SMTP_PORT", "1025")
	host := envy.Get("SMTP_HOST", "localhost")
	user := envy.Get("SMTP_USER", "")
	password := envy.Get("SMTP_PASSWORD", "")

	var err error
	smtp, err = mail.NewSMTPSender(host, port, user, password)

	if err != nil {
		log.Fatal(err)
	}

	r = render.New(render.Options{
		HTMLLayout:   "layout.plush.html",
		TemplatesBox: packr.New("app:mailers:templates", "../templates/mail"),
		Helpers:      render.Helpers{},
	})

	Mails[EmailWelcomeJob] = SendWelcomeEmails
}
