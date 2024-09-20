package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	del := 1
	y1 := 4
	for i := 1; i <= a; i++ {
		for z := 0; z <= a; z++ {
			for y := 2; y < z; y++ {
				if z%y == 0 && y != z {
					break
				}
				y1 = y
			}
			if z%y1 != 0 {
				del = z
				fmt.Print(z)
			}
		}
		a = a / del

	}
}
