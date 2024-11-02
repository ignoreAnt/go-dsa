package main

import "fmt"

/*
*
Variable Declaration:
---------------------

In Go, variables can be declared with an explicit type or inferred type.
Explicit Declaration: var x int = 10 specifies both the name and type (int) of the variable.
Type Inference: var y = 20.5 allows Go to infer the type based on the value provided (float64 in this case).
Short Declaration: z := "Hello, Go!" uses :=, a shorthand syntax allowed only within functions,
which infers the type based on the assigned value.

Multiple Variable Declaration:
-----------------------------
Go allows multiple variables to be declared on a single line using var a, b, c int = 1, 2, 3.
This can be useful when you need to declare multiple variables of the same type together.

Basic Data Types:
-----------------

Boolean (bool):

	Represents true or false.

Integer (int):

	Go has several integer types (int, int8, int16, int32, int64).
	By default, int is platform-dependent, but it’s typically int64 on a 64-bit system.

Floating-point (float32, float64):

	Go supports floating-point numbers with float32 and float64 types.
	By default, type inference will assign float64 for decimal values.

String (string):

	A sequence of characters.

Points Relevant to Data Structures
------------------------------------

Integer Types:

	When implementing data structures,
	choosing the appropriate integer type (int8, int16, etc.) can help optimize memory usage,
	especially for large data structures.

Boolean:

	Useful for flags in algorithms, like checking if an element has been visited in graph traversal.

String:

	Used for storing character sequences,
	which is useful in algorithms involving text processing (e.g., checking for anagrams or substrings).
	Understanding these types is essential for building data structures like arrays,
	stacks, and linked lists.
	Additionally, Go’s strict typing enforces clarity and prevents certain kinds of bugs,
	which is valuable in complex data structure implementations.
*/
func main() {
	// Explicit variable declaration with type
	var x int = 10
	fmt.Println("Explicitly declared variable:", x)

	// Type inference: Go will infer the type based on the assigned value
	y := 20.5
	fmt.Println("Type inferred variable:", y)

	// Short variable declaration (works only within functions)
	z := "Hello, Go!"
	fmt.Println("Short variable declaration:", z)

	// Multiple variable declaration
	a, b, c := 1, 2, 3
	fmt.Println("Multiple variable declaration:", a, b, c)

	// Basic data types
	var boolVar bool = true                // boolean
	var intVar int = 42                    // integer resolves automatically to architecture-dependent type integer
	var floatVar float64 = 3.14            // floating-point
	var stringVar string = "Hello, World!" // string
	fmt.Println("Basic data types:", boolVar, intVar, floatVar, stringVar)

	// Integer types
	var int8Var int8 = 127                   // 8-bit integer
	var int16Var int16 = 32767               // 16-bit integer
	var int32Var int32 = 2147483647          // 32-bit integer
	var int64Var int64 = 9223372036854775807 // 64-bit integer
	fmt.Println("Integer types:", int8Var, int16Var, int32Var, int64Var)

}
