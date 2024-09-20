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
	var str string
	var result string
	for i := 0; i < len(strk); i++ {
		for a := i; a <= a+dlina; a++ {
			if a+4 < len(strk) {
				str += strk[a]
				continue
			}
			for b := 0; b < len(simb); b++ {
				if strings.Contains(str, simb[b]) == true {
					result = str
					continue
				} else {
					result = "-1"
				}
			}
		}
		str = ""
	}
	fmt.Print(result)
}
