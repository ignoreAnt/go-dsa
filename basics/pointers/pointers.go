package main

import "fmt"

/**
 * @author: Aakash B
 * @date: 2024/11/11 20:15
 * @description: Demonstrates pointers in Go.
 * @file: pointers.go
 * @summary: This example shows how to use pointers in Go.
 * @link: https://go.dev/tour/pointers

Concept/Topic: Pointers in Go
-----------------------------
	Purpose:
		Pointers in Go provide a way to directly reference the memory address of variables.
		This capability is crucial for modifying values at specific memory addresses,
		achieving efficient memory usage, and managing resources in a program.
		Pointers are especially useful in data structures like linked lists and trees,
		where nodes must be linked using references.

	Core Components:
		Declaration:
			Define a pointer to a type using the * operator.
		Dereferencing:
			Access or modify the value at the pointer’s address.
		Nil Pointers:
			Pointers that don’t point to any valid memory location,
			useful for handling optional or missing data.
		Passing Pointers to Functions:
			Allows modification of the original variable without copying it.
	Explanation:

Pointer Declaration and Initialization:
---------------------------------------

	var p *int = &x:
		Declares p as a pointer to an integer and assigns it the address of x using &x.
		The & operator retrieves the memory address of x.
	Dereferencing:
		*p: Dereferences the pointer p to access the value of x.
		Using *p = 50 changes x to 50 through the pointer.
	Nil Pointers:
		A pointer declared without initialization (like var nilPointer *int) defaults to nil.
		This is useful for optional data, as nil can signify an absence of a value.

Passing Pointers to Functions:
-------------------------------

	modifyValue(&x):
		Passing the address of x to modifyValue.
		Inside the function, the pointer *ptr modifies the original x,
		avoiding a copy and directly altering the value.

Use Cases in Data Structures and Algorithms:
--------------------------------------------
	Linked Lists:
		Nodes are connected through pointers.
		Each node has a pointer to the next node, enabling dynamic list growth.

	Tree Structures:
		Tree nodes use pointers to link to child nodes, forming hierarchical data structures.

	Efficient Parameter Passing:
		Passing large data structures (like arrays) by reference is more efficient than copying,
		achieved by using pointers.

Key Points:
-----------
	Memory Efficiency:
		Using pointers helps avoid copying large amounts of data.
	Direct Modification:
		Pointers allow functions to modify the original variable's value without returning it.
	Nil Pointers:
		Useful to check if a pointer has been initialized or if data is missing.

Practical Usage:
----------------
	Data Structure Linking:
		Essential for dynamic structures like linked lists, trees, and graphs.
	In-Place Modifications:
		Modify large structures (e.g., arrays) within functions without creating additional copies.
	Resource Management:
		Allows developers to directly manage memory, providing more control in high-performance applications.
Notes:
------
	Pointers in Go are straightforward, with limited operations compared to languages like C.
	Go restricts pointer arithmetic, making pointer use safer and less prone to errors.

*/

func main() {
	//1. Pointer Declaration and Initialization
	var x int = 10
	var p *int = &x
	*p = 50
	fmt.Println("Value of x:", x)
	fmt.Println("Pointer p points to x:", p)
	fmt.Println("Value at pointer p (dereferenced):", *p)

	//2. Passing Pointers to Functions
	*p = 100
	fmt.Println("Updating value of x through pointer", x)

	//3. Nil Pointers
	var nilPointer *int
	fmt.Println("Nil pointer:", nilPointer)
	if nilPointer == nil {
		fmt.Println("nilPointer is nil")
	}

	//4. Pointer Arithmetic
	var y int = 20
	var q *int = &y
	var r *int = &y
	*q = 30
	*r = 40
	fmt.Println("Value of y:", y)
	fmt.Println("Pointer q points to y:", q)
	fmt.Println("Pointer r points to y:", r)
	fmt.Println("Value at pointer q (dereferenced):", *q)
	fmt.Println("Value at pointer r (dereferenced):", *r)

	//5. Passing Pointers to Functions
	modifyValue(&x)
	fmt.Println("Value of x after calling modifyValue:", x)
}

func modifyValue(i *int) {
	*i = *i + 10
}
