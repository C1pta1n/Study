package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	D := math.Pow(b, 2) - 4*a*c

	if D >= 0 {
		if D == 0 {
			fmt.Println(-b / (2 * a))
		} else if D > 0 {
			x1 := (-b + math.Sqrt(D)) / (2 * a)
			x2 := (-b - math.Sqrt(D)) / (2 * a)
			if x1 > x2 {
				fmt.Println(x2)
				fmt.Println(x1)
			} else {
				fmt.Println(x1)
				fmt.Println(x2)
			}
		}
	}
}
