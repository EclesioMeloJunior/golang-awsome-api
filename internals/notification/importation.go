package notification

import (
	"fmt"
	"log"
)

// Importation interface will abstracts some specific
// notification actions
type Importation interface {
	Notifier
	NotifySuccess(string) error
	NotifyFail(error) error
}

type importation struct {
	Emailer
}

// NewImportationNotifier will return an implementation of Importation
// interface that uses Emailer
func NewImportationNotifier(e Emailer) Importation {
	return &importation{e}
}

func (i *importation) NotifySuccess(message string) error {
	log.Printf("[SUCCESS] Import: %s\n", message)

	notify := "eclesiomelo.1@gmail.com"
	subject := "Successful import"
	emailTemplate := `
		To: %s
		Subject: %s

		The job to import Open Food Facts to mongodb was successful
		with the following message: %s
	`
	parsedTemplate := fmt.Sprintf(emailTemplate, notify, subject, message)

	i.Informations(notify, subject, []byte(parsedTemplate))
	return i.Notify()
}

func (i *importation) NotifyFail(err error) error {
	log.Printf("[FAIL] Import: %v\n", err)

	notify := "eclesiomelo.1@gmail.com"
	subject := "Importation was failed"
	emailTemplate := `
		To: %s
		Subject: %s

		Could not execute import from Open Food Facts data
		into mongo database. The error: %v
	`
	parsedTemplate := fmt.Sprintf(emailTemplate, notify, subject, err)

	i.Informations(notify, subject, []byte(parsedTemplate))
	return i.Notify()
}

func (i *importation) Notify() error {
	return i.Send()
}
