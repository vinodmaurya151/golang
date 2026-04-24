package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		fmt.Println("=== Scan() ===")
		// fmt package provide us Scan() which helps us to scan the input

		fmt.Println("=== Enter your name ===")
		var name string
		fmt.Scan(&name) // If we pass "Vinod Maurya" as name then it will take word only "Vinod" because after "Vinod" we have used space then "Maurya". Ex: "Vinod Maurya" so in Scan() it will take word till it didn't find the space so this is not the correct way
		fmt.Println("Welcome!: ", name)
	*/

	// In order to read input correctly we will use bufio.NewReader(os.Stdin)
	fmt.Println("=== bufio.NewReader(os.Stdin) ===")
	fmt.Println("=== Enter your name ===")

	// reader := bufio.NewReader(os.Stdin) // bufio.NewReader() use to reader the input os means operating system & Stdin means standard input

	// name, _ := reader.ReaderString('\n') // Means it will read the string till it didn't get '\n' new line. the _ is the blank identifier, and it’s used when you want to ignore a value. _ receives the error, but immediately discards it

	reader := bufio.NewReader(os.Stdin)
	name1, _ := reader.ReadString('\n')

	fmt.Println("Hello!", name1)

}
