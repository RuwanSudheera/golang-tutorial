package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "This is user input section welcom"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for course : ")

	// comma ok syntax // err err syntax
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rate, ", input)

}
