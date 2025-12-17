package main

import "fmt"

func main() {
	// =======================================
	// 1️⃣ Integers (int)
	// =======================================
	fmt.Println("Integer examples:")
	fmt.Println(1)              // single integer
	fmt.Println(1 + 2)          // addition
	fmt.Println(10 - 3)         // subtraction
	fmt.Println(4 * 5)          // multiplication
	fmt.Println(10 / 3)         // integer division (result truncated)
	fmt.Println(10 % 3)         // modulo operator

	// =======================================
	// 2️⃣ Floating-Point Numbers (float32, float64)
	// =======================================
	fmt.Println("\nFloat examples:")
	fmt.Println(2.4)            // single float
	fmt.Println(2.4 + 7.5)      // addition
	fmt.Println(5.0 / 2.0)      // division with floats
	fmt.Println(3.5 * 2.0)      // multiplication
	fmt.Println(7.5 - 2.2)      // subtraction

	// =======================================
	// 3️⃣ Strings
	// =======================================
	fmt.Println("\nString examples:")
	fmt.Println("Hello, World!") // simple string
	fmt.Println("Go " + "Lang")  // string concatenation

	name := "Aman"
	fmt.Printf("My name is %s\n", name) // formatted string using printf
	fmt.Printf("Hello %s, you are learning %s\n", name, "Go") // multiple placeholders

	// =======================================
	// 4️⃣ Boolean (bool)
	// =======================================
	fmt.Println("\nBoolean examples:")
	fmt.Println(true)              // true
	fmt.Println(false)             // false

	// Boolean expressions
	fmt.Println(3 > 2)             // true
	fmt.Println(5 == 5)            // true
	fmt.Println(7 != 3)            // true
	fmt.Println(4 < 1)             // false

	// =======================================
	// 5️⃣ Type Notes & Best Practices
	// =======================================

	// 1. Integer division truncates the decimal part
	// 2. Floats are by default float64 if not explicitly typed
	// 3. Strings can be concatenated using +
	// 4. Boolean expressions are widely used in conditions (if, for)
	// 5. Use fmt.Printf	