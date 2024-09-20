package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	_ = scan.Scan()
	str := scan.Text()
	cnt := 0
	str1 := ""
	for i := 0; i < len(str); i++ {
		if string(str[i]) == " " {
			cnt++
		}
		if (cnt > 1 && i+1 < len(str) && string(str[i+1]) != " ") || cnt == 0 {
			str1 = str1 + string(str[i])
		} else if i+1 < len(str) && string(str[i+1]) != " " {
			cnt = 0
		}
	}
	fmt.Print(str1)
}
