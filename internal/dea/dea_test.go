package dea

import "testing"

func TestIsDEA(t *testing.T) {
	known_dea := "10minutemail.com"
	known_good := "gmail.com"
	if IsDEA(known_good) {
		t.Fatalf("%s is not a DEA", known_good)
	}
	if !IsDEA(known_dea) {
		t.Fatalf("%s should have been DEA", known_dea)
	}
}
