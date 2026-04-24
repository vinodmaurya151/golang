package main

import "fmt"

func main() {

	fmt.Println("=== Declaring variable without initializing values ===")
	var age int
	var name string
	var height1 float32
	var height2 float64
	var isActive bool

	fmt.Println("Age: ", age)
	fmt.Println("Name: ", name)
	fmt.Println("Height1: ", height1)
	fmt.Println("Height2: ", height2)
	fmt.Println("isActive: ", isActive)

	/*
		OT:
		   Age:  0
		   Name:
		   Height1:  0
		   Height2:  0
		   isActive:  false
	*/

	fmt.Println("=== Assigning values ===")
	age = 15
	name = "Vinod Maurya"
	height1 = 15.23
	height2 = 16.526
	isActive = true

	fmt.Println("Age: ", age)
	fmt.Println("Name: ", name)
	fmt.Println("Height1: ", height1)
	fmt.Println("Height2: ", height2)
	fmt.Println("isActive: ", isActive)

	fmt.Println("=== Short hand ===")
	again_age := 32
	again_name := "Alice"
	again_height := 5.33
	again_isAtvice := false

	fmt.Println("Again age:", again_age)
	fmt.Println("Again name", again_name)
	fmt.Println("Again height", again_height)
	fmt.Println("Again isActive: ", again_isAtvice)

	fmt.Println("=== Printf() ===")
	fmt.Printf("Again age: %d\n", again_age)
	fmt.Printf("Again name: %s\n", again_name)
	fmt.Printf("Again height: %.2f\n", again_height)
	fmt.Printf("Again isActive: %t\n", isActive)
	fmt.Printf("Again isActive: %v\n", isActive) // %v can be used for any value

	fmt.Println("=== var() ===")
	// Method 2: Multiple variable declaration with var ().
	// What is var ()?
	// var () is a block syntax that lets you declare multiple variables together in a clean, organized way.
	var (
		var_age   int     = 25   // Has value
		var_name  string         // No value (gets "")
		var_score float64 = 95.5 // Has value
	)

	fmt.Printf("Var age: %d\n", var_age)
	fmt.Printf("Var name: %s\n", var_name)
	fmt.Printf("Var score: %.1f\n", var_score)

	fmt.Println("=== Check type ===")
	fmt.Printf("Data type: %T\n", var_age)
	fmt.Printf("Data type: %T\n", var_name)
	fmt.Printf("Data type: %T\n", var_score)

	/*
		%d = Integer
		%s = String
		%T = Data type check
		%.3f = Roundup to 3 decimal places
		%c = Character
	*/

}
