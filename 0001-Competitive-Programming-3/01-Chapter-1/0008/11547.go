package main

import "fmt"

func main() {
	var testCount, num int
	fmt.Scan(&testCount)
	for ; testCount > 0; testCount-- {
		fmt.Scan(&num)
		num = (((((num * 567) / 9) + 7492) * 235) / 47 - 498) % 100
		if num < 0 {
			num *= -1
		}
		fmt.Println(num / 10)
	}
}
