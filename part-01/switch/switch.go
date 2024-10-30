//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	age := rand.IntN(90)
	fmt.Println("I am", age, "years old")
	switch {
	case age == 0:
		//  - "newborn" when age is 0
		fmt.Println("newborn")
	case age >= 1 && age <= 3:
		//  - "toddler" when age is 1, 2, or 3
		fmt.Println("toddler")
	case age < 13:
		//  - "child" when age is 4 through 12
		fmt.Println("child")
	case age < 18:
		//  - "teenager" when age is 13 through 17
		fmt.Println("teenager")
	default:
		//  - "adult" when age is 18+
		fmt.Println("adult")
	}
}
