package main

import "fmt"

func main() {
	// =======================================
	// 1️⃣ Declaring Variables in Go
	// =======================================

	// Syntax: var variableName type = value
	var myVar string = "Hello Variables" // Explicit type
	fmt.Println("1:", myVar)

	// =======================================
	// 2️⃣ Type Inference (Shorthand)
	// =======================================

	// Go automatically infers type if you omit it
	var myVarTwo = "Aman" // inferred as string
	fmt.Println("2:", myVarTwo)

	var myBool = true // inferred as bool
	fmt.Println("3:", myBool)

	var myInt = 42 // inferred as int
	fmt.Println("4:", myInt)

	var myFloat = 4.0 // inferred as float64
	fmt.Println("5:", myFloat)

	// =======================================
	// 3️⃣ Short Variable Declaration (Most Common)
	// =======================================

	// := declares and initializes a variable in one step
	myStr := "Hey there!"
	fmt.Println("6:", myStr)

	// You can also declare first and assign later
	var anotherStr string
	anotherStr = "Assigned later"
	fmt.Println("7:", anotherStr)

	// =======================================
	// 4️⃣ Multiple Variable Declaration
	// =======================================

	// Declare multiple variables in one line
	var x, y, z int = 1, 2, 3
	fmt.Println("8:", x, y, z)

	// Short declaration with multiple variables
	a, b := "Hello", 100
	fmt.Println("9:", a, b)

	// =======================================
	// 5️⃣ Constants
	// =======================================

	// Constants cannot be changed once assigned
	const Pi = 3.14159
	fmt.Println("10: Pi =", Pi)

	// =======================================
	// 6️⃣ Zero Values
	// =======================================

	// Variables declared without an explicit initial value have a default zero value
	var defaultInt int    // 0
	var defaultStr string // ""
	var defaultBool bool  // false
	fmt.Println("11:", defaultInt, defaultStr, defaultBool)

	// =======================================
	// 7️⃣ Important Notes & Best Practices
	// =======================================

	// 1. Go does not allow unused variables: declaring without using will cause a compile-time error.
	// 2. Prefer short declarations (:=) inside functions for brevity.
	// 3. Use meaningful variable names for readability.
	// 4. Use explicit types when clarity is important, especially in exported functions.
	// 5. Constants are used for values that should not change.
	// 6. Go has strong type safety; type conversions must be explicit.

	// =======================================
	// 8️⃣ Summary of Declaration Methods
	// =======================================

	// 1) Explicit type:
	var country string = "India"

	// 2) Type inferred:
	var city = "Mumbai"

	// 3) Short declaration (inside functions only):
	name := "Aman"

	fmt.Println("12:", name, city, country)
}
