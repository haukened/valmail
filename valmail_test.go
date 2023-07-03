package valmail

import (
	"testing"
)

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

func TestInvalidSyntax(t *testing.T) {
	e := Parse("foo.com")
	if e.SyntaxValid {
		t.Fatalf("this email should not be valid, as it has no username")
	}
}

func TestIsDEA(t *testing.T) {
	good := Parse("noreply@gmail.com")
	bad := Parse("foo@10minutemail.com")
	if good.IsDEA() {
		t.Fatalf("%s should not be DEA", good.Domain)
	}
	if !bad.IsDEA() {
		t.Fatalf("%s should be DEA", bad.Domain)
	}
}
