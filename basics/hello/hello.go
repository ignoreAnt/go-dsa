package main

/**
Explanation

Package Declaration:

Every Go file starts with a package declaration.
main is the package used for the entry point of executable programs.
When we build an executable, Go looks for the main package and the main function as the starting point.

Imports:

Libraries are imported using the import keyword.
Go has a built-in library called fmt that handles basic input/output functions,
such as Println to display output.
In this example, we’re using fmt.Println to print "Hello, Go!" to the console.

Functions:

The main function is the entry point of the program.
The func keyword is used to define a function in Go.


Points Relevant to Data Structures

Understanding the basics of a Go program is essential as it forms the foundation for writing larger,
modular code.

Familiarity with the fmt package will help us debug our code,
especially when implementing data structures and algorithms.

For example, we’ll use fmt.Println to print out the structure of arrays,
linked lists, and other data types to verify their state.
*/

// Main function is the entry point of the program
func main() {
	println("Hello, Go language!")
}
