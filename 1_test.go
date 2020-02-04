package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("wew")
	want := "Hello, Bro"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
