package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func QuadE(x, y int) {
	// Check if x and y are positive numbers
	if x > 0 && y > 0 {
		// Print the first line
		z01.PrintRune('A')
		for i := 0; i < x-2; i++ {
			z01.PrintRune('B')
		}
		if x > 1 {
			z01.PrintRune('C')
			z01.PrintRune('\n')
		} else {
			z01.PrintRune('\n')
		}

		// Print the middle lines
		for j := 0; j < y-2; j++ {
			z01.PrintRune('B')
			for i := 0; i < x-2; i++ {
				fmt.Print(" ")
			}
			if x > 1 {
				fmt.Println("B")
			} else {
				z01.PrintRune('\n')
			}
		}

		// Print the last line
		if y > 1 {
			z01.PrintRune('C')
			for i := 0; i < x-2; i++ {
				z01.PrintRune('B')
			}
			if x > 1 {
				fmt.Println("A")
			} else {
				z01.PrintRune('\n')
			}
		}
	}
}
