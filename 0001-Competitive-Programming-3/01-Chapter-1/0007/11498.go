package main

import "fmt"

func main() {
	var testCount, n, m, x, y int

	for {
		fmt.Scan(&testCount, &n, &m)
		if testCount == 0 {
			break
		}
		for ; testCount != 0; testCount-- {
			fmt.Scan(&x, &y)
			if x ==n || y == m {
				fmt.Println("divisa")
				continue
			}

			if y > m {
				fmt.Print("N")
			} else {
				fmt.Print("S")
			}

			if x > n {
				fmt.Println("E")
			} else {
				fmt.Println("O")
			}
		}
	}
}
