package main

import (
	"fmt"
	"math"
)

/**
 * @author: Aakash B
 * @date: 2024/11/11 21:21
 * @description: Interfaces in Go.
 * @file: interfaces.go
 * @summary: Interfaces in Go allow you to define contracts for types.
 * @link: https://go.dev/tour/interfaces

Concept/Topic: Interfaces in Go
------------------------------
	Purpose:
		Interfaces in Go enable a form of polymorphism, allowing types to define behavior
		without relying on inheritance. They provide a way to specify methods that a type must implement,
		which makes interfaces valuable for defining abstractions and decoupling code.
		Interfaces facilitate flexible and reusable code by focusing on what a type can do rather
		than what it is.

Core Components:
----------------
	Interface Basics:
		Define a set of methods that a type must implement.
	Type Assertion:
		Allows access to the underlying concrete type from an interface variable.
	Type Switch:
		Evaluates the dynamic type of an interface, providing a way to handle different types.
	Using Interfaces for Abstraction:
		Define functions that operate on general behavior rather than specific types.

Explanation:
------------
	Interface Basics:
		type Shape interface { Area() float64 }: Defines the Shape interface with a single method, Area.
		Any type with an Area method returning a float64 satisfies this interface.
		Circle and Rectangle structs each implement the Shape interface by defining an Area method.

	Type Assertion:
		circle := s.(Circle): Asserts that the Shape interface holds a Circle type.
		If ok is true, the assertion succeeds, and we can access circle.Radius.
		Type assertions allow us to retrieve the concrete type from an interface when needed.

	Type Switch:
		switch t := s.(type): Performs type checking to determine the underlying type of s.
		Different cases (Circle and Rectangle) enable different behavior based on the type.
		With Type switches provide a safer and more concise way to handle multiple types.

Using Interfaces for Abstraction:
---------------------------------
	The printArea function takes any Shape as a parameter, demonstrating polymorphism.
	This decouples the function from concrete types and allows it to operate on any shape
	with an Area method.

Use Cases in Data Structures and Algorithms:
--------------------------------------------
	Abstracting Data Structures:
		Interfaces can define operations (e.g., Push, Pop for stacks) without depending on a specific implementation.
	Polymorphic Algorithms:
		Sorting or processing algorithms can work with a broader range of types by relying on interfaces.
	Type Flexibility:
		Algorithms can process mixed data types (e.g., shapes with different area calculations) by utilizing interfaces.

Key Points:
-----------
	Implicit Implementation:
		Types implement interfaces implicitly by having the required methods.
	Type Assertion vs. Type Switch:
		Use assertions to retrieve specific types and type switches to handle multiple potential types.
	Decoupling Code:
		Interfaces promote loose coupling by enabling functions to accept general
		behaviors rather than specific implementations.

Practical Usage:
----------------
	Data Structure Operations:
		Interfaces can define operations for collections like stacks and queues.
	Dependency Injection:
		Interfaces make testing easier by enabling dependency injection with mock types.
	Flexibility in Algorithms:
		Algorithms that operate on multiple types (e.g., different shapes) can be abstracted with interfaces.
Notes:
------
	Interfaces in Go encourage composition and enable powerful polymorphic behavior
	while maintaining Go’s simplicity and readability.
*/

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// Area Implementing the Shape interface for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area Implementing the Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// printArea Prints the area of any Shape interface.
func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

// typeSwitchExample Demonstrates a type switch to determine the underlying type of a Shape interface.
// It prints out the shape and its dimensions.
func typeSwitchExample(s Shape) {
	switch t := s.(type) {
	case Circle:
		fmt.Println("Circle with radius:", t.Radius)
	case Rectangle:
		fmt.Println("Rectangle with width and height:", t.Width, t.Height)
	default:
		fmt.Println("Unknown shape")
	}
}

// typeAssertion Demonstrates type assertion in Go.
// It takes any Shape and uses type assertion to access the underlying Circle type.
// If the type doesn't match Circle, it prints "Not a Circle!".
func typeAssertion(s Shape) {
	if circle, ok := s.(Circle); ok {
		fmt.Printf("Circle Radius: %.2f\n", circle.Radius)
	} else {
		fmt.Println("Not a Circle!")
	}
}
func main() {
	// Creating instances of Circle and Rectangle
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	// Using printArea to print the areas of the shapes
	printArea(circle)
	printArea(rectangle)

	// Using typeAssertion to access the Circle type
	typeAssertion(circle)
	typeAssertion(rectangle)

	// Using typeSwitchExample to print the shape and its dimensions
	typeSwitchExample(circle)
	typeSwitchExample(rectangle)

}
