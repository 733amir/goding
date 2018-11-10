package main

import "fmt"

func main() {
	var testCount, shopCount, shopPos, shopMin, shopMax int

	fmt.Scan(&testCount)
	for ; testCount != 0; testCount-- {
		fmt.Scan(&shopCount)
		shopMin, shopMax = 100, -1
		for ; shopCount != 0; shopCount-- {
			fmt.Scan(&shopPos)
			if shopPos < shopMin {
				shopMin = shopPos
			}
			if shopPos > shopMax {
				shopMax = shopPos
			}
		}
		fmt.Println((shopMax - shopMin) * 2)
	}
}
