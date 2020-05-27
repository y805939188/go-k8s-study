package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("ding")
	want := "Hello, world" + "ding"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}