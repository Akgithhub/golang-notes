package main

import "fmt"

func main() {
	// =======================================
	// 1️⃣ Standard for Loop
	// =======================================

	fmt.Println("Standard for loop:")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// =======================================
	// 2️⃣ While-style Loop (Go has no while)
	// =======================================

	fmt.Println("\nWhile-style loop:")
	a := 1
	for a <= 10 {
		fmt.Println(a)
		a++
	}

	// =======================================
	// 3️⃣ Infinite Loop (do-while equivalent)
	// =======================================

	// Go does not have a do-while loop.
	// This pattern is used as its closest equivalent.
	fmt.Println("\nDo-while equivalent using infinite loop:")
	b := 1
	for {
		fmt.Println(b)
		b++
		if b > 6 {
			break // exit loop
		}
	}

	// =======================================
	// 4️⃣ Range Iteration
	// =======================================

	fmt.Println("\nRange over slice:")
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	fmt.Println("\nRange over map:")
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
	}
	for key, value := range colors {
		fmt.Printf("%s: %s\n", key, value)
	}

	fmt.Println("\nRange over string (runes):")
	for i, char := range "Hello" {
		fmt.Printf("Index: %d, Char: %c\n", i, char)
	}

	// =======================================
	// 5️⃣ continue and break
	// =======================================

	fmt.Println("\nUsing continue:")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // skip current iteration
		}
		fmt.Println(i)
	}

	fmt.Println("\nUsing break:")
	for i := 0; i < 5; i++ {
		if i == 3 {
			break // exit loop
		}
		fmt.Println(i)
	}

	// =======================================
	// 6️⃣ Ignoring Values in range
	// =======================================

	fmt.Println("\nIgnoring index using _:")
	for _, value := range numbers {
		fmt.Println(value)
	}

	// =======================================
	// 7️⃣ Important Notes & Best Practices
	// =======================================

	// 1. Go has only ONE loop keyword: 'for'
	// 2. break exits the loop; continue skips the current iteration
	// 3. range works with slices, arrays, maps, and strings
	// 4. Iterating over strings uses UTF-8 runes, not bytes
	// 5. Use '_' to ignore unused variables
	// 6. Infinite loops must always have a clear exit condition
}
