package workers

import (
	"bufio"
	"os"
    "fmt"
)

func CountChars(name string) int {
    file, err := os.Open(name)
    if err != nil {
        fmt.Println("Err ", err)
        }
	characters := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        characters += len(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic("Something went wrong")
	}

	return characters
}
