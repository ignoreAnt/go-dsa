package main

import (
	"errors"
	"fmt"
)

/**
 * @author: Aakash B
 * @date: 2024/11/11 20:15
 * @description: Demonstrates error handling in Go.
 * @file: error_handling.go
 * @summary: This example shows how to handle errors in Go.
 * @link: https://go.dev/tour/error-handling

Concept/Topic: Error Handling in Go
-----------------------------------
	Purpose:
		Error handling in Go is straightforward and encourages explicit handling of errors,
		promoting robust code. Go uses the error type to signal issues, and by leveraging
		custom error types and idiomatic patterns, developers can manage errors effectively,
		ensuring graceful handling of unexpected situations in applications and libraries.

Core Components:
----------------
	Using error Types:
		Go's built-in error type represents errors, with functions often returning error
		as an additional return value.
	Custom Error Handling:
		Define custom error types to provide more detailed context and classification
		for specific error scenarios.
	Idiomatic Error Handling:
		Follow Go conventions like handling errors directly after they occur and
		avoiding exceptions for control flow.

Explanation:
-----------
	Using the Built-in error Type:
		Go’s standard library provides the error interface, which has a single method, Error() string.
		errors.New("cannot divide by zero"): Creates a new error with a descriptive message.
		The divide function returns an error if the divisor (b) is zero; otherwise,
		it returns the result and a nil error.

	Custom Error Handling:
		Custom error types provide more context by embedding additional fields.
		DivisionError struct defines a custom error with fields for dividend, divisor, and an error message.
		Implementing the Error() method allows DivisionError to satisfy the error interface,
		enabling it to be returned as an error in Go.

	Idiomatic Error Handling:
		Explicit Checking:
			Go idioms favor checking and handling errors immediately after they occur rather
			than deferring error handling.
		Minimal Panic:
			Avoid panic for error handling in production code; it is reserved for truly exceptional,
			unrecoverable errors.
		Error Wrapping:
			Use fmt.Errorf("...: %w", err) to wrap errors with additional context when propagating them.

Use Cases in Data Structures and Algorithms:
--------------------------------------------
	Error Reporting in Operations:
		Functions like searching in linked lists or binary trees can return custom errors for cases
		where the item is not found.
	File and Network Operations:
		Error handling is essential for I/O operations, where file or network errors are common.
	Error Propagation in Recursion:
		Recursive algorithms can propagate errors back up the call stack, halting execution
		when an error is encountered.

Key Points:
-----------
	Simple Errors:
		Use errors.New or fmt.Errorf for simple error messages.
	Custom Error Types:
		Define custom error types when more context is needed.
	Idiomatic Handling:
		Always check errors immediately, and avoid excessive use of panic in favor of explicit error returns.

Practical Usage:
----------------
	API Responses:
		Handle API response errors with descriptive messages and error types.
	Data Processing:
		Use custom errors to identify specific issues in data, such as validation errors.
	Chaining Errors:
		Add context to errors at each layer using error wrapping.

Notes:
------
	Idiomatic error handling in Go emphasizes simplicity and clarity,
	making it easier to trace issues. Custom errors should be used judiciously
	to avoid overcomplicating error handling, maintaining Go’s philosophy of clean, understandable code.
*/

// 1. Using the Built-in error Type
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// 2. Custom Error Handling

// DivisionError Define a custom error type
type DivisionError struct {
	dividend int
	divisor  int
	err      error
}

// Error implements the error interface, returning a formatted string
// describing the division error.
func (de DivisionError) Error() string {
	return fmt.Sprintf("cannot divide %d by %d: %v", de.dividend, de.divisor, de.err)
}

// divideWithCustomError returns the result of a/b, or an error if b is zero.
func divideWithCustomError(a, b int) (int, error) {
	if b == 0 {
		return 0, DivisionError{a, b, errors.New("cannot divide by zero")}
	}
	return a / b, nil
}

// main demonstrates idiomatic error handling with both the built-in error type and a custom error type.
func main() {
	// 1. Using the Built-in error Type
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// 2. Custom Error Handling
	result, err = divideWithCustomError(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
