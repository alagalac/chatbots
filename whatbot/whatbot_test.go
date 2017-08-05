package main

import (
    "io"
    "io/ioutil"
    "os"
    "testing"
)

func TestChat(t *testing.T) {
	// Create temporary file for us to use
	tmpfile, err := ioutil.TempFile("", "")
    if err != nil {
        t.Fatal(err)
    }

	// Clean up
    //defer tmpfile.Close()
	//defer os.Remove(tmpfile.Name())

	// Begin testing
	_, err = io.WriteString(tmpfile, "whatbot\n")
	if err != nil {
        t.Fatal(err)
    }

	_, err = tmpfile.Seek(0, os.SEEK_SET)
	if err != nil {
        t.Fatal(err)
    }

	result := recieveInput(tmpfile)
	if result != "LIGHTBULB!" {
		t.Error("Unexpected result. Returned:", result)
	}
}