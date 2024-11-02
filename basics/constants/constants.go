package main

import "fmt"

/**
Overview of Constants
---------------------
	In Go, constants are declared using the const keyword,
	and they must be assigned a value at the time of declaration.
	Unlike variables, constants:
	 - Cannot be modified once assigned.
	 - Must have values that are known at compile-time (they cannot be calculated at runtime).
	 - Can be typed or untyped.

Explanation of Each Section
---------------------------
1. Basic Constant
	const Pi = 3.14159
	Pi is an untyped constant. It doesn’t have a specific type until it is used in an expression.
	Untyped constants are more flexible because Go can adapt the type based on the context
	in which the constant is used. This is especially useful in mathematical calculations,
	as you might use Pi in different types of calculations without worrying about explicit
	type conversions.

2. Typed Constant
	const Threshold float64 = 0.75
	Threshold is a typed constant with an explicit type of float64.
	Typed constants have a specific type, which means they are restricted to contexts where
	that type is allowed. This can be useful when you need strict type checking,
	for instance, when comparing Threshold with other float64 values in your algorithms.

3. Grouped Constants
	const (
		Width  int = 1024
		Height int = 768
	)

	Grouped constants are a way to declare multiple constants of similar purpose together.
	Here, Width and Height represent dimensions that could be useful in an application where
	fixed dimensions are required. Grouping constants helps with organization and readability,
	especially when you have multiple related constants.

4. Derived Constants

	const (
		Area       = Width * Height
		HalfHeight = Height / 2
		Greeting   = "Hello, " + "Go!"
	)

	Derived constants are constants whose values are derived from other constants through
	compile-time expressions. Area and HalfHeight demonstrate mathematical operations with
	integer constants. Greeting is a string constant created by concatenating two string literals.
	These values are calculated at compile time, meaning they don’t incur any runtime cost,
	which can help optimize performance in critical parts of your program.

5. Enumerated Constants using iota

	const (
		Red = iota
		Green
		Blue
	)

	Enumerated constants use iota, which is a special identifier in Go that simplifies
	the creation of sequential values. In this example,
	Red is assigned 0,
	Green is assigned 1,
	and Blue is assigned 2.
	iota starts at 0 for the first constant in the block and increments by 1 for each
	successive constant. Enumerations are helpful in data structures when you need symbolic names
	for sequential values.
	For example, iota could be used to represent colors, days of the week, or status codes.

6. Complex Enumeration Patterns with iota

	const (
		One = 1 << iota  // 1 << 0 = 1
		Two              // 1 << 1 = 2
		Four             // 1 << 2 = 4
		Eight            // 1 << 3 = 8
	)

	This example uses iota with bit-shifting to create a sequence of powers of two: 1, 2, 4, 8.
	1 << iota means "shift the number 1 to the left by iota positions."
	Each time iota increments, it shifts the value one more position to the left.
	This pattern is useful for creating flags in bitwise operations,
	often used in data structures to represent multiple states in a compact way.
	For example, these constants could represent binary flags in a bit mask.

Points Relevant to Data Structures and Algorithms
-------------------------------------------------

Efficiency:
	Constants allow values to be computed at compile-time rather than runtime,
	which is useful in performance-critical parts of your code,
	such as fixed-size buffers or frequently used mathematical constants.

Code Clarity and Safety:
	By using constants, we signal to other developers (and to ourselves)
	that certain values are not meant to be changed. This can help avoid bugs caused
	by accidental modification.

Enumerations for Readability:
	Using iota for enumerations makes your code more readable.
	For instance, using Red, Green, and Blue instead of 0, 1, and 2 improves clarity
	when dealing with symbolic values in algorithms.

Bitwise Flags:
	iota with bit-shifting allows for compact and efficient representation of
	multiple states using bitwise flags, which can be useful in advanced
	data structures and algorithms (e.g., representing sets, states, or permissions in a single integer).


*/

// Pi 1. Basic Constant
const Pi = 3.14159265358979323846 // Untyped constant for Pi

// Threshold 2. Typed Constant
const Threshold float64 = 0.75 // Typed constant for Threshold

// 3. Grouped Constants
const (
	Width  int = 1024
	Height int = 768
)

// 4. Derived Constants
const (
	Area       = Width * Height
	HalfHeight = Height / 2
	Greeting   = "Hello, " + "Go!"
)

// 5. Enumerated Constants using iota
const (
	Red = iota
	Green
	Blue
)

// 6. Complex Enumeration Patterns with iota
const (
	One   = 1 << iota // 1 << 0 = 1
	Two               // 1 << 1 = 2
	Four              // 1 << 2 = 4
	Eight             // 1 << 3 = 8
)

// Prints examples of different types of constants:
// Basic constants, typed constants, grouped constants,
// derived constants, and complex enumeration patterns using iota.
func main() {
	// Print basic and typed constants
	fmt.Println("Pi:", Pi)
	fmt.Println("Threshold:", Threshold)

	// Print grouped constants
	fmt.Println("Screen Width:", Width)
	fmt.Println("Screen Height:", Height)

	// Print derived constants
	fmt.Println("Area:", Area)
	fmt.Println("HalfHeight:", HalfHeight)
	fmt.Println("Greeting:", Greeting)

	// Print enumerated constants using iota
	fmt.Println("Colors:", Red, Green, Blue)

	// Print complex enumeration patterns with iota
	fmt.Println("Power Constants:", One, Two, Four, Eight)

}
