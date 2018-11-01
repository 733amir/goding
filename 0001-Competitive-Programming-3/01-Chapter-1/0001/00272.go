package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner, isBeginning := bufio.NewScanner(os.Stdin), true
	for scanner.Scan() {
		text := scanner.Text() + "\n"
		output := strings.Builder{}
		for t := range text {
			if text[t] == '"' {
				if isBeginning {
					output.WriteString("``")
				} else {
					output.WriteString("''")
				}
				isBeginning = !isBeginning
			} else {
				output.WriteByte(text[t])
			}
		}
		fmt.Print(output.String())
	}
}
