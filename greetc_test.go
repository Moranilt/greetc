package greetc

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Jedi"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello(Jedi) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
