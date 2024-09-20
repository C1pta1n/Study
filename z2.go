package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	snow := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&snow[i])
	}

	for i := 0; i < n; i++ {
		if snow[i] == -1 {
			if i == 0 {
				snow[i] = 1
			} else {
				if snow[i-1] > 1 {
					snow[i] = snow[i-1]
				} else {
					snow[i] = 1
				}
			}
		} else {
			if i > 0 && snow[i] < snow[i-1] {
				fmt.Println("NO")
				return
			}
		}
	}

	for i := 1; i < n; i++ {
		if snow[i] < snow[i-1] {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(snow[i])
	}
}
