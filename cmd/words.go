package cmd

import (
	"bufio"
	"regexp"
	"strings"
    "os"
    "fmt"
)

func wordsInLine(line string) int {
	re := regexp.MustCompile("\\s+")

	str := strings.TrimSpace(line)
	clean := re.ReplaceAllString(str, " ")

	if len(clean) == 0 {
		return 0
	}

	return strings.Count(clean, " ") + 1
}

func WordsInLines(name string) int {
    file, err := os.Open(name)
    if err != nil {
        fmt.Println("Err ", err)
        }
	words := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words += wordsInLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic("Something went wrong")
	}

	return words
}
