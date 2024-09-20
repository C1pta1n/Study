package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)

	numbers := make(map[int]struct{})

	parts := strings.Split(input, ",")
	for _, part := range parts {
		if strings.Contains(part, "-") {
			rangeParts := strings.Split(part, "-")
			start, _ := strconv.Atoi(rangeParts[0])
			end, _ := strconv.Atoi(rangeParts[1])
			for i := start; i <= end; i++ {
				numbers[i] = struct{}{}
			}
		} else {
			num, _ := strconv.Atoi(part)
			numbers[num] = struct{}{}
		}
	}

	var result []int
	for num := range numbers {
		result = append(result, num)
	}
	sort.Ints(result)

	for i, num := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(num)
	}
}
