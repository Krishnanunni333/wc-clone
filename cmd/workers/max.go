package workers

import (
	"bufio"
	"os"
    "fmt"
)

func MaxLine(name string) int {
    file, err := os.Open(name)
    if err != nil {
        fmt.Println("Err ", err)
        }
	characters := 0
	max := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        if len(scanner.Text()) > max{
            max = len(scanner.Text())
        }
	}

	if err := scanner.Err(); err != nil {
		panic("Something went wrong")
	}

	return characters
}
