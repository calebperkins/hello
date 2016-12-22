package main

import "testing"

func TestGreeting(t *testing.T) {
	if Greeting("Caleb") != "Hello, Caleb!" {
		t.Error("Expected", "Hello, Caleb!", "; got:", Greeting("Caleb"))
	}
}
