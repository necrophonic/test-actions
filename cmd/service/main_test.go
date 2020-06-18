package main

import "testing"

func TestMsg(t *testing.T) {
	m := msg()

	if m != "This is doing something really cool! Honest!" {
		t.Error()
	}
}
