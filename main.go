package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.ReadFile("tetromino/good0.txt")
	if err != nil {
		log.Fatal("error")
	}

	file := string(f)
	count := 0
	tetromino := ""
	for i, c := range file {
		fmt.Println(i+1, c)
		if c == 10 && count > 0 && count != 4 {
			tetromino += "\n"
		}
		if c == '#' {
			if count+1 > 4 {
				fmt.Println("ERROR")
				break
			}
			count++
			isValid := false
			// First Line
			if i >= 0 && i <= 4 {

			}
			if file[i+1] == '#' || file[i-1] == '#' || file[i+4] == '#' || file[i-4] == '#' {
				isValid = true
			}
			if isValid {
				tetromino += "#"
			} else {
				tetromino += "."
			}
		}
	}

	fmt.Println(tetromino)
}
