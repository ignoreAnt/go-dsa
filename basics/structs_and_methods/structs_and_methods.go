package main

import "fmt"

/**
 * @author: Aakash B
 * @date: 2024/11/11 20:15
 * @description: Demonstrates structs and methods in Go.
 * @file: structs_and_methods.go
 * @summary: This example shows how to use structs and methods in Go.
 * @link: https://go.dev/tour/methods

Structs and Methods in Go
---------------------------
	Purpose:
		Structs in Go serve as a way to define complex data types by grouping multiple fields.
		Structs are fundamental in organizing data and are often used to represent
		real-world entities in a structured way. Methods in Go allow functions to be associated with structs,
		enabling object-oriented behavior. Struct embedding provides a mechanism for composition,
		allowing structs to reuse fields and methods from other structs, promoting code reuse.

Core Components:
----------------
	Struct Definition:
		Defines a custom data structure with specific fields.
	Methods with Receivers:
		Attaches functions to struct types, enabling method-like behavior.
	Struct Embedding:
		Allows one struct to embed another, gaining access to the embedded struct's fields and methods.

Explanation:
	Struct Definition:
		The Person struct is defined with fields FirstName, LastName, and Age.
		Each field has a specific type, and all are public (accessible outside the package)
		because they start with an uppercase letter.

	Method with Receiver:
		The FullName method is attached to Person using a receiver (p *Person),
		allowing instances of Person to call this method. The receiver type *Person enables this
		method to access and modify the struct's fields directly, though in this example it only
		reads fields.

		This function concatenates the first and last name, returning the full name.

	Struct Embedding:
		The Employee struct embeds the Person struct, gaining access to its fields and methods.
		In employee.FullName(), FullName() is accessible directly as if it were a method of Employee,
		thanks to struct embedding.
		Embedding promotes code reuse and is Go's way of achieving composition.

Use Cases in Data Structures and Algorithms:
--------------------------------------------
	Organizing Complex Data:
		For data structures like trees, graphs, or linked lists, structs are essential to define nodes
		with specific fields (e.g., pointers to children or neighbors).
	Methods for Structs:
		Methods like AddNode, RemoveNode, FindNode, etc., can be attached to structs to provide
		intuitive interfaces for data structure operations.
	Struct Embedding for Composition:
		Allows flexibility in data structures by embedding base node types and extending them
		for specific implementations (e.g., binary vs. generic trees).

Key Points:
------------

	Encapsulation:
		Group related fields in structs to represent entities (e.g., nodes in data structures).
	Receiver Types:
		Receivers can be pointers (for in-place modifications) or values (for non-modifying actions).
	Composition over Inheritance:
		Go promotes struct embedding for composition, avoiding the need for traditional inheritance.

Practical Usage:
----------------
	Linked List Nodes:
		Use structs to define nodes with data and pointer fields.
	Binary Tree Nodes:
		Structs are ideal for representing nodes with left and right child pointers.
	Graph Nodes and Edges:
		Represent nodes and connections in graph-based algorithms.

Notes:
------
	Structs and methods are Go's primary way to represent objects and encapsulate related behavior.
	Struct Embedding allows Go to achieve code reuse, while avoiding the pitfalls of
	deep inheritance hierarchies common in other languages.

*/

//1. Struct Definition

// Person Define a struct to represent a person
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// 2. Method with Receiver

// FullName returns the full name of the person
func (p *Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// 3. Struct Embedding

// Employee is a struct that embeds Person
type Employee struct {
	Person
	JobTitle string
}

func main() {

	//Initialize a Person struct
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       22,
	}

	//Print the full name of the person
	fmt.Println("Full Name:", person.FullName())

	//Initialize an Employee struct
	employee := Employee{
		Person: Person{
			FirstName: "Jane",
			LastName:  "Smith",
			Age:       25,
		},
		JobTitle: "Marketing Manager",
	}

	//Print the full name of the employee
	fmt.Println("Full Name:", employee.FullName())

	//Print the job title of the employee
	fmt.Println("Job Title:", employee.JobTitle)

}
