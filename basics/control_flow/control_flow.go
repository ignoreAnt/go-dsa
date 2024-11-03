package main

import "fmt"

/**
 * @author: Aakash B
 * @date: 2024/11/02 22:50
 * @description: Demonstrates control flow statements in Go.
 * @file: control_flow.go
 * @summary: This example shows how to use if, else, and switch statements in Go.
 * @link: https://go.dev/tour/controlflow

Control Flow in Go
------------------
	Control flow in Go enables you to define the logic and flow of your program,
	allowing you to make decisions, repeat actions, and branch code paths.
	This is crucial when working with data structures and algorithms,
	as control flow structures allow us to implement conditions, loops, and branching.

Overview of Control Flow Constructs
-----------------------------------
	Control flow constructs in Go include:
		- If Statements: Conditional execution of code blocks.
 		- For Loops: Repeating actions multiple times. Go’s only loop construct,
					 which can function as a traditional loop, a while loop, or an infinite loop.
		- Switch Statements: Branching based on multiple conditions in a concise way.

	Each control flow construct is explained with syntax, use cases,
	and examples relevant to data structures and algorithms.

1. If Statements
----------------
	Purpose: if statements allow you to conditionally execute blocks of code based
			 on specified conditions. This is fundamental for decision-making in programming.

	if condition {
		// code to execute if condition is true
	} else if anotherCondition {
		// code to execute if anotherCondition is true
	} else {
		// code to execute if no condition is true
	}

Use Cases in Data Structures and Algorithms

	Boundary Checks:
		Checking if an index is within array bounds.
	Conditional Insertion/Deletion:
		For example, deciding where to insert a node in a binary search tree.
	Exit Conditions:
		Useful for recursive functions or loops to exit when a specific
		condition is met (e.g., finding a target value in a linked list).

2. For Loops
------------
	Purpose: for is the only loop in Go, but it’s very flexible.
			 It can be used as a traditional loop, a while-like loop, or an infinite loop.

Syntax
	Standard For Loop

	for initialization; condition; post {
		// code to execute
	}

	While-Style For Loop

	for condition {
		// code to execute while condition is true
	}

	Infinite Loop

	for {
		// code to execute indefinitely
	}

Use Cases in Data Structures and Algorithms
-------------------------------------------
	Traversing Arrays or Linked Lists:
		Iterating through arrays or linked lists to perform operations like insertion, deletion, or search.
	Repeated Operations:
		For example, summing values in an array or counting occurrences.
	Loop with Condition:
		Used in binary search or similar algorithms to run until a specific condition is met.
	Infinite Loop:
		Common in event-driven programming or algorithms that wait for a specific condition.

3. Switch Statements
--------------------
	Purpose:
		switch provides a clean way to handle multiple conditions.
		It is particularly useful when you need to branch based on the
		value of a variable that can take on a limited set of possible values.

	Syntax:

		switch expression {
		case value1:
			// code for value1
		case value2:
			// code for value2
		default:
			// code if no cases match
		}

Use Cases in Data Structures and Algorithms
-------------------------------------------
	Decision Making:
		Efficiently handle multiple conditions, which can replace complex if-else chains.
	Handling Different Operations:
		For example, in a data structure that has different operations based on user input (insert, delete, find).
	Categorizing Values:
		Used to classify or categorize data, such as checking if a character is a vowel or consonant.

*/

func main() {
	// 1. If Statements
	fmt.Println("If Statement Example:")
	age := 20
	if age >= 18 {
		fmt.Println("You are eligible to vote.")
	} else {
		fmt.Println("You are not eligible to vote.")
	}

	// 2. Short If Statement with initialization
	if score := 90; score > 80 {
		fmt.Println("Grade A")
	} else if score >= 75 {
		fmt.Println("Grade B")
	} else {
		fmt.Println("Grade C")
	}

	// 3. For Loop
	fmt.Println("For Loop Example:")

	// Basic For Loop
	for i := 0; i < 5; i++ {
		fmt.Println("i : ", i)
	}

	// For loop as a while loop
	j := 0
	for j < 3 {
		fmt.Println("j: ", j)
		j++
	}

	// Infinite loop for with break
	k := 0
	for {
		fmt.Println("Infinite loop, k: ", k)
		k++
		if k == 2 {
			break
		}
	}

	// 3.Switch Statements
	fmt.Println("\nSwitch Case examples: ")
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Unknown day")
	}

	// Switch with multiple values in a case
	letter := "a"
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("Vowel")
	default:
		fmt.Println("Consonant")
	}
}
