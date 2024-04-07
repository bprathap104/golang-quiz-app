package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var num_questions int = 2
	var correct int = 0
	fmt.Println("Enter your name:")
	var name string = ""
	fmt.Scanln(&name)
	fmt.Printf("Hello %v!\n", name)
	fmt.Println("Enter your age:")
	var age int = 0
	fmt.Scanln(&age)
	if age < 10 {
		fmt.Println("You are not eligible for Quiz")
		return
	}
	fmt.Println("Continue")
	fmt.Println("Quiz1: Which is better, RTX 3080 or RTX 3090?")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	strippedStr := line[:len(line)-1]
	if strippedStr == "RTX 3080" || strippedStr == "rtx 3080" {
		fmt.Println("Correct!")
		correct++
	} else {
		fmt.Println("Incorrect!")
	}
	fmt.Println("Quiz2: Go is a compiled language?")
	var answer2 string = ""
	fmt.Scanln(&answer2)
	if answer2 == "true" {
		fmt.Println("Correct!")
		correct++
	} else {
		fmt.Println("Incorrect!")
	}
	var result float64 = 0
	result = (float64(correct) / float64(num_questions)) * 100
	fmt.Printf("Your result %v%%!\n", result)
}
