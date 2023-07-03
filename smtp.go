package valmail

import (
	"fmt"
	"net/smtp"
	"strings"
)

func (e *EmailAddress) SMTPValid(useTLS ...bool) bool {
	// We can't do this without valid MX records
	if !e.MXValid() {
		return false
	}
	// grab the optional parameter
	port := 587
	if len(useTLS) > 0 {
		if !useTLS[0] {
			port = 25
		}
	}
	// validate each MX record by connecting and saying hello
	for _, mx := range e.mxRecords {
		// remove the trailing period, and add the port number
		addr := fmt.Sprintf("%s:%d", strings.TrimSuffix(mx.Host, "."), port)
		// Dial
		client, err := smtp.Dial(addr)
		if err != nil {
			return false
		}
		// Say HELO/EHLO
		err = client.Hello("localhost")
		if err != nil {
			return false
		}
		// QUIT
		client.Quit()
	}
	// if they all validate, return true
	return true
}
