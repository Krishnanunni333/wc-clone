package workers

import (
	"bufio"
	"regexp"
	"strings"
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

func WordsInLines(str string) int {

    words := 0
	scanner := bufio.NewScanner(strings.NewReader(str))

	for scanner.Scan() {
		words += wordsInLine(scanner.Text())
	}

	return words
}
