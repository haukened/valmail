package valmail

import "testing"

func TestPopulateMX(t *testing.T) {
	e := Parse("noreply@gmail.com")
	err := e.populateMX()
	if err != nil {
		t.Fatalf("populate mx returned error: %s", err)
	}
	if !e.MXValid() {
		t.Fatalf("domain %s should have valid DNS MX records", e.Domain)
	}
}

func TestFailMX(t *testing.T) {
	e := Parse("nobody@thisprobablyisnotadomain.com")
	err := e.populateMX()
	if err == nil {
		t.Fatalf("somehow we got MX records for %s and should not have", e.Domain)
	}
	if e.MXValid() {
		t.Fatal("there is NO way that this has valid records")
	}
}

func TestNoMXRecords(t *testing.T) {
	e := Parse("nobody@haukeness.com")
	if e.MXValid() {
		t.Fatalf("%s is a valid domain with no MX records, but package reports valid", e.Domain)
	}
}

func TestInvalidMX(t *testing.T) {
	bad := Parse("foo.co")
	if bad.MXValid() {
		t.Fatal("invalid syntax mail is returning valid MX")
	}
}
