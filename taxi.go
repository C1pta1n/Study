package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	size := make([]int, a)
	for i := 0; i < a; i++ {
		fmt.Scan(&size[i])
	}
	cnt := 0
	for z := 0; z < a; z++ {
		for i := 0; i < a; i++ {
			if size[z] == 4 {
				cnt++
				size[z] = 5
			}
			if size[z]+size[i] < 4 && z != i {
				cnt++
				size[z] = 5
				size[i] = 5
			}
			if size[i] < 4 {
				cnt++
				size[i] = 5
			}
		}
		break
	}
	fmt.Print(cnt)
}
