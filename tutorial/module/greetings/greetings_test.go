package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Minzhi"
	want := regexp.MustCompile(`\b` + name + `\b`)
	message, err := Hello(name)
	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello(%v) = %q, %v, want match for %#q, <nil>`, name, message, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")
	if message != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, message, err)
	}
}
