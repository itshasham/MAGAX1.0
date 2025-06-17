package mailer

import (
	"fmt"
	"log"

	"gopkg.in/gomail.v2"
)

type SMTPMailer struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

func NewSMTPMailer(host string, port int, username, password, from string) *SMTPMailer {
	// Debugging config values

	return &SMTPMailer{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		From:     from,
	}
}

func (m *SMTPMailer) SendContactConfirmation(to, name, subject, message string) error {
	// Validate SMTP config before dialing
	if m.Host == "" || m.Username == "" || m.Password == "" || m.From == "" {
		return fmt.Errorf("❌ Missing SMTP config (host/user/pass/from)")
	}

	// Prepare message body
	body := fmt.Sprintf(
		"Hi %s,\n\n"+
			"Thank you for contacting Magmox. Your message has been received, and our team will respond as soon as possible.\n\n"+
			"📩 Subject: %s\n\n"+
			"📝 Message:\n%s\n\n"+
			"Warm regards,\nMagmox Team\nhttps://magmox.com",
		name, subject, message,
	)
	msg := gomail.NewMessage()
	msg.SetHeader("From", m.From) // should be: no-reply@magmox.com
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "We've received your message – Magmox Support")
	msg.SetBody("text/plain", body)

	// Create SMTP dialer
	dialer := gomail.NewDialer(m.Host, m.Port, m.Username, m.Password)

	// Use SSL only if you're using port 465
	if m.Port == 465 {
		dialer.SSL = true
	}

	log.Printf("📧 Attempting to send email to %s via %s:%d", to, m.Host, m.Port)

	if err := dialer.DialAndSend(msg); err != nil {
		log.Printf("❌ EMAIL SEND FAILED: %v", err)
		return err
	}

	log.Printf("✅ Email sent successfully to %s", to)
	return nil
}
