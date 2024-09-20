package main

import (
	"fmt"
)

func main() {
	var str1 string
	fmt.Scan(&str1)
	cnt := 0
	for i := 0; i < len(str1); i++ {
		if int(str1[cnt]) > 60 && int(str1[cnt]) < 122 {
			if (int(str1[i]) > 1 && int(str1[i]) < 90) || (int(str1[cnt]) == 95) || (int(str1[cnt]) > 97 && int(str1[cnt]) < 122) {
				cnt = 1
			} else {
				cnt = 0
				break
			}
		} else {
			cnt = 0
			break
		}
	}
	if cnt == 1 {
		fmt.Print("YES")
	} else if cnt == 0 {
		fmt.Print("NO")
	}
}