package main

import (
    "fmt"
    "os"
	"regexp"
)

func main() {
	fmt.Printf("Hello, human\n")
	fmt.Printf("My name is- Wait. Can you tell me my name?\n")
    fmt.Printf(recieveInput(os.Stdin) + "\n")
}

func recieveInput(input *os.File) string {

	text := ""
    fmt.Fscanln(input, &text)

	whatRegex := regexp.MustCompile(`(?i).*what.*`)

	if whatRegex.MatchString(text) {
		return "LIGHTBULB!"
	}
	
	return "Huh?"
}
