package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	fmt.Scan(&x1, &y1, &x2, &y2)
	if (x1 == x2 || y1 == y2) || (((math.Abs(x1-x2) == 1) && (math.Abs(y1-y2) == 1)) || ((math.Abs(x1-x2) == 2) && (math.Abs(y1-y2) == 2)) || ((math.Abs(x1-x2) == 3) && (math.Abs(y1-y2) == 3)) || ((math.Abs(x1-x2) == 4) && (math.Abs(y1-y2) == 4))) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
