package main

import "fmt"

/**
Functions in Go
---------------

1. Basic Function Definition and Calling
-----------------------------------------
	Purpose:
		Basic functions encapsulate a specific task that can be reused throughout the program.
		Defining a function with the func keyword and calling it allows you to break down complex
		logic into manageable pieces.

	Code Example: add function

		func add(x int, y int) int {
			return x + y
		}

	Explanation:
		add is a function that takes two integers (x and y) and returns their sum.
		It uses the func keyword, followed by the function name, parameters, and return type.

	Use Cases in Data Structures and Algorithms:
		Modularization:
			Basic functions like add are useful in algorithms where operations need to be reused,
			such as summing elements in an array.
		Helper Functions:
			Small utility functions are often used as helpers in larger algorithms,
			such as calculating the sum of node values in a binary tree.

2. Parameters and Return Values
--------------------------------
	Purpose:
		Functions can accept parameters (inputs) and return values (outputs),
		enabling data to be passed in and out. This is foundational for writing reusable code.

	Code Example: subtract function

		func subtract(a int, b int) int {
			return a - b
		}

	Explanation:
		subtract takes two integers as input and returns their difference.
		Parameters allow you to pass data into the function, and the return
		value sends the result back to the caller.

Use Cases in Data Structures and Algorithms:
--------------------------------------------

	Arithmetic Operations:
		Parameterized functions are often used in algorithms where repeated
		calculations are required, such as calculating distances or costs.
	Search and Compare Functions:
		Functions like subtract are helpful for comparing values,
		such as finding the minimum difference in a set of numbers.

3. Multiple Return Values
-------------------------
	Purpose:
		Go functions can return multiple values, which is commonly used for returning
		both a result and an error.

	Code Example: divide function

	func divide(a int, b int) (int, int) {
		quotient := a / b
		remainder := a % b
		return quotient, remainder
	}

	Explanation:
		divide returns both the quotient and remainder when dividing a by b.
		Go’s multiple return values allow functions to return additional information
		without needing complex data structures.

	Use Cases in Data Structures and Algorithms:

		Error Handling:
			Functions can return results along with errors,
			which is crucial in data structures like hash tables where lookups may fail.
		Partitioning:
			Returning multiple values is useful for divide-and-conquer algorithms,
			such as quicksort, where a function may need to return multiple pieces of data.

4. Named Return Values
----------------------
	Purpose:
		Named return values allow you to pre-declare return variables
		in the function signature, making code easier to read and simplifying the return statement.

	Code Example: rectangleProperties function
		func rectangleProperties(length, width int) (area, perimeter int) {
			area = length * width
			perimeter = 2 * (length + width)
			return  // returns area and perimeter automatically
		}
	Explanation:
		rectangleProperties calculates and returns both the area and perimeter of a rectangle.
		Named return values allow us to omit the variables in the return statement,
		as they’re automatically returned.

	Use Cases in Data Structures and Algorithms:
		Geometry Calculations:
			Named return values make functions cleaner when performing calculations,
			like finding the dimensions of geometric shapes.
		Clarity:
			Named return values improve readability, which is especially useful in
			complex algorithms where the purpose of each return value needs to be clear.

5. Variadic Functions
---------------------
	Purpose:
		A variadic function accepts a variable number of arguments, useful when the
		number of inputs isn’t fixed.

	Code Example: sum function

		func sum(nums ...int) int {
			total := 0
			for _, num := range nums {
				total += num
			}
			return total
		}

	Explanation:
		sum can take any number of integers and calculates their total.
		The ... syntax allows for variable arguments, making it flexible for different input sizes.

	Use Cases in Data Structures and Algorithms:
		Aggregation:
			Useful for aggregation functions, like summing values or calculating
			averages in a set of data.
		Flexible Input:
			Variadic functions are ideal when working with lists of data, such as merging arrays.

6. Passing Functions as Parameters
----------------------------------

	Purpose:
		Passing functions as parameters allows you to create higher-order functions,
		where the behavior can be customized by passing different functions.

	Code Example: applyOperation function

		func applyOperation(a, b int, op func(int, int) int) int {
			return op(a, b)
		}

	Explanation:
		applyOperation takes two integers and a function (op) as parameters,
		applying the function to the integers.
		This allows you to customize the behavior by passing different operations.

	Use Cases in Data Structures and Algorithms:
		Custom Comparators:
			Useful for implementing custom sorting or comparison functions.
		Functional Programming:
			Higher-order functions enable functional programming patterns, like map, filter, and reduce.

7. Anonymous Functions
----------------------
	Purpose:
		Anonymous functions are unnamed functions, often used for short, inline operations.

	Code Example: anonymousFunctionExample

		func anonymousFunctionExample() {
			// Assigned to a variable
			greet := func(name string) {
				fmt.Println("Hello,", name)
			}
			greet("Alice")

			// Immediately invoked
			func(msg string) {
				fmt.Println(msg)
			}("This is an immediately invoked function!")
		}

	Explanation:
		The first example assigns an anonymous function to greet and then calls it.
		The second example defines and immediately invokes an anonymous function.

	Use Cases in Data Structures and Algorithms:
		Deferred Execution:
			Anonymous functions are useful in defer statements for cleanup tasks.
		Single-use Operations:
			Ideal for quick, single-use functions, like in event handling or callbacks.

8. Closures
-----------
	Purpose:
		A closure is an anonymous function that "remembers" variables from its outer scope.
		This makes it useful for maintaining state across multiple function calls.

	Code Example: closureExample function
		func closureExample() func() int {
			counter := 0
			return func() int {
				counter++
				return counter
			}
		}

	Explanation:
		closureExample returns an anonymous function that increments counter each time it’s called.
		The anonymous function "closes over" counter, preserving its state between calls.

	Use Cases in Data Structures and Algorithms:
		Stateful Iterators:
			Closures are useful for creating iterators that retain state.
		Function Factories:
			Useful in data structures for generating custom functions, such as generating
			hash functions with specific seeds.

9. Recursion
------------
	Purpose:
		Recursion is a technique where a function calls itself.
		It’s useful for problems that can be broken down into smaller sub-problems,
		such as tree traversal and divide-and-conquer algorithms.

	Code Example: factorial function

		func factorial(n int) int {
			if n == 0 {
				return 1
			}
			return n * factorial(n-1)
		}

	Explanation:

		factorial calculates the factorial of a number by calling itself with a decremented value of n.
		The base case (n == 0) prevents infinite recursion and provides an end condition.

	Use Cases in Data Structures and Algorithms:
		Tree Traversals:
			Recursion is often used for traversing trees (preorder, inorder, postorder).
		Divide and Conquer:
			Recursive functions are used in algorithms like binary search and quicksort.
		Combinatorial Algorithms:
			Many algorithms, such as generating subsets or calculating permutations, rely on recursion.

*/

func main() {

	// 1. Basic Function Definition and Calling
	fmt.Println(add(1, 2))

	// 2. Parameters and Return Values
	fmt.Println(subtract(5, 3))

	// 3. Multiple Return Values
	fmt.Println(divide(10, 5))

	// 4. Named Return Values
	rectangleProperties(5, 3)

	// 5. Variadic Functions
	fmt.Println(sum(1, 2, 3, 4, 5))

	// 6. Passing Functions as Parameters
	fmt.Println(applyOperation(2, 3, multiply))

	// 7. Anonymous Functions
	fmt.Println("\nAnonymous Functions Example:")
	anonymousFunctionExample()

	// 8. Closures
	fmt.Println("\nClosures Example:")
	counter := closureExample()
	fmt.Println("counter : ", counter())
	fmt.Println("counter : ", counter())
	fmt.Println("counter : ", counter())

	// 9. Recursion
	fmt.Println("\nRecursion Example:")
	fmt.Println(factorial(5))

}

// 1. Basic Function Definition and Calling
// add takes two integers as input and returns their sum
func add(x int, y int) int {
	return x + y
}

// 2.Parameters and Return Values
// subtract takes two integers as input and returns their difference
func subtract(a int, b int) int {
	return a - b
}

// 3. Multiple Return Values
// divide returns both the quotient and remainder when dividing a by b
func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// 4. Named Return Values
// rectangleProperties calculates and returns both the area and perimeter of a rectangle
func rectangleProperties(length, width int) (area, perimeter int) {
	area = length * width
	perimeter = 2 * (length + width)
	return // returns area and perimeter automatically
}

// 5. Variadic Functions
// sum takes any number of integers and calculates their total.
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// 6. Passing Functions as Parameters
// applyOperation takes two integers and a function (op) as parameters,
// applying the function to the integers.
func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// multiply takes two integers and returns their product.
func multiply(a, b int) int {
	return a * b
}

// 7. Anonymous Functions
// anonymousFunctionExample demonstrates the use of an anonymous function.
func anonymousFunctionExample() {
	// Assigned to a variable
	greet := func(name string) {
		fmt.Println("Hello,", name)
	}
	greet("Alice")

	func(msg string) {
		fmt.Println(msg)
	}("This is an immediately invoked function!")
}

// 8. Closures
// closureExample returns an anonymous function that increments counter each time it’s called.
func closureExample() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}

// 9. Recursion
// factorial calculates the factorial of a number by calling itself with a decremented value of n.
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
