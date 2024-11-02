# go-dsa

> A comprehensive repository of Data Structures and Algorithms in Go.

This project is a structured, in-depth resource on Data Structures and Algorithms (DSA) implemented in the Go programming language. It is designed to take learners from basic programming concepts to advanced data structures and algorithms, providing a foundation for competitive programming and technical interviews at top companies.

## Overview

The **go-dsa** repository covers the following topics:
1. **Go Language Fundamentals**: All essential Go language features required to work with data structures and algorithms.
2. **Analysis of Algorithms**: Asymptotic analysis, order of growth, Big O notation, space complexity, and recursion analysis.
3. **Mathematics for DSA**: Fundamental math concepts like counting, prime factorization, and number theory that are often used in algorithms.
4. **Bitwise Operations**: Basic to advanced bitwise techniques used for optimization and solving specific algorithm problems.
5. **Recursion**: Techniques and applications of recursion, including classic problems like the Tower of Hanoi.
6. **Core Data Structures**: Arrays, linked lists, stacks, queues, trees, graphs, heaps, and tries.
7. **Advanced Data Structures**: Segment trees, binary indexed trees, and disjoint sets.
8. **Algorithms**: Sorting, searching, greedy algorithms, dynamic programming, and backtracking.

Each topic includes well-organized code files and test cases to help with understanding and application of the concepts.

---

## Project Structure

The repository is organized as follows:

```plaintext
go-dsa/
├── basics/                  # Go language fundamentals
├── analysis/                # Algorithm analysis techniques
├── data-structures/         # Data structure implementations
│   ├── arrays/
│   ├── linked-lists/
│   ├── stacks/
│   ├── queues/
│   └── trees/
├── algorithms/              # Algorithm implementations
│   ├── sorting/
│   ├── searching/
│   ├── dynamic-programming/
│   └── graph-algorithms/
├── tests/                   # Additional test cases and utilities
├── README.md                # Project overview and setup instructions
└── go.mod                   # Go module file
```

Getting Started
Prerequisites
Go: Ensure that Go is installed on your system. You can download it from golang.org.
Installation

_**Clone the repository:**_

`git clone https://github.com/yourusername/go-dsa.git
cd go-dsa`

_**Initialize Go modules:**_

`go mod init github.com/yourusername/go-dsa`


**_Run tests to verify the setup:_**
`go test ./...`

This will run all the test cases in the project to confirm that everything is set up correctly.

## Topics Covered
**1. Go Language Fundamentals**
   This section covers:

Basic syntax, variables, control structures, and functions.
Pointers and error handling.
Structs, methods, slices, maps, and interfaces.
**2. Analysis of Algorithms**
   This section includes:

Asymptotic analysis (Big O, Omega, and Theta notation).
Best, average, and worst-case analysis.
Complexity analysis for loops and recursion.
Space complexity.

**3. Mathematics for DSA**
   This section includes:

Counting digits, prime numbers, and divisors.
Factorials, GCD, LCM, and prime factorization.
Sieve of Eratosthenes for finding primes efficiently.

**4. Bitwise Operations**
   This section covers:

Basic bitwise operations, including AND, OR, XOR, shifts.
Applications such as checking power of two, counting set bits, and bitwise subsets.

**5. Recursion**
   This section covers:

Basic recursion concepts and common problems (factorial, palindrome check, etc.).
Advanced problems like subset generation, Tower of Hanoi, and Josephus Problem.

**6. Core Data Structures**
   This section covers:

Arrays: Basic operations, searching, sorting, and optimizations.
Linked Lists: Singly, doubly, and circular linked lists, along with problems like reversing a list and detecting cycles.
Stacks and Queues: Array and linked-list implementations, along with classic problems like balanced parentheses and the stock span problem.
Trees and Graphs: Binary trees, binary search trees (BST), traversals, and graph representations (BFS, DFS).

**7. Advanced Data Structures**
   This section includes:

Segment Trees and Binary Indexed Trees for range queries.
Disjoint Sets for union-find operations with path compression and union by rank.

**8. Algorithms**
   This section includes:

Sorting: Bubble, selection, insertion, merge, quicksort, and more.
Searching: Binary search, search in rotated sorted array, and median finding.
Dynamic Programming: Problems like longest increasing subsequence, knapsack, and matrix chain multiplication.
Greedy Algorithms: Activity selection, fractional knapsack, and Huffman coding.
Backtracking: Problems like N-Queens, Sudoku, and Rat in a Maze.

## Running Tests
Each module has corresponding test files (e.g., variables_test.go in the basics folder). You can run all tests across the project using:

`go test ./...`
To run tests for a specific file, navigate to that folder and run:

`go test`
Tests are written using Go's standard testing package.

## Contributing
Contributions are welcome! If you’d like to improve the content, fix bugs, or add new topics, please follow these steps:

**Fork the repository**.
Create a new branch for your feature/bugfix.
Commit your changes with descriptive messages.
Push to your branch and create a pull request.
Please ensure that all code changes are covered by tests.

## License
This project is licensed under the MIT License. See the LICENSE file for more details.

## Roadmap
The repository is divided into the following phases:

**Go Fundamentals:** Covering basic language constructs.

**Algorithm Analysis:** Complexity analysis and asymptotic notation.

**Mathematics and Bitwise Operations:** Essential math and bitwise tricks.

**Core Data Structures:** Arrays, linked lists, stacks, queues, trees, and graphs.

**Advanced Data Structures:** Segment trees, binary indexed trees, and disjoint sets.

**Algorithms:** Sorting, searching, dynamic programming, greedy algorithms, and backtracking.

The repository is structured to allow for a step-by-step understanding of each concept before moving on to more complex topics.


Happy Coding!

