package greetings

import (
	"regexp"
	"testing"
)

func TestSingleGreeting(t *testing.T) {
	name := "Ed"
	want := regexp.MustCompile(`\b` + name + `\b`)
	message, err := SingleGreeting(name)
	if !want.MatchString(message) || err != nil {
		t.Fatalf(`SingleGreeting(string) = %q, %v, want match for %#q, nil`, message, err, want)
	}
}

func TestSingleEmptyGreeting(t *testing.T) {
	emptyName := ""
	message, err := SingleGreeting(emptyName)
	if message != emptyName || err == nil {
		t.Fatalf(`SingleGreeting(string) = %q, %v, want "", error`, message, err)
	}
}

func TestMultipleGreetings(t *testing.T) {
	names := []string{"Ken", "Nick", "Jim"}

	messages, err := MultipleGreetings(names)

	for i := range names {
		want := regexp.MustCompile(`\b` + names[i] + `\b`)
		if !want.MatchString(messages[names[i]]) || err != nil {
			t.Fatalf(`MultipleGreetings(names []string) = %q, %v, want match for %#q, nil`, messages, err, want)
		}
	}

	want := len(names)
	if want != len(messages) || err != nil {
		t.Fatalf(`MultipleGreetings(names []string) = %q, %v, want match for %#q size, nil`, messages, err, want)
	}

}
