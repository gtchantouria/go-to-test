package main

import "testing"

func TestSayHi(t *testing.T) {
	expected := "Hi Guille"
	greeting := sayHi("Guille")

	if greeting != expected {
		t.Errorf("Greeting was incorrect, got: '%s', want: '%s'", greeting, expected)
	}
}
