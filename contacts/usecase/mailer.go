package usecase

type Mailer interface {
	SendContactConfirmation(toEmail, name, subject, message string) error
}
