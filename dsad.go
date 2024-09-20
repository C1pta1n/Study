package main

import (
	"bufio"
	"fmt"
	"os"
)

func IsDigitsOnly(s string) int {
	var a, b int
	for _, c := range s {
		if c > '0' && c < '9' {
			a = 'c'
			if c > '0' && c < '9' && a != 0 {
				b = 'c'
				for i := a; i <= b; i++ {
					return i
				}
			}
		}
	}
	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	name := scanner.Text()
	name1 := IsDigitsOnly(name)
	fmt.Print(name1)
}
