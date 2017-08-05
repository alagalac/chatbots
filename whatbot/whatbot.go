package main

import (
    "fmt"
    "os"
	"regexp"
)

func main() {
	fmt.Printf("Hello, human\n")
	fmt.Printf("My name is- Wait. Can you tell me my name?\n")
	var finished bool = false
	var answer string
	for !finished {
		answer, finished = recieveInput(os.Stdin)
		fmt.Printf(answer + "\n")
	}
}

func recieveInput(input *os.File) (string, bool) {

	text := ""
    fmt.Fscanln(input, &text)

	whatRegex := regexp.MustCompile(`(?i).*what.*`)
	lightbulbRegex := regexp.MustCompile(`(?i).*lightbulb.*`)

	if whatRegex.MatchString(text) {
		return "LIGHTBULB!", true
	} else if lightbulbRegex.MatchString(text) {
		return "Nooooo!!!", true
	}
	
	return "That's not my name!\nWhat is my name?", false
}
