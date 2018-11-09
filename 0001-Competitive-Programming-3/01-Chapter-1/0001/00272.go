package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	buffer, isBeginning := make([]byte, 32), true
	for {
		characterCount, err := os.Stdin.Read(buffer)
		if characterCount == 0 || err != nil {
			break
		}

		output := strings.Builder{}
		for i := 0; i < characterCount; i++ {
			if buffer[i] == '"' {
				if isBeginning {
					output.WriteString("``")
				} else {
					output.WriteString("''")
				}
				isBeginning = !isBeginning
			} else {
				output.WriteByte(buffer[i])
			}
		}
		fmt.Print(output.String())
	}
}
