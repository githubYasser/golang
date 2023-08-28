package main

import (
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Glays"
	want := name
	got, _ := Hello("la" + name)
	if want != got {
		// add t.fatal to stop the test when it fails
		t.Fatalf("got %q want %q", got, want)

	}

}
