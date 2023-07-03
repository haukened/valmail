package valmail

import "testing"

func TestSMTP(t *testing.T) {
	e := Parse("nobody@gmail.com")
	if !e.SMTPValid() {
		t.Fatalf("unable to test SMTP for %s", e.Domain)
	}
}
