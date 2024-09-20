package main

import (
	"fmt"
)

func main() {
	var b int
	fmt.Scan(&b)
	min := -99999999999
	max := -9999999999
	ch := 0
	for b != 0 {
		ch = b
		fmt.Scan(&b)
		if min < b {
			min = b
		}
		if max < min && max < b {
			max = ch
		}
	}
	fmt.Print(max)
}
