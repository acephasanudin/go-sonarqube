package main

import "testing"

func TestGetMessage(t *testing.T) {
	expected := "Hello, Go and SonarQube!"
	result := getMessage()

	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
