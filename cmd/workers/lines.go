package workers

import (
	"bufio"
	"os"
    "fmt"
)

func CountLines(name string) int {
    file, err := os.Open(name)
    if err != nil {
        fmt.Println("Err ", err)
        }
	lines := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
	}

	if err := scanner.Err(); err != nil {
		panic("Something went wrong")
	}

	return lines
}
