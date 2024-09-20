package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	name := scanner.Text()
	re := regexp.MustCompile("[0-9]+")
	ints := re.FindAllString(name, -1)
	a, b := "", ""
	var c, d int
	var err, err1 error
	for i := 0; i < len(ints); i += 2 {
		if i+1 < len(ints) {
			a = ints[i]
			b = ints[i+1]
			c, err = strconv.Atoi(a)
			d, err1 = strconv.Atoi(b)
			for g := c; g <= d; g++ {
				fmt.Print(g)
			}
			if err != nil && err1 != nil {
				fmt.Print("")
			}
		}
	}
	fmt.Print(ints[len(ints)-1])
}
