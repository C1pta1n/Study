package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	balance := 1000

	deposit := func(amount int) {
		defer wg.Done()

		balance += amount
	}

	withdraw := func(amount int) {
		defer wg.Done()

		balance -= amount
	}

	attempts := 1000

	for i := 0; i < attempts; i++ {
		wg.Add(2)

		go deposit(500)
		go withdraw(302)
	}

	wg.Wait()

	fmt.Println(balance)
}
