package smtpHelper

import (
	"api/common/utils"
	"api/config"
	"bytes"
	"fmt"
	"net/smtp"
)

// SMTP configuration
type SMTPConfig struct {
	Host         string
	Port         int
	Username     string
	Password     string
	FromEmail    string
	FromUsername string
}

var smtpCfg *SMTPConfig

func instance() *SMTPConfig {
	if smtpCfg == nil {
		smtpCfg = &SMTPConfig{
			Host:     config.Env.SmtpHost,
			Port:     config.Env.SmtpPort,
			Username: config.Env.SmtpUsername,
			Password: config.Env.SmtpPassword,
		}
	}
	return smtpCfg
}

func SendEmailTo(fromEmail string, fromUsername string, recipient string, subject string, body []byte) error {
	newInstance := instance()
	newInstance.FromEmail = fromEmail
	newInstance.FromUsername = fromUsername
	return sendHTMLEmailTo(newInstance, recipient, subject, body)
}

func SendEmailBCC(recipients []string, subject string, body []byte) error {
	newInstance := instance()
	return sendHTMLEmailBCC(newInstance, recipients, subject, body)
}

// sendHTMLEmailTo sends an HTML email using the provided SMTP configuration.
// If username or password are empty, no authentication is used.
func sendHTMLEmailTo(smtpCfg *SMTPConfig, recipient string, subject string, body []byte) error {
	if smtpCfg == nil {
		return fmt.Errorf("SMTP configuration is nil")
	}
	if !utils.IsEmailValid(recipient) {
		return fmt.Errorf("recipient email is invalid: %s", recipient)
	}

	from := fmt.Sprintf("%s <%s>", smtpCfg.FromUsername, smtpCfg.FromEmail)

	headers := map[string]string{
		"From":         from,
		"To":           recipient,
		"Subject":      subject,
		"MIME-Version": "1.0",
		"Content-Type": "text/html; charset=\"UTF-8\"",
	}

	var msg bytes.Buffer
	for k, v := range headers {
		msg.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	msg.WriteString("\r\n") // Separator between headers and body
	msg.Write(body)

	addr := fmt.Sprintf("%s:%d", smtpCfg.Host, smtpCfg.Port)

	// Use authentication only if both username and password are provided
	var auth smtp.Auth
	if len(smtpCfg.Username) > 0 && len(smtpCfg.Password) > 0 {
		auth = smtp.PlainAuth("", smtpCfg.Username, smtpCfg.Password, smtpCfg.Host)
	} else {
		auth = nil
	}

	// Send the email
	return smtp.SendMail(addr, auth, smtpCfg.FromEmail, []string{recipient}, msg.Bytes())
}

// sendHTMLEmailTo sends an HTML email to multiple recipients using BCC,
// hiding their email addresses from each other.
// If username or password are empty, no authentication is used.
func sendHTMLEmailBCC(smtpCfg *SMTPConfig, recipients []string, subject string, body []byte) error {
	if smtpCfg == nil {
		return fmt.Errorf("SMTP configuration is nil")
	}

	from := fmt.Sprintf("%s <%s>", smtpCfg.FromUsername, smtpCfg.FromEmail)

	headers := map[string]string{
		"From": from,
		// "To":           smtpCfg.FromEmail, // To can be the sender's email or left empty
		"Subject":      subject,
		"MIME-Version": "1.0",
		"Content-Type": "text/html; charset=\"UTF-8\"",
	}

	var msg bytes.Buffer
	for k, v := range headers {
		msg.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	msg.WriteString("\r\n") // Separator between headers and body
	msg.Write(body)

	addr := fmt.Sprintf("%s:%d", smtpCfg.Host, smtpCfg.Port)

	// Auth conditionnelle
	var auth smtp.Auth
	if len(smtpCfg.Username) > 0 && len(smtpCfg.Password) > 0 {
		auth = smtp.PlainAuth("", smtpCfg.Username, smtpCfg.Password, smtpCfg.Host)
	} else {
		auth = nil
	}

	// Envoi aux destinataires en BCC
	return smtp.SendMail(addr, auth, smtpCfg.FromEmail, recipients, msg.Bytes())
}
