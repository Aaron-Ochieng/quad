package piscine

import (
	"github.com/01-edu/z01"
)

func QuadB(x, y int) {
	if x > 0 && y > 0 {
		// TopLine
		z01.PrintRune('/')
		for i := 0; i < x-2; i++ {
			z01.PrintRune('*')
		}
		if x > 1 {
			z01.PrintRune('\\')
			z01.PrintRune('\n')
		} else {
			z01.PrintRune('\n')

		}

		// Middle lines
		for i := 0; i < y-2; i++ {
			z01.PrintRune('*')
			for j := 0; j < x-2; j++ {
				z01.PrintRune(' ')
			}
			if x > 1 {
				z01.PrintRune('*')
				z01.PrintRune('\n')
			} else {
				z01.PrintRune('\n')
			}
		}

		// Bottom line
		if y > 1 {
			if x > 1 {
				z01.PrintRune('\\')
				for i := 0; i < x-2; i++ {
					z01.PrintRune('*')
				}
				z01.PrintRune('/')
				z01.PrintRune('\n')
			} else {
				z01.PrintRune('\\')
				z01.PrintRune('\n')
			}
		}
	}
}
