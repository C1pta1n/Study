package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	ch := 999
	num1, num2, num3, num4 := 0, 0, 0, 0
	for b := 0; b < 5; b++ {
		for i := 0; i <= 9; i++ {
			for z := 0; z < 5; z++ {
				if z == 0 {
					num4 = a % 10
					a = a / 10
				}
				if z == 1 {
					num3 = a % 10
					a = a / 10
				}
				if z == 2 {
					num2 = a % 10
					a = a / 10
				}
				if z == 3 {
					num1 = a % 10
					a = a / 10
				}
			}
			if b == 0 {
				ch = num1
				if ch != i && i != 0 && i >= num1 {
					num1 = i
					break
				}
			}
			if b == 1 {
				ch = num2
				if ch != i && ch != num1 && i >= num2 {
					num2 = i
					break
				}
			}
			if b == 2 {
				ch = num3
				if ch != i && ch != num1 && ch != num2 && i >= num3 {
					num3 = i
					break
				}
			}
			if b == 3 {
				ch = num4
				if ch != i && ch != num3 && ch != num1 && ch != num2 && i >= num4 {
					num4 = i
					break
				}
			}
			if num1 != 0 && num2 != 0 && num3 != 0 && num4 != 0 {
				fmt.Print(num1, num2, num3, num4)
				break
			}
		}
	}
}
