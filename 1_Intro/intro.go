package main

import "fmt"

func main() {
	// =======================================
	// 1️⃣ Hello World in Go
	// =======================================

	// fmt.Println prints the output followed by a newline
	fmt.Println("Hello World")

	// fmt.Printf can be used for formatted output
	name := "Aman"
	fmt.Printf("Hello %s, welcome to Go!\n", name)

	// =======================================
	// 2️⃣ Running Go Programs
	// =======================================

	// There are basically two main ways to run a Go program:

	// 1️⃣ Using 'go build' (compiles to an executable)
	//    Command:
	//       go build path/to/file.go
	//    This generates an executable file in the current directory
	//    Example:
	//       go build main.go
	//       ./main           // run the executable
	//    Notes:
	//       - Use this when you want a standalone executable
	//       - Changes in code require re-building

	// 2️⃣ Using 'go run' (compiles and runs in one step)
	//    Command:
	//       go run path/to/file.go
	//    Example:
	//       go run main.go
	//    Notes:
	//       - Fast for testing small scripts
	//       - Does not generate a separate executable

	// =======================================
	// 3️⃣ Important Notes
	// =======================================

	// 1. Every Go program starts execution from the main() function in package main
	// 2. 'package main' is required for creating executable programs
	// 3. 'fmt' package is used for input/output (printing to console)
	// 4. Use 'go fmt file.go' to format your code according to Go standards
	// 5. Comments can be single-line (//) or multi-line (/* */)

	// =======================================
	// ✅ Example: Running
	// =======================================

	// Using go run:
	//   go run main.go
	// Output:
	//   Hello World
	//   Hello Aman, welcome to Go!

	// Using go build:
	//   go build main.go
	//   ./main
	// Output is the same as above
}
