package valmail

import (
	"net"
	"net/mail"
	"strings"

	"github.com/haukened/valmail/internal/dea"
)

// EmailAddress holds a parsed address for functional validation
type EmailAddress struct {
	Username    string    // the portion of the address preceeding the last @
	Domain      string    // the remainder of the address, following the last @
	Raw         string    // the raw address, as input by you
	SyntaxValid bool      // returns true if the email is syntactically correct
	mxRecords   []*net.MX // holds MX records for further validation
}

// Parse takes a raw address string, and returns an EmailAddress struct for further validation
func Parse(address string) (email EmailAddress) {
	email.Raw = address
	addr, err := mail.ParseAddress(address)
	if err != nil {
		email.SyntaxValid = false
		return
	}
	email.SyntaxValid = true
	email.Username, email.Domain = parseEmail(addr.Address)
	return
}

// parseEmail just splits the address at the final @ in the address
func parseEmail(addr string) (user, domain string) {
	at := strings.LastIndex(addr, "@")
	if at >= 0 {
		user, domain = addr[:at], addr[at+1:]
	}
	return
}

func (e *EmailAddress) IsDEA() bool {
	return dea.IsDEA(e.Domain)
}
