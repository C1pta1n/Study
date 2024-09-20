package main

import (
	"fmt"
	"strings"
)

func main() {
	var stroka string
	fmt.Scan(&stroka)

	var treb string
	fmt.Scan(&treb)

	var dlina int
	fmt.Scan(&dlina)

	simb := strings.Split(treb, "")
	strk := strings.Split(stroka, "")
	fmt.Println(simb)
	fmt.Print(strk)
	var str []string
	var result string
	for i := 0; i < len(strk); i++ {
		for a := i; a < a+dlina; a++ {
			if a+dlina < len(strk) {
				str = append(str, strk[a])
			}
		}
		str = nil
	}
	fmt.Print(result)
}
