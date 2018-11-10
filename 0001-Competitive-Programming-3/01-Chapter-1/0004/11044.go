package main

import (
	"fmt"
	"math"
)

func main() {
	var testCount, x, y int
	fmt.Scan(&testCount)
	for ; testCount != 0; testCount-- {
		fmt.Scan(&x, &y)
		fmt.Println(int64(math.Ceil(float64(x-2)/3) * math.Ceil(float64(y-2)/3)))
	}
}
