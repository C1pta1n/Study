package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func countDivisors(n int) int {
	if n < 1 {
		return 0
	}
	count := 0
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			if n/i == i {
				count++
			} else {
				count += 2
			}
		}
	}
	return count
}

func countCompositeWithPrimeDivisors(l int, r int) int {
	count := 0
	for i := l; i <= r; i++ {
		if i > 1 && !isPrime(i) {
			divisorCount := countDivisors(i)
			if isPrime(divisorCount) {
				count++
			}
		}
	}
	return count
}

func main() {
	var l, r int
	fmt.Scan(&l, &r)
	result := countCompositeWithPrimeDivisors(l, r)
	fmt.Print(result)
}
