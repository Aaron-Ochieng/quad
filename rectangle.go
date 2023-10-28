package main

import "fmt"

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		fmt.Print("o")
		for i := 0; i < x-2; i++ {
			fmt.Print("-")
		}
		fmt.Println("o")

		for i := 0; i < y-2; i++ {
			fmt.Print("|")
			for j := 0; j < x-2; j++ {
				fmt.Print(" ")
			}
			fmt.Println("|")
		}

		if y > 1 {
			fmt.Print("o")
			for i := 0; i < x-2; i++ {
				fmt.Print("-")
			}
			fmt.Println("o")
		}
	}
}

func main() {
	QuadA(5, 3)
}
