package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "5"
	num, _ := strconv.Atoi(str)
	fmt.Println(num * 2) // вывод 10
	var num1 int = 5
	str1 := strconv.Itoa(num1)
	fmt.Println(str1 + str1) // вывод 55
}
