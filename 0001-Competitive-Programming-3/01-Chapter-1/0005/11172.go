package main

import "fmt"

func main() {
	var testCount int
	var a, b int64

	fmt.Scan(&testCount)
	for ; testCount != 0; testCount-- {
		fmt.Scan(&a, &b)
		if a > b {
			fmt.Println(">")
		} else if a < b {
			fmt.Println("<")
		} else {
			fmt.Println("=")
		}
	}
}
