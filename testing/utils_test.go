package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	expectedStr := "Hello, Testing!"
	result := hello()
	t.Log(expectedStr)
	if result != expectedStr {
		t.Fatalf("Expected %s, got %s", expectedStr, result)
	}
}
