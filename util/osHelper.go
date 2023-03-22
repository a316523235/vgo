package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadLine reads a line of text from the console.
func ReadLine(prompt ...string) string {
	if len(prompt) > 0 {
		fmt.Print(prompt[0])
	} else {
		fmt.Print("Please enter a line of text: ")
	}
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// ReadInt reads an integer from the console.
func ReadInt(prompt ...string) int {
	if len(prompt) == 0 {
		prompt = []string{"Please enter an integer: "}
	}
	for {
		text := ReadLine(prompt...)
		num, err := strconv.Atoi(text)
		if err == nil {
			return num
		}
		fmt.Println("Invalid input. Please enter an integer.")
	}
}

// ReadFloat reads a float from the console.
func ReadFloat(prompt ...string) float64 {
	if len(prompt) == 0 {
		prompt = []string{"Please enter a float:"}
	}
	for {
		text := ReadLine()
		num, err := strconv.ParseFloat(text, 64)
		if err == nil {
			return num
		}
		fmt.Println("Invalid input. Please enter a float.")
	}
}
// ReadMultiLine reads multiple lines of text from the console.
func ReadMultiLine(prompt ...string) string {
	if len(prompt) > 0 {
		fmt.Print(prompt[0])
	} else {
		fmt.Println("Please enter multiple lines of text, and press enter twice to finish:")
	}
	var lines []string
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		if text == "\n" {
			break
		}
		lines = append(lines, text)
	}
	return strings.Join(lines, "")
}
