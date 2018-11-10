package main

import "fmt"

func main() {
	var start, stop1, stop2, stop3 int
	for {
		fmt.Scanf("%d%d%d%d\n", &start, &stop1, &stop2, &stop3)
		if start == 0 && stop1 == 0 && stop2 == 0 && stop3 == 0 {
			break
		}
		degree := 360 * 3
		if s := start - stop1; s >= 0 {
			degree += s * 9
		} else {
			degree += 360 + s * 9
		}
		if s := stop2 - stop1; s >= 0 {
			degree += s * 9
		} else {
			degree += 360 + s * 9
		}
		if s := stop2 - stop3; s >= 0 {
			degree += s * 9
		} else {
			degree += 360 + s * 9
		}
		fmt.Println(degree)
	}
}
