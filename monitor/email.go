package monitor

import (
	"log"
	"os"
	"strconv"

	"github.com/maxcnunes/go-uptime-api/monitor/entities"
	"gopkg.in/gomail.v1"
)

type emailConfig struct {
	from     string
	username string
	password string
	host     string
	port     int
}

func getEmailConfig() emailConfig {
	conf := emailConfig{
		from:     os.Getenv("EMAIL_FROM"),
		username: os.Getenv("EMAIL_USERNAME"),
		password: os.Getenv("EMAIL_PASSWORD"),
		host:     os.Getenv("EMAIL_HOST"),
	}

	if conf.from == "" || conf.username == "" || conf.password == "" {
		log.Fatalln("Missing email configurations")
	}
	if conf.host == "" {
		conf.host = "smtp.gmail.com"
	}
	if os.Getenv("EMAIL_PORT") == "" {
		conf.port = 587
	} else {
		if port, err := strconv.Atoi(os.Getenv("EMAIL_PORT")); err == nil {
			conf.port = port
		}
	}

	return conf
}

// SendNotificaton sends notifications to all emails related to a target
// The notification can be about a uptime or downtime depending in the current target's status
func SendNotificaton(target entities.Target) {
	if len(target.Emails) == 0 {
		return
	}

	conf := getEmailConfig()

	msg := gomail.NewMessage()
	msg.SetHeader("From", conf.from)
	msg.SetHeader("To", target.Emails...)
	if target.Status < 500 {
		msg.SetHeader("Subject", "Target is UP: "+target.URL)
		msg.SetBody("text/html", "Hi,<br>The target "+target.URL+" is back UP (HTTP "+strconv.Itoa(target.Status)+").")
	} else {
		msg.SetHeader("Subject", "Target is DOWN: "+target.URL)
		msg.SetBody("text/html", "The target "+target.URL+" is currently DOWN (HTTP "+strconv.Itoa(target.Status)+").<br>Uptime Robot will alert you when it is back up.")
	}

	mailer := gomail.NewMailer(conf.host, conf.username, conf.password, conf.port)
	if err := mailer.Send(msg); err != nil {
		panic(err)
	}
}
