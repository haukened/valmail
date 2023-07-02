package valmail

import (
	"testing"

	"github.com/haukened/valmail/internal/dea"
)

func TestIsDEA(t *testing.T) {
	known_dea := "10minutemail.com"
	known_good := "gmail.com"
	if dea.IsDEA(known_good) {
		t.Fatalf("%s is not a DEA", known_good)
	}
	if !dea.IsDEA(known_dea) {
		t.Fatalf("%s should have been DEA", known_dea)
	}
}

func TestParse(t *testing.T) {
	good_email := "david@hauken.us"
	e := Parse(good_email)
	if !e.SyntaxValid {
		t.Fatalf("%s should have been valid syntax", good_email)
	}
	if e.Username != "david" {
		t.Fatalf("username should have been david, got %s", e.Username)
	}
	if e.Domain != "hauken.us" {
		t.Fatalf("domain should have been hauken.us, got %s", e.Domain)
	}
}
