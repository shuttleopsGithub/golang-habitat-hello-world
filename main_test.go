package main

import (
	"testing"
)

func test_message(t testing.T) {
	var expectedMessage string = "Hello World"
	message := message()
	if message != expectedMessage {
		t.Fatalf("Message %s does not match expected message %s", message, expectedMessage)
	}
}