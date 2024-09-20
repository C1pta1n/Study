package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	name := scanner.Text()
	var a, b int
	var n []int
	for c := 0; c < len(name); c++ {
		if _, err := strconv.Atoi(name[c]); err == nil {
			if a != 0 {
				b = strconv.Atoi(name[i])
				for i := a; i <= b; i++ {
					n = append(n, i)
				}
				a = 0
				b = 0
			}
			a = strconv.Atoi(name[i])
		}
	}
	fmt.Print(n)
}
