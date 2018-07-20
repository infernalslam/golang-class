package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "tak")
	got := buffer.String()
	want := "Hello, tak"

	if got != want {
		t.Errorf("got %d want %d,", got, want)
	}
}
