package utils

import "log"

func SendEmail(content string) error {
	// Placeholder: Use SendGrid/Mailgun APIs
	log.Println("Sending email: ", content)
	return nil
}

func FetchTargetRecipients() []string {
	// Placeholder: Fetch target recipients
	return []string{"developer1@example.com", "team@techcompany.com"}
}
