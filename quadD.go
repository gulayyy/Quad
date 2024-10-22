package piscine

import "fmt"

func QuadD(x, y int) {
	if x > 0 && y > 0 {
		for a := 0; a < x; a++ {
			if a == 0 {
				fmt.Print("A")
			} else if a == x-1 {
				fmt.Print("C")
			} else {
				fmt.Print("B")
			}
		}
		fmt.Println()

		for b := 0; b < y-2; b++ {
			fmt.Print("B")
			for c := 0; c < x-2; c++ {
				fmt.Print(" ")
			}
			if x > 1 {
				fmt.Print("B")
			}
			fmt.Println()
		}
		if y > 1 {
			for a := 0; a < x; a++ {
				if a == 0 {
					fmt.Print("A")
				} else if a == x-1 {
					fmt.Print("C")
				} else {
					fmt.Print("B")
				}
			}
		}
	}
}
