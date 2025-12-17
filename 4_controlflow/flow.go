package main

import "fmt"

func main() {
	// =======================================
	// 1️⃣ if – else if – else Statement
	// =======================================

	age := 21

	if age <= 25 {
		fmt.Println("You have time to learn and grow.")
	} else if age <= 40 {
		fmt.Println("You still have time, but with more responsibilities.")
	} else {
		fmt.Println("You have done well. Keep going!")
	}

	// =======================================
	// 2️⃣ switch Statement
	// =======================================

	day := "Wednesday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week")

	case "Wednesday":
		fmt.Println("Midweek – still a long way to go")

	case "Friday":
		fmt.Println("Weekend is near")

	case "Saturday":
		fmt.Println("Finally, time to relax!")

	default:
		fmt.Println("Sunday – rest and recharge")
	}

	// =======================================
	// 3️⃣ switch with Multiple Values
	// =======================================

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend 🎉")
	default:
		fmt.Println("It's a weekday")
	}

	// =======================================
	// 4️⃣ switch without Expression
	// =======================================

	score := 85

	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 75:
		fmt.Println("Grade: B")
	case score >= 60:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: D")
	}

	// =======================================
	// 5️⃣ Important Notes & Best Practices
	// =======================================

	// 1. Go does NOT require parentheses around conditions
	// 2. Curly braces {} are mandatory
	// 3. switch in Go automatically breaks after a case
	// 4. 'fallthrough' can be used if you want to continue to the next case
	// 5. switch is preferred over long if-else chains for readability
}
