package main

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gopal"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gopal")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gopal")  = %q, %v, want match for %#q, nil`, msg, err, want)
	}

}
func TestHelloEmptyName(t *testing.T) {
	name := ""
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("")  = %q, %v, want match for %#q, nil`, msg, err, want)
	}

}
