package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	tr := 15
	size := make([][]int, a)
	for i := 0; i < a; i++ {
		size[i] = make([]int, a)
		for z := 0; z < a; z++ {
			fmt.Scan(&size[i][z])

		}
	}
	for i := 0; i < a; i++ {
		for z := 0; z < a; z++ {
			if size[i][z] == size[z][i] && i != z || a == 1 {
				tr = 1
			} else if size[i][z] != size[z][i] && i != z {
				tr = 0
				break
			}
		}
	}
	if tr == 1 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
