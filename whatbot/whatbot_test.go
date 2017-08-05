package main

import (
    "io"
    "io/ioutil"
    "os"
    "testing"
)

func TestChat(t *testing.T) {
	verifyInput(t, "what", "LIGHTBULB!")
	verifyInput(t, "whatbot", "LIGHTBULB!")
	verifyInput(t, "why", "That's not my name!\nWhat is my name?")
}

func verifyInput(t *testing.T, input string, output string) {
	// Create temporary file for us to use
	tmpfile, err := ioutil.TempFile("", "")
    if err != nil {
        t.Fatal(err)
    }

	// Clean up
    defer tmpfile.Close()
	defer os.Remove(tmpfile.Name())

	// Begin testing
	_, err = io.WriteString(tmpfile, input)
	if err != nil {
        t.Fatal(err)
    }

	_, err = tmpfile.Seek(0, os.SEEK_SET)
	if err != nil {
        t.Fatal(err)
    }

	result, _ := recieveInput(tmpfile)
	if result != output {
		t.Errorf("Unexpected result. Entered: %q, Expected: %q, Actual: %q", input, output, result)
	}
}